// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin/category": {
            "get": {
                "description": "分类列表，无分页，递归实现。",
                "consumes": [
                    "multipart/form-data",
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "分类管理"
                ],
                "summary": "分类列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.ApiResp"
                        }
                    },
                    "302": {
                        "description": "Found",
                        "schema": {
                            "$ref": "#/definitions/util.ApiResp"
                        }
                    }
                }
            },
            "put": {
                "description": "编辑分类",
                "consumes": [
                    "multipart/form-data",
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "分类管理"
                ],
                "summary": "分类更新",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "分类id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "父级id",
                        "name": "pid",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "排序，数值越大越靠前",
                        "name": "order",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "分类名称",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.ApiResp"
                        }
                    },
                    "302": {
                        "description": "Found",
                        "schema": {
                            "$ref": "#/definitions/util.ApiResp"
                        }
                    }
                }
            },
            "post": {
                "description": "新增分类",
                "consumes": [
                    "multipart/form-data",
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "分类管理"
                ],
                "summary": "分类新增",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "父级id",
                        "name": "pid",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "排序，数值越大越靠前",
                        "name": "order",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "分类名称",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.ApiResp"
                        }
                    },
                    "302": {
                        "description": "Found",
                        "schema": {
                            "$ref": "#/definitions/util.ApiResp"
                        }
                    }
                }
            },
            "delete": {
                "description": "删除分类 暂仅支持单个删除",
                "consumes": [
                    "multipart/form-data",
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "分类管理"
                ],
                "summary": "分类删除",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "分类id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.ApiResp"
                        }
                    },
                    "302": {
                        "description": "Found",
                        "schema": {
                            "$ref": "#/definitions/util.ApiResp"
                        }
                    }
                }
            }
        },
        "/admin/login": {
            "post": {
                "description": "统一登录接口，暂时接收账号密码登录和手机验证码登录",
                "consumes": [
                    "multipart/form-data",
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户信息"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "登录类型1:用户名密码登录，2手机号验证码登录",
                        "name": "type",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "登录邮箱",
                        "name": "email",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "手机号",
                        "name": "phone",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "手机验证码",
                        "name": "captcha",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.ApiResp"
                        }
                    },
                    "302": {
                        "description": "Found",
                        "schema": {
                            "$ref": "#/definitions/util.ApiResp"
                        }
                    }
                }
            }
        },
        "/admin/matieral/token": {
            "get": {
                "description": "上传图片获取七牛上传token",
                "consumes": [
                    "multipart/form-data",
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "素材管理"
                ],
                "summary": "获取上传token",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.ApiResp"
                        }
                    },
                    "302": {
                        "description": "Found",
                        "schema": {
                            "$ref": "#/definitions/util.ApiResp"
                        }
                    }
                }
            }
        },
        "/admin/tag": {
            "get": {
                "description": "标签列表。",
                "consumes": [
                    "multipart/form-data",
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "标签管理"
                ],
                "summary": "标签列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页显示条数",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "名称筛选",
                        "name": "title",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ListData"
                        }
                    },
                    "302": {
                        "description": "Found",
                        "schema": {
                            "$ref": "#/definitions/util.ApiResp"
                        }
                    }
                }
            },
            "put": {
                "description": "编辑分类",
                "consumes": [
                    "multipart/form-data",
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "分类管理"
                ],
                "summary": "分类更新",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "分类id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "父级id",
                        "name": "pid",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "排序，数值越大越靠前",
                        "name": "order",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "分类名称",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.ApiResp"
                        }
                    },
                    "302": {
                        "description": "Found",
                        "schema": {
                            "$ref": "#/definitions/util.ApiResp"
                        }
                    }
                }
            },
            "post": {
                "description": "新增标签",
                "consumes": [
                    "multipart/form-data",
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "标签管理"
                ],
                "summary": "标签新增",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "排序，数值越大越靠前",
                        "name": "order",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "标签名称",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.ApiResp"
                        }
                    },
                    "302": {
                        "description": "Found",
                        "schema": {
                            "$ref": "#/definitions/util.ApiResp"
                        }
                    }
                }
            },
            "delete": {
                "description": "删除标签 暂仅支持单个删除",
                "consumes": [
                    "multipart/form-data",
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "分类管理"
                ],
                "summary": "标签删除",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "标签id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.ApiResp"
                        }
                    },
                    "302": {
                        "description": "Found",
                        "schema": {
                            "$ref": "#/definitions/util.ApiResp"
                        }
                    }
                }
            }
        },
        "/admin/user": {
            "get": {
                "description": "管理员列表",
                "consumes": [
                    "multipart/form-data",
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户信息"
                ],
                "summary": "管理员列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页显示条数",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.ApiResp"
                        }
                    },
                    "302": {
                        "description": "Found",
                        "schema": {
                            "$ref": "#/definitions/util.ApiResp"
                        }
                    }
                }
            },
            "put": {
                "description": "编辑管理员",
                "consumes": [
                    "multipart/form-data",
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户信息"
                ],
                "summary": "管理员更新",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户昵称",
                        "name": "nickname",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "用户真实姓名-用于登录使用",
                        "name": "realname",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "手机号-可用于登录",
                        "name": "phone",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户头像",
                        "name": "avatar",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.ApiResp"
                        }
                    },
                    "302": {
                        "description": "Found",
                        "schema": {
                            "$ref": "#/definitions/util.ApiResp"
                        }
                    }
                }
            },
            "post": {
                "description": "新增管理员",
                "consumes": [
                    "multipart/form-data",
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户信息"
                ],
                "summary": "管理员新增",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户昵称",
                        "name": "nickname",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "用户真实姓名-用于登录使用",
                        "name": "realname",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "手机号-可用于登录",
                        "name": "phone",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户头像",
                        "name": "avatar",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.ApiResp"
                        }
                    },
                    "302": {
                        "description": "Found",
                        "schema": {
                            "$ref": "#/definitions/util.ApiResp"
                        }
                    }
                }
            },
            "delete": {
                "description": "删除管理员-暂只支持单个删除",
                "consumes": [
                    "multipart/form-data",
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户信息"
                ],
                "summary": "管理员删除",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.ApiResp"
                        }
                    },
                    "302": {
                        "description": "Found",
                        "schema": {
                            "$ref": "#/definitions/util.ApiResp"
                        }
                    }
                }
            }
        },
        "/admin/user/:id": {
            "get": {
                "description": "根据主键ID获取管理员详情",
                "consumes": [
                    "multipart/form-data",
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户信息"
                ],
                "summary": "管理员详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.ApiResp"
                        }
                    },
                    "302": {
                        "description": "Found",
                        "schema": {
                            "$ref": "#/definitions/util.ApiResp"
                        }
                    }
                }
            }
        },
        "/captcha": {
            "get": {
                "description": "获取图形验证码图片",
                "consumes": [
                    "multipart/form-data",
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公共接口"
                ],
                "summary": "图形验证码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "手机号",
                        "name": "phone",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.ApiResp"
                        }
                    },
                    "302": {
                        "description": "Found",
                        "schema": {
                            "$ref": "#/definitions/util.ApiResp"
                        }
                    }
                }
            }
        },
        "/sms": {
            "post": {
                "description": "发送短信接口",
                "consumes": [
                    "multipart/form-data",
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公共接口"
                ],
                "summary": "短信接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "手机号",
                        "name": "phone",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "图形验证码的值",
                        "name": "captcha",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.ApiResp"
                        }
                    },
                    "302": {
                        "description": "Found",
                        "schema": {
                            "$ref": "#/definitions/util.ApiResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.ListData": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "list": {
                    "type": "object"
                }
            }
        },
        "util.ApiResp": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "msg": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
	swag.Register(swag.Name, &s{})
}
