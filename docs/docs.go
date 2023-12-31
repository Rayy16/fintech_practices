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
            "name": "rliu",
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
        "/download/public/{file_type}/{file_name}": {
            "get": {
                "description": "下载公共文件的统一接口，数字人、封面图片、素材库素材均通过本接口下载",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "*/*"
                ],
                "tags": [
                    "download"
                ],
                "summary": "下载公共文件接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "下载的文件类型, 类型为枚举值：dp、resource、cover_image",
                        "name": "file_type",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "下载的文件名称",
                        "name": "file_name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "file"
                        }
                    }
                }
            }
        },
        "/download/{file_type}/{file_name}": {
            "get": {
                "description": "下载文件的统一接口，数字人、封面图片、素材库素材均通过本接口下载",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "*/*"
                ],
                "tags": [
                    "download"
                ],
                "summary": "下载文件接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "下载的文件类型, 类型为枚举值：dp、resource、cover_image",
                        "name": "file_type",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "下载的文件名称",
                        "name": "file_name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "file"
                        }
                    }
                }
            }
        },
        "/dp": {
            "get": {
                "description": "查询用户所拥有的数字人信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "digital person"
                ],
                "summary": "查询数字人接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "分页查询页数，默认为1",
                        "name": "page_no",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "分页查询页大小，默认为10",
                        "name": "page_size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "查询返回的排序字段，默认为创建时间",
                        "name": "order_field",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "排序方式，默认为倒序",
                        "name": "method",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.GetDpResp"
                        }
                    }
                }
            },
            "post": {
                "description": "创建用户数字人信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "digital person"
                ],
                "summary": "创建数字人接口",
                "parameters": [
                    {
                        "description": "数字人名称、形象link、音频link、音色link、文本内容。如传输音频link，则音色link与文本内容可为空字符串。如音频link为空字符串，则后二者必须传输",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.CreateDpReq"
                        }
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.CommResp"
                        }
                    }
                }
            }
        },
        "/dp/{dp_link}": {
            "delete": {
                "description": "删除用户所拥有的数字人信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "digital person"
                ],
                "summary": "删除数字人接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "需删除的数字人id",
                        "name": "dp_link",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.CommResp"
                        }
                    }
                }
            }
        },
        "/hotvedio": {
            "get": {
                "description": "获取首页数字人视频",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "*/*"
                ],
                "tags": [
                    "homepage"
                ],
                "summary": "首页视频接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "分页查询，页码",
                        "name": "pageNo",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "分页查询，页大小",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.GetDpResp"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "登录以获取token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "登录接口",
                "parameters": [
                    {
                        "description": "用户账号与加密的用户密码",
                        "name": "user_account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.AuthReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.AuthResp"
                        }
                    }
                }
            }
        },
        "/pubkey": {
            "get": {
                "description": "获取rsa公钥",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authorization"
                ],
                "summary": "获取公钥接口",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.PubKeyResp"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "注册用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "注册用户接口",
                "parameters": [
                    {
                        "description": "用户名称、用户账户与加密的用户密码",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.RegisterReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.CommResp"
                        }
                    }
                }
            }
        },
        "/resource": {
            "post": {
                "description": "创建用户素材库素材",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "resource lib"
                ],
                "summary": "创建素材库素材接口",
                "parameters": [
                    {
                        "description": "素材描述、素材类型(tone、image)， IsPng(如果是image类型，是否是图片形象素材)",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.CreateResourceReq"
                        }
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.CreateResourceResp"
                        }
                    }
                }
            }
        },
        "/resource/{resource_link}": {
            "delete": {
                "description": "删除用户所拥有的素材库素材",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "resource lib"
                ],
                "summary": "删除素材库素材接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "素材连接，tone or image",
                        "name": "resource_link",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.CommResp"
                        }
                    }
                }
            }
        },
        "/resource/{resource_type}": {
            "get": {
                "description": "查询用户所拥有的 or 公共的素材库素材信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "resource lib"
                ],
                "summary": "查询素材库接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "素材类型，tone or image",
                        "name": "resource_type",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "分页查询页数，默认为1",
                        "name": "page_no",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "分页查询页大小，默认为10",
                        "name": "page_size",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "是否查询公共素材，默认为否",
                        "name": "is_public",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.GetResourceResp"
                        }
                    }
                }
            }
        },
        "/upload/{file_type}/{file_name}": {
            "post": {
                "description": "上传文件的统一接口，数字人、封面图片、素材库素材均通过本接口下载",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "*/*"
                ],
                "tags": [
                    "download"
                ],
                "summary": "上传文件接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "上传的文件类型，类型为枚举值：audio、resource",
                        "name": "file_type",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "上传的文件名称，需要带上相应后缀，例如audio为.wav, resource为 .png 或 .mp4",
                        "name": "file_name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "上传的文件",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.CommResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "schema.AuthReq": {
            "type": "object",
            "required": [
                "decrypt_data",
                "user_account"
            ],
            "properties": {
                "decrypt_data": {
                    "type": "string"
                },
                "user_account": {
                    "type": "string"
                }
            }
        },
        "schema.AuthResp": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "msg": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "schema.CommResp": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "schema.CreateDpReq": {
            "type": "object",
            "required": [
                "audio_link",
                "content",
                "dp_name",
                "image_link",
                "tone_link"
            ],
            "properties": {
                "audio_link": {
                    "type": "string"
                },
                "content": {
                    "type": "string"
                },
                "dp_name": {
                    "type": "string"
                },
                "image_link": {
                    "type": "string"
                },
                "tone_link": {
                    "type": "string"
                }
            }
        },
        "schema.CreateResourceReq": {
            "type": "object",
            "required": [
                "resource_describe",
                "resource_type"
            ],
            "properties": {
                "is_png": {
                    "type": "boolean"
                },
                "resource_describe": {
                    "type": "string"
                },
                "resource_type": {
                    "type": "string"
                }
            }
        },
        "schema.CreateResourceResp": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "msg": {
                    "type": "string"
                },
                "resource_link": {
                    "type": "string"
                }
            }
        },
        "schema.DpEntity": {
            "type": "object",
            "properties": {
                "cover_image_link": {
                    "type": "string"
                },
                "create_time": {
                    "type": "string"
                },
                "dp_link": {
                    "type": "string"
                },
                "dp_name": {
                    "type": "string"
                },
                "dp_status": {
                    "type": "integer"
                },
                "hot_score": {
                    "type": "integer"
                },
                "owner": {
                    "type": "string"
                },
                "update_time": {
                    "type": "string"
                }
            }
        },
        "schema.GetDpResp": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.DpEntity"
                    }
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "schema.GetResourceResp": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.ResourceEntity"
                    }
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "schema.PubKeyResp": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "string"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "schema.RegisterReq": {
            "type": "object",
            "required": [
                "decrypt_data",
                "user_account",
                "user_name"
            ],
            "properties": {
                "decrypt_data": {
                    "type": "string"
                },
                "user_account": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "schema.ResourceEntity": {
            "type": "object",
            "properties": {
                "cover_image_link": {
                    "type": "string"
                },
                "create_time": {
                    "type": "string"
                },
                "resouce_id": {
                    "type": "string"
                },
                "resource_describe": {
                    "type": "string"
                },
                "resource_link": {
                    "type": "string"
                },
                "update_time": {
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
	BasePath:         "",
	Schemes:          []string{},
	Title:            "cc fintech practices API",
	Description:      "榕树平台API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
