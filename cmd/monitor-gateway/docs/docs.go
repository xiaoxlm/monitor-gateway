// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://www.swagger.io/support",
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
        "/monitor-gateway/api/v1/metrics/batch-query": {
            "post": {
                "description": "更具PromQL查询指标",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BatchQuery"
                ],
                "summary": "BatchQuery",
                "operationId": "BatchQuery",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization Basic token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_xiaoxlm_monitor-gateway_api_request.MetricsBatchQueryBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {}
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrorRESP"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrorRESP"
                        }
                    }
                }
            }
        },
        "/monitor-gateway/api/v1/metrics/mapping": {
            "get": {
                "description": "获取指标映射",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ListMetricsMapping"
                ],
                "summary": "ListMetricsMapping",
                "operationId": "ListMetricsMapping",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization Basic token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/internal_model.MetricsMapping"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrorRESP"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "datatypes.JSONMap": {
            "type": "object",
            "additionalProperties": true
        },
        "github_com_xiaoxlm_monitor-gateway_api_request.MetricsBatchQueryBody": {
            "type": "object",
            "required": [
                "queries"
            ],
            "properties": {
                "queries": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_xiaoxlm_monitor-gateway_pkg_metrics_interface.QueryFormItem"
                    }
                }
            }
        },
        "github_com_xiaoxlm_monitor-gateway_pkg_metrics_interface.QueryFormItem": {
            "type": "object",
            "required": [
                "end",
                "query",
                "start",
                "step"
            ],
            "properties": {
                "end": {
                    "description": "结束时间",
                    "type": "integer"
                },
                "query": {
                    "description": "查询语句",
                    "type": "string"
                },
                "start": {
                    "description": "开始时间",
                    "type": "integer"
                },
                "step": {
                    "description": "步长",
                    "type": "integer"
                }
            }
        },
        "httputil.ErrorRESP": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string"
                },
                "type": {
                    "type": "integer"
                }
            }
        },
        "internal_model.MetricsMapping": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/mysql.DeletedTime"
                },
                "description": {
                    "description": "描述",
                    "type": "string"
                },
                "expression": {
                    "description": "表达式",
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "labels": {
                    "description": "指标标签",
                    "allOf": [
                        {
                            "$ref": "#/definitions/datatypes.JSONMap"
                        }
                    ]
                },
                "metricUniqueID": {
                    "description": "告警唯一标识",
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "mysql.DeletedTime": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
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
	Title:            "monitor-gateway",
	Description:      "This is a monitor gateway",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
