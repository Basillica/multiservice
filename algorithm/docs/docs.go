// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "The Algorithm Service",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/json": {
            "get": {
                "description": "login for auth users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "auth login as a user",
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/types.ArrayRequest"
                            }
                        }
                    },
                    "401": {
                        "description": "unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "forbidden",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "types.ArrayRequest": {
            "type": "object",
            "required": [
                "array",
                "integer"
            ],
            "properties": {
                "array": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "integer": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "The Algorith Service api documentation",
	Description:      "The api documentation of the core apis of the algorithm service.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}