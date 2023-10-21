// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "No terms",
        "contact": {
            "name": "Support",
            "email": "florian.charpentier67@gmail.com"
        },
        "license": {
            "name": "MIT",
            "url": "http://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/download/url": {
            "get": {
                "description": "Execute the Youtube-DL job using Cloud Task",
                "consumes": [
                    "application/json"
                ],
                "summary": "Download and save a new music file",
                "parameters": [
                    {
                        "description": "Parameters to send to job",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/download.downloadPayload"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            }
        },
        "/music-researcher/genres": {
            "get": {
                "description": "List available genre in Spotify API",
                "consumes": [
                    "application/json"
                ],
                "summary": "List genres",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/music-researcher/search": {
            "get": {
                "description": "Searchs for music in Spotify API",
                "consumes": [
                    "application/json"
                ],
                "summary": "Music search",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Main user query",
                        "name": "q",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "Genre list",
                        "name": "genre",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Limit result count",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "download.downloadPayload": {
            "type": "object",
            "properties": {
                "meta": {
                    "description": "metadata infos to format the file",
                    "type": "object",
                    "properties": {
                        "album": {
                            "description": "the album",
                            "type": "string"
                        },
                        "artist": {
                            "description": "the artist",
                            "type": "string"
                        },
                        "track": {
                            "description": "the music track",
                            "type": "string"
                        }
                    }
                },
                "url": {
                    "description": "the youtube url to use for download",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "api.dadard.fr",
	BasePath:         "/",
	Schemes:          []string{"https"},
	Title:            "Gateway Front API",
	Description:      "This application provides a front gateway allowing you\nto interact with multiple GRPC microservices hosted\nin Google Cloud",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
