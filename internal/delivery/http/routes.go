package http

import (
	"os"

	"dapp-moderator/docs"
	_ "dapp-moderator/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func (h *httpDelivery) registerRoutes() {
	h.RegisterDocumentRoutes()
	h.RegisterV1Routes()
}

func (h *httpDelivery) RegisterV1Routes() {
	h.Handler.Use(h.MiddleWare.LoggingMiddleware)
	h.Handler.Use(h.MiddleWare.Pagination)
	//api
	api := h.Handler.PathPrefix("/dapp/api").Subrouter()
	//AUTH
	auth := api.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/nonce", h.generateMessage).Methods("POST")
	auth.HandleFunc("/nonce/verify", h.verifyMessage).Methods("POST")

	//quicknode
	quicknode := api.PathPrefix("/quicknode").Subrouter()
	quicknode.HandleFunc("/address/{walletAddress}/balance", h.addressBalance).Methods("GET")

	//nftExplorer
	nftExplorer := api.PathPrefix("/nft-explorer").Subrouter()
	nftExplorer.HandleFunc("/collections", h.collections).Methods("GET")
	nftExplorer.HandleFunc("/collections/{contractAddress}", h.collectionDetail).Methods("GET")
	nftExplorer.HandleFunc("/collections/{contractAddress}/nfts", h.collectionNfts).Methods("GET")
	nftExplorer.HandleFunc("/collections/{contractAddress}/nfts/{tokenID}", h.collectionNftDetail).Methods("GET")
	nftExplorer.HandleFunc("/collections/{contractAddress}/nfts/{tokenID}/content", h.collectionNftContent).Methods("GET")
	nftExplorer.HandleFunc("/nfts", h.nfts).Methods("GET")
	nftExplorer.HandleFunc("/owner-address/{ownerAddress}/nfts", h.nftByWalletAddress).Methods("GET")

	nftExplorerAuth := api.PathPrefix("/nft-explorer").Subrouter()
	nftExplorerAuth.Use(h.MiddleWare.ValidateAccessToken)
	nftExplorerAuth.HandleFunc("/collections/{contractAddress}", h.updateCollectionDetail).Methods("PUT")

	//bfs services
	bfsServices := api.PathPrefix("/bfs-service").Subrouter()
	bfsServices.HandleFunc("/files/{walletAddress}", h.bfsFiles).Methods("GET")
	bfsServices.HandleFunc("/browse/{walletAddress}", h.bfsBrowseFile).Methods("GET")
	bfsServices.HandleFunc("/info/{walletAddress}", h.bfsFileInfo).Methods("GET")
	bfsServices.HandleFunc("/content/{walletAddress}", h.bfsFileContent).Methods("GET")

	//bns services
	bnsServices := api.PathPrefix("/bns-service").Subrouter()
	bnsServices.HandleFunc("/names", h.bnsNames).Methods("GET")
	bnsServices.HandleFunc("/names/{name}", h.bnsName).Methods("GET")
	bnsServices.HandleFunc("/names/{name}/available", h.bnsNameAvailable).Methods("GET")
	bnsServices.HandleFunc("/names/owned/{wallet_address}", h.bnsNameOwnedByWalletAddress).Methods("GET")

	// token explorer
	tokenRoutes := api.PathPrefix("/token-explorer").Subrouter()
	tokenRoutes.HandleFunc("/tokens", h.getTokens).Methods("GET")
	tokenRoutes.HandleFunc("/token/{address}", h.getToken).Methods("GET")
	tokenRoutes.HandleFunc("/token/{address}", h.updateToken).Methods("PUT")

	walletInfoGroup := api.PathPrefix("/wallets").Subrouter()
	walletInfoGroup.HandleFunc("/{walletAddress}", h.walletInfo).Methods("GET")
	walletInfoGroup.HandleFunc("/{walletAddress}/txs", h.walletTxs).Methods("GET")

	//profile
	profile := api.PathPrefix("/profile").Subrouter()
	profile.HandleFunc("/wallet/{walletAddress}", h.profileByWallet).Methods("GET")
	profile.HandleFunc("/wallet/{walletAddress}/allowed-list/existed", h.profileByWalletExistedAllowedList).Methods("GET")
	profile.HandleFunc("/wallet/{walletAddress}/histories", h.currentUerProfileHistories).Methods("GET")
	profile.HandleFunc("/wallet/{walletAddress}/collections", h.currentUserProfileCollections).Methods("GET")
	profile.HandleFunc("/wallet/{walletAddress}/tokens/bought", h.currentUserProfileBoughtTokens).Methods("GET")

	profileAuth := api.PathPrefix("/profile").Subrouter()
	profileAuth.Use(h.MiddleWare.AuthorizationFunc)
	profileAuth.HandleFunc("/me", h.currentUerProfile).Methods("GET")
	profileAuth.HandleFunc("/histories", h.createProfileHistory).Methods("POST")
	profileAuth.HandleFunc("/histories", h.confirmProfileHistory).Methods("PUT")

	uploadRoute := api.PathPrefix("/upload").Subrouter()
	// uploadRoute.Use(h.MiddleWare.AuthorizationFunc) // temp pause
	uploadRoute.HandleFunc("/file", h.uploadFile).Methods("POST")

	tools := api.PathPrefix("/tools").Subrouter()
	tools.HandleFunc("/compile-contract", h.compileContract).Methods("POST")

	dappInfo := api.PathPrefix("/dapp-info").Subrouter()
	dappInfo.HandleFunc("/create", h.createDAppInfo).Methods("POST")
	dappInfo.HandleFunc("/list", h.listDAppInfo).Methods("GET")

	//admin
	admin := api.PathPrefix("/admin").Subrouter()
	admin.HandleFunc("/redis", h.getRedisKeys).Methods("GET")
	admin.HandleFunc("/redis/{key}", h.getRedis).Methods("GET")
	admin.HandleFunc("/redis", h.upsertRedis).Methods("POST")
	admin.HandleFunc("/redis", h.deleteAllRedis).Methods("DELETE")
	admin.HandleFunc("/redis/{key}", h.deleteRedis).Methods("DELETE")

	// uniswap
	swapRoutes := api.PathPrefix("/swap").Subrouter()
	swapRoutes.HandleFunc("/tm-scan-event", h.swapTmTokenScanEvents).Methods("GET")
	swapRoutes.HandleFunc("/scan-event", h.swapScanEvents).Methods("GET")
	swapRoutes.HandleFunc("/total-supply", h.jobUpdateTotalSupply).Methods("GET")

	swapRoutes.HandleFunc("/btc-price", h.jobGetBtcPrice).Methods("GET")
	swapRoutes.HandleFunc("/scan-pair-event", h.swapScanPairEvents).Methods("GET")
	swapRoutes.HandleFunc("/scan", h.swapScanHash).Methods("GET")
	swapRoutes.HandleFunc("/clear-cache", h.clearCache).Methods("GET")
	swapRoutes.HandleFunc("/update-sync", h.jobUpdateDataSwapSync).Methods("GET")
	swapRoutes.HandleFunc("/update-history", h.jobUpdateDataSwapHistory).Methods("GET")
	swapRoutes.HandleFunc("/update-pair", h.jobUpdateDataSwapPair).Methods("GET")
	swapRoutes.HandleFunc("/update-token", h.jobUpdateDataSwapToken).Methods("GET")
	swapRoutes.HandleFunc("/fe-log", h.addFrontEndLog).Methods("POST")
	swapRoutes.HandleFunc("/report/slack", h.getSlackReport).Methods("GET")

	jobRoutes := swapRoutes.PathPrefix("/job").Subrouter()
	jobRoutes.HandleFunc("/update-ido", h.swapJobUpdateIdoStatus).Methods("GET")

	swapTransactions := swapRoutes.PathPrefix("/transactions").Subrouter()
	swapTransactions.HandleFunc("/pending", h.findPendingTransactionHistories).Methods("GET")


	swapTokensRoutes := swapRoutes.PathPrefix("/token").Subrouter()
	swapTokensRoutes.HandleFunc("/list", h.getTokensInPool).Methods("GET")
	swapTokensRoutes.HandleFunc("/route", h.getRoutePair).Methods("GET")
	swapTokensRoutes.HandleFunc("/report", h.getTokensReport).Methods("GET")
	swapTokensRoutes.HandleFunc("/price", h.getTokensPrice).Methods("GET")

	userRoutes := swapRoutes.PathPrefix("/user").Subrouter()
	userRoutes.HandleFunc("/trade-histories", h.findUserSwapHistories).Methods("GET")

	swapPairRoutes := swapRoutes.PathPrefix("/pair").Subrouter()
	swapPairRoutes.HandleFunc("/list", h.findSwapPairs).Methods("GET")
	swapPairRoutes.HandleFunc("/trade-histories", h.findSwapHistories).Methods("GET")
	swapPairRoutes.HandleFunc("/apr", h.getLiquidityApr).Methods("GET")

	transactions := api.PathPrefix("/transactions").Subrouter()
	transactions.HandleFunc("/scan-txs", h.swapTransactions).Methods("GET")

	idoRoutes := swapRoutes.PathPrefix("/ido").Subrouter()
	idoRoutes.HandleFunc("/", h.addOrUpdateSwapIdo).Methods("POST")
	idoRoutes.HandleFunc("/list", h.findSwapIdoHistories).Methods("GET")
	idoRoutes.HandleFunc("/detail", h.findSwapIdoDetail).Methods("GET")
	idoRoutes.HandleFunc("/delete", h.deleteSwapIdo).Methods("DELETE")
	idoRoutes.HandleFunc("/tokens", h.getSwapIdoTokens).Methods("GET")

	tmRoutes := swapRoutes.PathPrefix("/tm").Subrouter()
	tmRoutes.HandleFunc("/histories", h.findTmTokenHistories).Methods("GET")
}

func (h *httpDelivery) RegisterDocumentRoutes() {
	documentUrl := `/dapp/swagger/`
	domain := os.Getenv("swagger_domain")
	docs.SwaggerInfo.Host = domain
	docs.SwaggerInfo.BasePath = "/dapp/api"
	swaggerURL := documentUrl + "swagger/doc.json"
	h.Handler.PathPrefix(documentUrl).Handler(httpSwagger.Handler(
		httpSwagger.URL(swaggerURL), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		//httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	))
}
