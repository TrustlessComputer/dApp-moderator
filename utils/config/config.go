package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Debug         bool
	StartHTTP     bool
	Context       *Context
	Databases     *Databases
	Redis         RedisConfig
	ENV           string
	ServicePort   string
	QuickNode     string
	NftExplorer   string
	TokenExplorer string
	BFSService    string
	Gcs           *GCS
}

type Context struct {
	TimeOut int
}

type Databases struct {
	Postgres *DBConnection
	Mongo    *DBConnection
}

type DBConnection struct {
	Host    string
	Port    string
	User    string
	Pass    string
	Name    string
	Sslmode string
	Scheme  string
}

type Mongo struct {
	DBConnection
}

type GCS struct {
	ProjectId string
	Bucket    string
	Auth      string
	Endpoint  string
	Region    string
	AccessKey string
	SecretKey string
}

type RedisConfig struct {
	Address  string
	Password string
	DB       string
	ENV      string
}

func NewConfig(filePaths ...string) (*Config, error) {
	if len(filePaths) > 0 {
		godotenv.Load(filePaths[0])
	} else {
		godotenv.Load()
	}
	services := make(map[string]string)
	isDebug, _ := strconv.ParseBool(os.Getenv("DEBUG"))
	isStartHTTP, _ := strconv.ParseBool(os.Getenv("START_HTTP"))

	timeOut, err := strconv.Atoi(os.Getenv("CONTEXT_TIMEOUT"))
	if err != nil {
		panic(err)
	}

	services["og"] = os.Getenv("OG_SERVICE_URL")
	conf := &Config{
		ENV:       os.Getenv("ENV"),
		StartHTTP: isStartHTTP,
		Context: &Context{
			TimeOut: timeOut,
		},
		Debug:         isDebug,
		ServicePort:   os.Getenv("SERVICE_PORT"),
		QuickNode:     os.Getenv("QUICKNODE_URL"),
		NftExplorer:   os.Getenv("NFT_EXPLORER_URL"),
		TokenExplorer: os.Getenv("TOKEN_EXPLORER_URL"),
		BFSService:    os.Getenv("BFS_SERVICE_URL"),
		Databases: &Databases{
			Mongo: &DBConnection{
				Host:   os.Getenv("MONGO_HOST"),
				Port:   os.Getenv("MONGO_PORT"),
				User:   os.Getenv("MONGO_USER"),
				Pass:   os.Getenv("MONGO_PASSWORD"),
				Name:   os.Getenv("MONGO_DB"),
				Scheme: os.Getenv("MONGO_SCHEME"),
			},
		},
		Redis: RedisConfig{
			Address:  os.Getenv("REDIS_ADDR"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       os.Getenv("REDIS_DB"),
			ENV:      os.Getenv("REDIS_ENV"),
		},
		Gcs: &GCS{
			ProjectId: os.Getenv("GCS_PROJECT_ID"),
			Bucket:    os.Getenv("GCS_BUCKET"),
			Auth:      os.Getenv("GCS_AUTH"),
			Endpoint:  os.Getenv("GCS_ENDPOINT"),
			Region:    os.Getenv("GCS_REGION"),
			AccessKey: os.Getenv("GCS_ACCESS_KEY"),
			SecretKey: os.Getenv("GCS_SECRET_KEY"),
		},
	}

	return conf, nil
}
