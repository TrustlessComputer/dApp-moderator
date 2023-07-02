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
	nftExplorer.HandleFunc("/refresh-nft/contracts/{contractAddress}/token/{tokenID}", h.refreshNft).Methods("GET")

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
	bnsServices.HandleFunc("/names/{token_id}", h.bnsName).Methods("GET")
	bnsServices.HandleFunc("/names/{token_id}/available", h.bnsNameAvailable).Methods("GET")
	bnsServices.HandleFunc("/names/owned/{wallet_address}", h.bnsNameOwnedByWalletAddress).Methods("GET")
	bnsServices.HandleFunc("/default/{wallet_address}", h.bnsDefault).Methods("GET")

	bnsServicesAuth := api.PathPrefix("/bns-service").Subrouter()
	bnsServicesAuth.Use(h.MiddleWare.ValidateAccessToken)
	bnsServicesAuth.HandleFunc("/default/{wallet_address}", h.updateBnsDefault).Methods("PUT")

	//auction
	auctionRoutesPublic := api.PathPrefix("/auction").Subrouter()
	//auctionRoutesPublic.Use(h.MiddleWare.ValidateAccessToken)
	auctionRoutesPublic.HandleFunc("/detail/{contractAddress}/{tokenID}", h.auctionDetail).Methods("GET")
	auctionRoutesPublic.HandleFunc("/list-bid", h.listBid).Methods("GET")

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
	uploadRoute.HandleFunc("/file", h.filterUploadedFile).Methods("GET")
	uploadRoute.HandleFunc("/file", h.uploadFile).Methods("POST")
	uploadRoute.HandleFunc("/file-size", h.calculateUploadedFile).Methods("POST")
	//uploadRoute.HandleFunc("/file/multipart-fake", h.uploadFileMultiPartFake).Methods("POST")
	uploadRoute.HandleFunc("/file/{file_id}/tx_hash/{tx_hash}", h.updateTxHashUploadedFile).Methods("PUT")
	uploadRoute.HandleFunc("/file/{file_id}/chunks", h.fileChunks).Methods("GET")
	uploadRoute.HandleFunc("/file/{file_id}/chunks/{chunk_id}", h.getChunkByID).Methods("GET")
	uploadRoute.HandleFunc("/file/{file_id}/chunks/{chunk_id}/tx_hash/{tx_hash}", h.updateTxHashForAChunk).Methods("PUT")

	uploadRoute.HandleFunc("/multipart", h.CreateMultipartUpload).Methods("POST")
	uploadRoute.HandleFunc("/multipart/{uploadID}", h.UploadPart).Methods("PUT")
	uploadRoute.HandleFunc("/multipart/{uploadID}", h.CompleteMultipartUpload).Methods("POST")

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
	swapRoutes.HandleFunc("/fe-log", h.addFrontEndLog).Methods("POST")
	swapRoutes.HandleFunc("/bot-config", h.addSwapBotConfig).Methods("POST")
	swapRoutes.HandleFunc("/report/slack", h.getSlackReport).Methods("GET")

	swapTransactions := swapRoutes.PathPrefix("/transactions").Subrouter()
	swapTransactions.HandleFunc("/pending", h.findPendingTransactionHistories).Methods("GET")

	jobRoutes := swapRoutes.PathPrefix("/job").Subrouter()
	jobRoutes.Use(h.MiddleWare.SwapAuthorizationJobFunc)
	jobRoutes.HandleFunc("/update-ido", h.swapJobUpdateIdoStatus).Methods("GET")
	jobRoutes.HandleFunc("/update-pair", h.jobUpdateDataSwapPair).Methods("GET")
	jobRoutes.HandleFunc("/update-token", h.jobUpdateDataToken).Methods("GET")
	jobRoutes.HandleFunc("/claim-test-mainnet", h.gmPaymentClaimTestMainnet).Methods("GET")

	swapTokensRoutes := swapRoutes.PathPrefix("/token").Subrouter()
	swapTokensRoutes.HandleFunc("/list", h.getTokensInPool).Methods("GET")
	swapTokensRoutes.HandleFunc("/route", h.getRoutePair).Methods("GET")
	swapTokensRoutes.HandleFunc("/report", h.getTokensReport).Methods("GET")
	swapTokensRoutes.HandleFunc("/list/v1", h.getTokensInPoolV1).Methods("GET")
	swapTokensRoutes.HandleFunc("/route/v1", h.getRoutePairV1).Methods("GET")
	// swapTokensRoutes.HandleFunc("/report/v1", h.getTokensReport).Methods("GET")
	swapTokensRoutes.HandleFunc("/price", h.getTokensPrice).Methods("GET")
	swapTokensRoutes.HandleFunc("/summary", h.getTokenSummary).Methods("GET")

	userRoutes := swapRoutes.PathPrefix("/user").Subrouter()
	userRoutes.HandleFunc("/trade-histories", h.findUserSwapHistories).Methods("GET")

	swapPairRoutes := swapRoutes.PathPrefix("/pair").Subrouter()
	swapPairRoutes.HandleFunc("/list", h.findSwapPairs).Methods("GET")
	swapPairRoutes.HandleFunc("/trade-histories", h.findSwapHistories).Methods("GET")
	swapPairRoutes.HandleFunc("/apr", h.getLiquidityApr).Methods("GET")
	swapPairRoutes.HandleFunc("/apr/list", h.getListLiquidityAprReport).Methods("GET")

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

	// walletRoutes := swapRoutes.PathPrefix("/wallet").Subrouter()
	// walletRoutes.Use(h.MiddleWare.SwapAuthorizationJobFunc)
	// walletRoutes.HandleFunc("/update", h.addOrUpdateSwapWallet).Methods("PUT")
	// walletRoutes.HandleFunc("/detail", h.getSwapWallet).Methods("GET")

	gmRoutes := swapRoutes.PathPrefix("/gm").Subrouter()
	gmRoutes.Use(h.MiddleWare.SwapRecaptchaV2Middleware)
	gmRoutes.HandleFunc("/claim", h.gmPaymentClaim).Methods("GET")

	// evm bytescode check
	evmRoutes := api.PathPrefix("/evm").Subrouter()
	evmRoutes.HandleFunc("/bytescode", h.checkEvmBytescode).Methods("POST")

	//marketplace
	marketplace := api.PathPrefix("/marketplace").Subrouter()
	marketplace.HandleFunc("/collections", h.mkpCollections).Methods("GET")
	marketplace.HandleFunc("/collections/{contract_address}", h.mkpCollectionDetail).Methods("GET")
	marketplace.HandleFunc("/collections/{contract_address}/activities", h.getCollectionActivities).Methods("GET")
	marketplace.HandleFunc("/collections/{contract_address}/attributes", h.mkpCollectionAttributes).Methods("GET")
	marketplace.HandleFunc("/collections/{contract_address}/chart", h.getCollectionChart).Methods("GET")
	marketplace.HandleFunc("/collections/{contract_address}/nfts", h.mkplaceNftsOfACollection).Methods("GET")
	marketplace.HandleFunc("/collections/{contract_address}/nft-owners", h.mkplaceNftOwnerCollection).Methods("GET")
	marketplace.HandleFunc("/nfts", h.mkplaceNfts).Methods("GET")
	marketplace.HandleFunc("/collections/{contract_address}/nfts/{token_id}", h.mkplaceNftDetail).Methods("GET")
	marketplace.HandleFunc("/listing/{contract_address}/token/{token_id}", h.getListingViaGenAddressTokenID).Methods("GET")
	marketplace.HandleFunc("/offers/{contract_address}/token/{token_id}", h.getOfferViaGenAddressTokenID).Methods("GET")
	marketplace.HandleFunc("/wallet/{wallet_address}/listing", h.getListingOfAProfile).Methods("GET")
	marketplace.HandleFunc("/wallet/{wallet_address}/offer", h.getOffersOfAProfile).Methods("GET")
	marketplace.HandleFunc("/contract/{contract_address}/token/{token_id}/activities", h.getTokenActivities).Methods("GET")
	marketplace.HandleFunc("/contract/{contract_address}/token/{token_id}/soul_histories", h.getSoulHistories).Methods("GET")

	soul := api.PathPrefix("/soul").Subrouter()
	soul.HandleFunc("/signature", h.SoulCreateSignature).Methods("POST")
	soul.HandleFunc("/capture", h.SoulCaptureImage).Methods("POST")
	soul.HandleFunc("/nfts", h.soulNfts).Methods("GET")
	soul.HandleFunc("/nfts/{token_id}", h.soulNft).Methods("GET")

	api.HandleFunc("/time", h.serverTime).Methods("GET")

	//dao
	dao := api.PathPrefix("/dao").Subrouter()
	dao.HandleFunc("/proposals", h.proposals).Methods("GET")
	dao.HandleFunc("/proposals", h.createDraftProposals).Methods("POST")
	dao.HandleFunc("/proposals/{proposal_id}", h.getProposal).Methods("GET")
	dao.HandleFunc("/proposals/{proposal_id}/votes", h.getProposalVotes).Methods("GET")
	dao.HandleFunc("/proposals/{id}/{proposal_id}", h.mapOffAndOnChainProposal).Methods("PUT")

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
