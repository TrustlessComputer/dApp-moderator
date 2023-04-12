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

	//profile
	profile := api.PathPrefix("/profile").Subrouter()
	profile.Use(h.MiddleWare.AuthorizationFunc)
	profile.HandleFunc("/me", h.currentUerProfile).Methods("GET")
	//profile.HandleFunc("/me/collections", h.currentUerProfile).Methods("GET")
	profile.HandleFunc("/wallet/{walletAddress}", h.profileByWallet).Methods("GET")
	profile.HandleFunc("/histories", h.createProfileHistory).Methods("POST")
	profile.HandleFunc("/histories/{txHash}/confirm", h.confirmProfileHistory).Methods("PUT")

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
