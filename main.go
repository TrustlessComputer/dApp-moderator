package main

import (
	"context"
	"dapp-moderator/external/block_stream"
	"dapp-moderator/external/blockchain_api"
	"dapp-moderator/external/bns_service"
	moralis2 "dapp-moderator/external/moralis"
	"dapp-moderator/external/token_explorer"
	"dapp-moderator/internal/generative_respository"
	discordclient "dapp-moderator/utils/discord"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"time"

	"go.uber.org/zap"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	"dapp-moderator/external/bfs_service"
	"dapp-moderator/external/nft_explorer"
	"dapp-moderator/external/quicknode"
	"dapp-moderator/internal/delivery"
	httpHandler "dapp-moderator/internal/delivery/http"
	"dapp-moderator/internal/delivery/txTCServer"
	"dapp-moderator/internal/repository"
	"dapp-moderator/internal/usecase"
	"dapp-moderator/utils/config"
	"dapp-moderator/utils/connections"
	"dapp-moderator/utils/global"
	"dapp-moderator/utils/googlecloud"
	_logger "dapp-moderator/utils/logger"
	"dapp-moderator/utils/oauth2service"
	"dapp-moderator/utils/redis"

	"github.com/gorilla/mux"
)

var logger _logger.Ilogger
var mongoConnection connections.IConnection
var generativeMongoConnection connections.IConnection
var conf *config.Config

func init() {
	logger = _logger.NewLogger(true)

	c, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	mongoCnn := fmt.Sprintf("%s://%s:%s@%s/?retryWrites=true&w=majority", c.Databases.Mongo.Scheme, c.Databases.Mongo.User, c.Databases.Mongo.Pass, c.Databases.Mongo.Host)
	if c.ENV == "local" {
		mongoCnn = "mongodb://127.0.0.1:27017/tc_dapp?retryWrites=true&w=majority"
	}

	mongoDbConnection, err := connections.NewMongo(mongoCnn)
	if err != nil {
		logger.AtLog().Logger.Error("Cannot connect mongoDB ", zap.Error(err))
		panic(err)
	}

	generativeMongoCnn := fmt.Sprintf("%s://%s:%s@%s/?retryWrites=true&w=majority", c.Databases.GenerativeMongo.Scheme, c.Databases.GenerativeMongo.User, c.Databases.GenerativeMongo.Pass, c.Databases.GenerativeMongo.Host)
	if c.ENV == "dev" {
		generativeMongoCnn = "mongodb://127.0.0.1:27017/?retryWrites=true&w=majority"
	}

	gccnn, err := connections.NewMongo(generativeMongoCnn)
	if err != nil {
		logger.AtLog().Logger.Error("Cannot connect mongoDB ", zap.Error(err))
		//panic(err)
	}

	conf = c
	mongoConnection = mongoDbConnection
	generativeMongoConnection = gccnn
}

// @title tcDAPP APIs
// @version 1.0.0
// @description This is a sample server TC-DAPP server.

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @BasePath /dapp-moderator/v1
func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	// log.Println("init sentry ...")
	// sentry.InitSentry(conf)
	startServer()
}

func startServer() {
	logger.AtLog().Logger.Info("starting server ...")
	cache, redisClient := redis.NewRedisCache(conf.Redis)
	r := mux.NewRouter()
	gcs, err := googlecloud.NewDataGCStorage(*conf)
	qn := quicknode.NewQuickNode(conf, cache)
	bst := block_stream.NewBlockStream(conf, cache)
	nex := nft_explorer.NewNftExplorer(conf, cache)
	bfs := bfs_service.NewBfsService(conf, cache)
	bns := bns_service.NewBNSService(conf, cache)
	tke := token_explorer.NewTokenExplorer(conf, cache)
	dcl := discordclient.NewClient()
	bca := blockchain_api.NewBlockChainApi(conf, cache)
	moralis := moralis2.NewMoralisService(cache)

	auth2Service := oauth2service.NewAuth2()

	s3Adapter := googlecloud.NewS3Adapter(googlecloud.S3AdapterConfig{
		BucketName: conf.Gcs.Bucket,
		Endpoint:   conf.Gcs.Endpoint,
		Region:     conf.Gcs.Region,
		AccessKey:  conf.Gcs.AccessKey,
		SecretKey:  conf.Gcs.SecretKey,
	}, redisClient)

	g := global.Global{
		MuxRouter:              r,
		Conf:                   conf,
		DBConnection:           mongoConnection,
		GenerativeDBConnection: generativeMongoConnection,
		Cache:                  cache,
		GCS:                    gcs,
		QuickNode:              qn,
		BlockStream:            bst,
		NftExplorer:            nex,
		BfsService:             bfs,
		BnsService:             bns,
		TokenExplorer:          tke,
		Auth2:                  auth2Service,
		DiscordClient:          dcl,
		BlockChainApi:          bca,
		Moralis:                moralis,
		S3Adapter:              &s3Adapter,
	}

	repo, err := repository.NewRepository(&g)
	if err != nil {
		logger.AtLog().Logger.Error("Cannot init repository", zap.Error(err))
		return
	}

	generativeRepo, err := generative_respository.NewGenerativeRepository(&g)
	if err != nil {
		logger.AtLog().Logger.Error("Cannot init repository", zap.Error(err))
		return
	}

	// migration - we don't need this
	//migrate.SetDatabase(repo.DB)
	//if migrateErr := migrate.Up(-1); migrateErr != nil {
	//	logger.AtLog().Error("migrate failed", zap.Error(err))
	//}

	uc, err := usecase.NewUsecase(&g, repo, generativeRepo)
	if err != nil {
		logger.AtLog().Error("LoadUsecases - Cannot init usecase", zap.Error(err))
		return
	}
	//uc.TestSendNotify()
	//uc.UpdateAllCollectionIndex(context.Background())

	servers := make(map[string]delivery.AddedServer)
	// api fixed run:
	h, _ := httpHandler.NewHandler(&g, *uc)
	servers["http"] = delivery.AddedServer{
		Server:  h,
		Enabled: false,
	}

	txConsumerStatr := os.Getenv("TX_CONSUMER_START")
	txConsumerStatrBool, err := strconv.ParseBool(txConsumerStatr)
	if err != nil {
		txConsumerStatrBool = true //alway start this server, if config is missing
	}
	txConsumerStatrBool = true

	trendingStart := os.Getenv("TRENDING_SERVER_START")
	trendingStartBool, err := strconv.ParseBool(trendingStart)
	if err != nil {
		trendingStartBool = true //alway start this server, if config is missing
	}

	jobSendDiscordStart := os.Getenv("JOB_SEND_DISCORD_START")
	jobSendDiscordStartBool, err := strconv.ParseBool(jobSendDiscordStart)
	if err != nil {
		jobSendDiscordStartBool = true //alway start this server, if config is missing
	}

	tx, _ := txTCServer.NewTxTCServer(&g, *uc)
	servers["tx-consumer"] = delivery.AddedServer{
		Server:  tx,
		Enabled: txConsumerStatrBool,
	}

	servers["job-discord"] = delivery.AddedServer{
		Server:  txTCServer.NewJobDisCord(&g, *uc),
		Enabled: jobSendDiscordStartBool,
	}

	trendingServer, _ := txTCServer.NewTrendingServer(&g, *uc)
	servers["job-trending"] = delivery.AddedServer{
		Server:  trendingServer,
		Enabled: trendingStartBool,
	}

	//var wait time.Duration
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Run our server in a goroutine so that it doesn't block.
	for name, server := range servers {
		if server.Enabled {
			if server.Server != nil {
				go server.Server.StartServer()
			}
			logger.AtLog().Logger.Info(fmt.Sprintf("%s is enabled", name))
		} else {
			logger.AtLog().Logger.Info(fmt.Sprintf("%s is disabled", name))
		}
	}

	// Block until we receive our signal.
	<-c
	wait := time.Second
	// // Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// // Doesn't block if no connections, but will otherwise wait
	// // until the timeout deadline.
	// err := srv.Shutdown(ctx)
	// if err != nil {
	// 	logger.AtLog().Logger.Error("httpDelivery.StartServer - Server can not shutdown", err)
	// 	return
	// }
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	<-ctx.Done() //if your application should wait for other services
	// to finalize based on context cancellation.
	logger.AtLog().Logger.Warn("httpDelivery.StartServer - server is shutting down")
	tracer.Stop()
	os.Exit(0)

}
