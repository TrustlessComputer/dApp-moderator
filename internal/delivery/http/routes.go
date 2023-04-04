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

	//admin
	admin := api.PathPrefix("/admin").Subrouter()
	admin.HandleFunc("/redis", h.getRedisKeys).Methods("GET")
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
