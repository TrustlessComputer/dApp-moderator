// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/nonce": {
            "post": {
                "description": "Generate a message for user's wallet",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Generate a message",
                "parameters": [
                    {
                        "description": "Generate message request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structure.GenerateMessage"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResponse"
                        }
                    }
                }
            }
        },
        "/auth/nonce/verify": {
            "post": {
                "description": "Verified the generated message",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Verified the generated message",
                "parameters": [
                    {
                        "description": "Verify message request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structure.VerifyMessage"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResponse"
                        }
                    }
                }
            }
        },
        "/bfs-service/browse/{walletAddress}": {
            "get": {
                "description": "Browse files of a wallet (uploader's wallet address)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BFS-service"
                ],
                "summary": "Browse files of a wallet",
                "parameters": [
                    {
                        "type": "string",
                        "description": "walletAddress",
                        "name": "walletAddress",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "path",
                        "name": "path",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResponse"
                        }
                    }
                }
            }
        },
        "/bfs-service/content/{walletAddress}": {
            "get": {
                "description": "Get file content of a wallet address (uploader's wallet address)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BFS-service"
                ],
                "summary": "Get content file",
                "parameters": [
                    {
                        "type": "string",
                        "description": "walletAddress",
                        "name": "walletAddress",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "path",
                        "name": "path",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResponse"
                        }
                    }
                }
            }
        },
        "/bfs-service/files/{walletAddress}": {
            "get": {
                "description": "Get files of a wallet (uploader's wallet address)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BFS-service"
                ],
                "summary": "Get files of a wallet",
                "parameters": [
                    {
                        "type": "string",
                        "description": "walletAddress",
                        "name": "walletAddress",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResponse"
                        }
                    }
                }
            }
        },
        "/bfs-service/info/{walletAddress}": {
            "get": {
                "description": "Get file info of a wallet address (uploader's wallet address)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BFS-service"
                ],
                "summary": "Get file info",
                "parameters": [
                    {
                        "type": "string",
                        "description": "walletAddress",
                        "name": "walletAddress",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "path",
                        "name": "path",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResponse"
                        }
                    }
                }
            }
        },
        "/bns-explorer/bns": {
            "get": {
                "description": "get Bns",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BNS-explorer"
                ],
                "summary": "get Bns",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResponse"
                        }
                    }
                }
            }
        },
        "/nft-explorer/collections": {
            "get": {
                "description": "Get Collections",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "nft-explorer"
                ],
                "summary": "Get Collections",
                "parameters": [
                    {
                        "type": "string",
                        "description": "owner",
                        "name": "owner",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "contract",
                        "name": "contract",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "default deployed_at_block",
                        "name": "sort_by",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "default -1",
                        "name": "sort",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResponse"
                        }
                    }
                }
            }
        },
        "/nft-explorer/collections/{contractAddress}": {
            "get": {
                "description": "Get Collections",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "nft-explorer"
                ],
                "summary": "Get Collections",
                "parameters": [
                    {
                        "type": "string",
                        "description": "contractAddress",
                        "name": "contractAddress",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResponse"
                        }
                    }
                }
            }
        },
        "/nft-explorer/collections/{contractAddress}/nfts": {
            "get": {
                "description": "Get nfts of a Collectionc",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "nft-explorer"
                ],
                "summary": "Get nfts of a Collectionc",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "contractAddress",
                        "name": "contractAddress",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResponse"
                        }
                    }
                }
            }
        },
        "/nft-explorer/collections/{contractAddress}/nfts/{tokenID}": {
            "get": {
                "description": "Get nft detail of a Collection",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "nft-explorer"
                ],
                "summary": "Get nft detail of a Collection",
                "parameters": [
                    {
                        "type": "string",
                        "description": "contractAddress",
                        "name": "contractAddress",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "tokenID",
                        "name": "tokenID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResponse"
                        }
                    }
                }
            }
        },
        "/nft-explorer/collections/{contractAddress}/nfts/{tokenID}/content": {
            "get": {
                "description": "Get nft content of a Collection",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "nft-explorer"
                ],
                "summary": "Get nft content of a Collection",
                "parameters": [
                    {
                        "type": "string",
                        "description": "contractAddress",
                        "name": "contractAddress",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "tokenID",
                        "name": "tokenID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResponse"
                        }
                    }
                }
            }
        },
        "/nft-explorer/nfts": {
            "get": {
                "description": "Get Nfts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "nft-explorer"
                ],
                "summary": "Get Nfts",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResponse"
                        }
                    }
                }
            }
        },
        "/nft-explorer/owner-address/{ownerAddress}/nfts": {
            "get": {
                "description": "Get nfts of a wallet address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "nft-explorer"
                ],
                "summary": "Get nfts of a wallet address",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "ownerAddress",
                        "name": "ownerAddress",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResponse"
                        }
                    }
                }
            }
        },
        "/quicknode/address/{walletAddress}/balance": {
            "get": {
                "description": "getaddress balance",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "QuickNode"
                ],
                "summary": "qn_addressBalance RPC Method",
                "parameters": [
                    {
                        "type": "string",
                        "description": "BTC walletAddress",
                        "name": "walletAddress",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.JsonResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/response.RedisResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/token-explorer/search": {
            "get": {
                "description": "search tokens",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "token-explorer"
                ],
                "summary": "search tokens",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "searching key",
                        "name": "key",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResponse"
                        }
                    }
                }
            }
        },
        "/token-explorer/token/{address}": {
            "get": {
                "description": "Get token detail",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "token-explorer"
                ],
                "summary": "Get token detail",
                "parameters": [
                    {
                        "type": "string",
                        "description": "contractAddress",
                        "name": "address",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResponse"
                        }
                    }
                }
            }
        },
        "/token-explorer/tokens": {
            "get": {
                "description": "Get tokens",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "token-explorer"
                ],
                "summary": "Get tokens",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResponse"
                        }
                    }
                }
            }
        },
        "/wallets/{walletAddress}": {
            "get": {
                "description": "Get Wallet's info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wallet"
                ],
                "summary": "Get Wallet's info",
                "parameters": [
                    {
                        "type": "string",
                        "description": "walletAddress",
                        "name": "walletAddress",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "response.JsonResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {
                    "$ref": "#/definitions/response.RespondErr"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "response.RedisResponse": {
            "type": "object",
            "properties": {
                "value": {
                    "type": "string"
                }
            }
        },
        "response.RespondErr": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "structure.GenerateMessage": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "walletType": {
                    "type": "string"
                }
            }
        },
        "structure.VerifyMessage": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "addressBTC": {
                    "description": "taproot address",
                    "type": "string"
                },
                "addressBTCSegwit": {
                    "type": "string"
                },
                "addressPayment": {
                    "type": "string"
                },
                "ethsignature": {
                    "type": "string"
                },
                "messagePrefix": {
                    "type": "string"
                },
                "signature": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0.0",
	Host:        "",
	BasePath:    "/dapp-moderator/v1",
	Schemes:     []string{},
	Title:       "tcDAPP APIs",
	Description: "This is a sample server TC-DAPP server.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
