// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/index": {
            "get": {
                "responses": {}
            }
        },
        "/user/createUser": {
            "post": {
                "responses": {
                    "200": {
                        "description": "code\",\"message\",\"data\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/deleteUser": {
            "put": {
                "responses": {
                    "200": {
                        "description": "code\",\"message\",\"data\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "responses": {
                    "200": {
                        "description": "code\",\"message\",\"data\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/getUser.email": {
            "get": {
                "responses": {
                    "200": {
                        "description": "code\",\"message\",\"data\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/getUser.id": {
            "get": {
                "responses": {
                    "200": {
                        "description": "code\",\"message\",\"data\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/getUser.name": {
            "get": {
                "responses": {
                    "200": {
                        "description": "code\",\"message\",\"data\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/getUser.phone": {
            "get": {
                "responses": {
                    "200": {
                        "description": "code\",\"message\",\"data\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/getUserList": {
            "get": {
                "responses": {
                    "200": {
                        "description": "code\",\"message\",\"data\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "responses": {
                    "200": {
                        "description": "code\",\"message\",\"data\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/ws/sendMsg": {
            "get": {
                "responses": {}
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
