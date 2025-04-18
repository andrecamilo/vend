// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
            "name": "API Support",
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
        "/contextos": {
            "get": {
                "description": "Retorna a lista de todos os contextos cadastrados",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "contextos"
                ],
                "summary": "Listar contextos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Contexto"
                            }
                        }
                    },
                    "501": {
                        "description": "Not Implemented",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Cria um novo contexto no sistema",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "contextos"
                ],
                "summary": "Criar contexto",
                "parameters": [
                    {
                        "description": "Dados do contexto",
                        "name": "contexto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Contexto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/domain.Contexto"
                        }
                    },
                    "501": {
                        "description": "Not Implemented",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/contextos/{id}": {
            "get": {
                "description": "Retorna os dados de um contexto específico",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "contextos"
                ],
                "summary": "Buscar contexto",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID do contexto",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Contexto"
                        }
                    },
                    "501": {
                        "description": "Not Implemented",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Atualiza os dados de um contexto específico",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "contextos"
                ],
                "summary": "Atualizar contexto",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID do contexto",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Dados do contexto",
                        "name": "contexto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Contexto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Contexto"
                        }
                    },
                    "501": {
                        "description": "Not Implemented",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "Remove um contexto do sistema",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "contextos"
                ],
                "summary": "Deletar contexto",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID do contexto",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "501": {
                        "description": "Not Implemented",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/pessoas": {
            "get": {
                "description": "Retorna a lista de todas as pessoas cadastradas",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pessoas"
                ],
                "summary": "Listar pessoas",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Pessoa"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Cria uma nova pessoa no sistema",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pessoas"
                ],
                "summary": "Criar pessoa",
                "parameters": [
                    {
                        "description": "Dados da pessoa",
                        "name": "pessoa",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Pessoa"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/domain.Pessoa"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/pessoas/{id}": {
            "get": {
                "description": "Retorna os dados de uma pessoa específica",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pessoas"
                ],
                "summary": "Buscar pessoa",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID da pessoa",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Pessoa"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Atualiza os dados de uma pessoa específica",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pessoas"
                ],
                "summary": "Atualizar pessoa",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID da pessoa",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Dados da pessoa",
                        "name": "pessoa",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Pessoa"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Pessoa"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "Remove uma pessoa do sistema",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pessoas"
                ],
                "summary": "Deletar pessoa",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID da pessoa",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/prompts": {
            "get": {
                "description": "Retorna a lista de todos os prompts cadastrados",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "prompts"
                ],
                "summary": "Listar prompts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Prompt"
                            }
                        }
                    },
                    "501": {
                        "description": "Not Implemented",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Cria um novo prompt no sistema",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "prompts"
                ],
                "summary": "Criar prompt",
                "parameters": [
                    {
                        "description": "Dados do prompt",
                        "name": "prompt",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Prompt"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/domain.Prompt"
                        }
                    },
                    "501": {
                        "description": "Not Implemented",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/prompts/{id}": {
            "get": {
                "description": "Retorna os dados de um prompt específico",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "prompts"
                ],
                "summary": "Buscar prompt",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID do prompt",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Prompt"
                        }
                    },
                    "501": {
                        "description": "Not Implemented",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Atualiza os dados de um prompt específico",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "prompts"
                ],
                "summary": "Atualizar prompt",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID do prompt",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Dados do prompt",
                        "name": "prompt",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Prompt"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Prompt"
                        }
                    },
                    "501": {
                        "description": "Not Implemented",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "Remove um prompt do sistema",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "prompts"
                ],
                "summary": "Deletar prompt",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID do prompt",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "501": {
                        "description": "Not Implemented",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/telefones": {
            "get": {
                "description": "Retorna a lista de todos os telefones cadastrados",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "telefones"
                ],
                "summary": "Listar telefones",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Telefone"
                            }
                        }
                    },
                    "501": {
                        "description": "Not Implemented",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Cria um novo telefone no sistema",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "telefones"
                ],
                "summary": "Criar telefone",
                "parameters": [
                    {
                        "description": "Dados do telefone",
                        "name": "telefone",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Telefone"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/domain.Telefone"
                        }
                    },
                    "501": {
                        "description": "Not Implemented",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/telefones/{id}": {
            "get": {
                "description": "Retorna os dados de um telefone específico",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "telefones"
                ],
                "summary": "Buscar telefone",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID do telefone",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Telefone"
                        }
                    },
                    "501": {
                        "description": "Not Implemented",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Atualiza os dados de um telefone específico",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "telefones"
                ],
                "summary": "Atualizar telefone",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID do telefone",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Dados do telefone",
                        "name": "telefone",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Telefone"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Telefone"
                        }
                    },
                    "501": {
                        "description": "Not Implemented",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "Remove um telefone do sistema",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "telefones"
                ],
                "summary": "Deletar telefone",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID do telefone",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "501": {
                        "description": "Not Implemented",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Contexto": {
            "type": "object",
            "required": [
                "nome"
            ],
            "properties": {
                "data_fim": {
                    "type": "string"
                },
                "data_inicio": {
                    "type": "string"
                },
                "descricao": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "nome": {
                    "type": "string"
                },
                "pessoas": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.Pessoa"
                    }
                },
                "prompts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.Prompt"
                    }
                }
            }
        },
        "domain.Pessoa": {
            "type": "object",
            "required": [
                "email",
                "nome"
            ],
            "properties": {
                "contextos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.Contexto"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "nome": {
                    "type": "string"
                },
                "telefones": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.Telefone"
                    }
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "domain.Prompt": {
            "type": "object",
            "required": [
                "conteudo"
            ],
            "properties": {
                "conteudo": {
                    "type": "string"
                },
                "contexto_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "domain.Telefone": {
            "type": "object",
            "required": [
                "numero",
                "tipo"
            ],
            "properties": {
                "id": {
                    "type": "string"
                },
                "numero": {
                    "type": "string"
                },
                "pessoa_id": {
                    "type": "string"
                },
                "tipo": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Vend API",
	Description:      "API para gerenciamento de vendas com integração ChatGPT",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
