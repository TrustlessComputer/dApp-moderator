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
	//api
	api := h.Handler.PathPrefix("/dapp/api").Subrouter()

	//quicknode
	quicknode := api.PathPrefix("/quicknode").Subrouter()
	quicknode.HandleFunc("/address/{walletAddress}/balance", h.addressBalance).Methods("GET")
	
	//quicknode
	nftExplorer := api.PathPrefix("/nft-explorer").Subrouter()
	nftExplorer.HandleFunc("/collections", h.collections).Methods("GET")
	nftExplorer.HandleFunc("/collections/{contractAddress}", h.collectionDetail).Methods("GET")
	nftExplorer.HandleFunc("/collections/{contractAddress}/nfts", h.collectionNfts).Methods("GET")
	nftExplorer.HandleFunc("/collections/{contractAddress}/nfts/{tokenID}", h.collectionNftDetail).Methods("GET")
	nftExplorer.HandleFunc("/collections/{contractAddress}/nfts/{tokenID}/content", h.collectionNftContent).Methods("GET")
	nftExplorer.HandleFunc("/nfts", h.nfts).Methods("GET")
	nftExplorer.HandleFunc("/owner-address/{walletAddress}/nfts", h.nftByWalletAddress).Methods("GET")

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
