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
        "/exists": {
            "post": {
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Проверяет есть ли путь(файл/папка)",
                "parameters": [
                    {
                        "description": "Путь до директории",
                        "name": "path",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/mkdir": {
            "post": {
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "text/plain"
                ],
                "summary": "Создает папку по указанному пути",
                "parameters": [
                    {
                        "description": "Путь до директории",
                        "name": "path",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/mv": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "summary": "Перемещает по указанному пути",
                "parameters": [
                    {
                        "description": "Список файлов",
                        "name": "files",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/handlers.FilesToMove"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/pwd": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Список файлов/папок из директории",
                "parameters": [
                    {
                        "description": "Путь до директории",
                        "name": "path",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/handlers.File"
                            }
                        }
                    }
                }
            }
        },
        "/rm": {
            "post": {
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Удаляет папку/файл по переданному пути",
                "parameters": [
                    {
                        "description": "Путь до директории",
                        "name": "path",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.File": {
            "type": "object",
            "properties": {
                "isdir": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "path": {
                    "description": "Type string ` + "`" + `json:\"type\"` + "`" + `",
                    "type": "string"
                }
            }
        },
        "handlers.FilesToMove": {
            "type": "object",
            "properties": {
                "files": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "to": {
                    "type": "string"
                }
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
