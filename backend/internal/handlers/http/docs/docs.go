// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth3/change": {
            "post": {
                "description": "Change password",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Change password",
                "parameters": [
                    {
                        "description": "Login payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/UserLoginChangeDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ResponseDTO-UserToken"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/auth3/forgot": {
            "post": {
                "description": "Forgot password",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Forgot password",
                "parameters": [
                    {
                        "description": "Login payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/UserLoginForgotDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ResponseDTO-CodeDeliveryDetails"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/auth3/login": {
            "post": {
                "description": "User login",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "Login payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/UserLoginDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ResponseDTO-UserToken"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/auth3/status": {
            "get": {
                "description": "healthcheck",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "healthcheck",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/EmptyResponseDTO"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "CodeDeliveryDetails": {
            "type": "object",
            "properties": {
                "deliveryMedium": {
                    "type": "string"
                }
            }
        },
        "EmptyResponseDTO": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "message": {
                    "type": "string",
                    "example": "Operação realizada com sucesso"
                },
                "status": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "ResponseDTO-CodeDeliveryDetails": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "data": {
                    "$ref": "#/definitions/CodeDeliveryDetails"
                },
                "message": {
                    "type": "string",
                    "example": "Operação realizada com sucesso"
                },
                "status": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "ResponseDTO-UserToken": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "data": {
                    "$ref": "#/definitions/UserToken"
                },
                "message": {
                    "type": "string",
                    "example": "Operação realizada com sucesso"
                },
                "status": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "UserLoginChangeDTO": {
            "type": "object",
            "required": [
                "email",
                "newPassword",
                "passwordResetCode"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "newPassword": {
                    "type": "string",
                    "minLength": 6
                },
                "passwordResetCode": {
                    "type": "string"
                }
            }
        },
        "UserLoginDTO": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                }
            }
        },
        "UserLoginForgotDTO": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "UserToken": {
            "type": "object",
            "properties": {
                "AccessToken": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Auth3 Rest API",
	Description:      "Auth3 Rest API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
