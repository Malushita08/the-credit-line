// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
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
        "/creditLines": {
            "post": {
                "description": "Create a creditLine",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "creditLine"
                ],
                "summary": "Create a creditLine",
                "parameters": [
                    {
                        "description": "creditLine Data",
                        "name": "creditLine",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreditLineRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseBody"
                        }
                    }
                }
            }
        },
        "/creditLines/foundingName/{foundingName}": {
            "get": {
                "description": "Get all the creditLines requests a foundingName did",
                "tags": [
                    "creditLine"
                ],
                "summary": "Get all the creditLines requests a foundingName did",
                "parameters": [
                    {
                        "type": "string",
                        "description": "creditLine foundingName",
                        "name": "foundingName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.CreditLine"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.CreditLine": {
            "type": "object",
            "properties": {
                "allowedRequest": {
                    "type": "boolean"
                },
                "attemptAcceptedNumber": {
                    "type": "integer"
                },
                "attemptNumber": {
                    "type": "integer"
                },
                "cashBalance": {
                    "type": "number"
                },
                "foundingName": {
                    "type": "string"
                },
                "foundingType": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastAcceptedRequestDate": {
                    "type": "string"
                },
                "monthlyRevenue": {
                    "type": "number"
                },
                "recommendedCreditLine": {
                    "type": "number"
                },
                "requestedCreditLine": {
                    "type": "number"
                },
                "requestedDate": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                }
            }
        },
        "models.CreditLineRequestBody": {
            "type": "object",
            "properties": {
                "cashBalance": {
                    "type": "number"
                },
                "foundingName": {
                    "type": "string"
                },
                "foundingType": {
                    "type": "string"
                },
                "monthlyRevenue": {
                    "type": "number"
                },
                "requestedCreditLine": {
                    "type": "number"
                }
            }
        },
        "models.CreditLineResponseBody": {
            "type": "object",
            "properties": {
                "cashBalance": {
                    "type": "number"
                },
                "foundingName": {
                    "type": "string"
                },
                "foundingType": {
                    "type": "string"
                },
                "monthlyRevenue": {
                    "type": "number"
                },
                "recommendedCreditLine": {
                    "type": "number"
                },
                "requestedCreditLine": {
                    "type": "number"
                },
                "requestedDate": {
                    "type": "string"
                }
            }
        },
        "models.ResponseBody": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.CreditLineResponseBody"
                },
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "The Credit Line API",
	Description:      "API that calculates a recommended creditLine for a Founding based on its type and other fields.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
