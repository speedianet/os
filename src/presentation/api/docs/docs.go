// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://speedia.net/tos/",
        "contact": {
            "name": "Speedia Engineering",
            "url": "https://speedia.net/",
            "email": "eng+swagger@speedia.net"
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
        "/auth/login/": {
            "post": {
                "description": "Generate JWT with credentials",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "GenerateJwtWithCredentials",
                "parameters": [
                    {
                        "description": "Login",
                        "name": "loginDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.AccessToken"
                        }
                    }
                }
            }
        },
        "/database/{dbType}/": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "List databases names, users and sizes.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "database"
                ],
                "summary": "GetDatabases",
                "parameters": [
                    {
                        "enum": [
                            "mysql",
                            "postgres"
                        ],
                        "type": "string",
                        "description": "DatabaseType",
                        "name": "dbType",
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
                                "$ref": "#/definitions/entity.Database"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Add a new database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "database"
                ],
                "summary": "AddDatabase",
                "parameters": [
                    {
                        "enum": [
                            "mysql",
                            "postgres"
                        ],
                        "type": "string",
                        "description": "DatabaseType",
                        "name": "dbType",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "AddDatabase",
                        "name": "addDatabaseDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AddDatabase"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "DatabaseAdded",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/o11y/overview/": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Show system information and resource usage.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "o11y"
                ],
                "summary": "O11yOverview",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.O11yOverview"
                        }
                    }
                }
            }
        },
        "/services/": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "List services and their status.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "services"
                ],
                "summary": "GetServices",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Service"
                            }
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Start, stop, install or uninstall a service.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "services"
                ],
                "summary": "UpdateServiceStatus",
                "parameters": [
                    {
                        "description": "UpdateServiceStatusDetails",
                        "name": "updateSvcStatusDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateSvcStatus"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ServiceStatusUpdated",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/user/": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Update an user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "UpdateUser",
                "parameters": [
                    {
                        "description": "UpdateUserDetails",
                        "name": "updateUserDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "UserUpdated message or NewKeyString",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Add a new user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "AddNewUser",
                "parameters": [
                    {
                        "description": "NewUserDetails",
                        "name": "addUserDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AddUser"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "UserCreated",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/user/{userId}/": {
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Delete an user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "DeleteUser",
                "parameters": [
                    {
                        "type": "string",
                        "description": "UserId",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "UserDeleted",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.AddDatabase": {
            "type": "object",
            "properties": {
                "dbName": {
                    "type": "string"
                }
            }
        },
        "dto.AddUser": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.Login": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.UpdateSvcStatus": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/valueObject.ServiceStatus"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "dto.UpdateUser": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "shouldUpdateApiKey": {
                    "type": "boolean"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "entity.AccessToken": {
            "type": "object",
            "properties": {
                "expiresIn": {
                    "type": "integer"
                },
                "tokenStr": {
                    "type": "string"
                },
                "type": {
                    "$ref": "#/definitions/valueObject.AccessTokenType"
                }
            }
        },
        "entity.Database": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "size": {
                    "type": "integer"
                },
                "type": {
                    "$ref": "#/definitions/valueObject.DatabaseType"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.DatabaseUser"
                    }
                }
            }
        },
        "entity.DatabaseUser": {
            "type": "object",
            "properties": {
                "dbName": {
                    "type": "string"
                },
                "dbType": {
                    "$ref": "#/definitions/valueObject.DatabaseType"
                },
                "privileges": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "entity.O11yOverview": {
            "type": "object",
            "properties": {
                "currentUsage": {
                    "$ref": "#/definitions/valueObject.CurrentResourceUsage"
                },
                "hostname": {
                    "type": "string"
                },
                "publicIp": {
                    "type": "string"
                },
                "runtimeContext": {
                    "$ref": "#/definitions/valueObject.RuntimeContext"
                },
                "specs": {
                    "$ref": "#/definitions/valueObject.HardwareSpecs"
                },
                "uptimeSecs": {
                    "type": "integer"
                }
            }
        },
        "entity.Service": {
            "type": "object",
            "properties": {
                "cpuUsagePercent": {
                    "type": "number"
                },
                "memUsagePercent": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                },
                "pids": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "status": {
                    "$ref": "#/definitions/valueObject.ServiceStatus"
                },
                "uptimeSecs": {
                    "type": "number"
                }
            }
        },
        "valueObject.AccessTokenType": {
            "type": "string",
            "enum": [
                "sessionToken",
                "userApiKey"
            ],
            "x-enum-varnames": [
                "sessionToken",
                "userApiKey"
            ]
        },
        "valueObject.CurrentResourceUsage": {
            "type": "object",
            "properties": {
                "cpuUsagePercent": {
                    "type": "number"
                },
                "memUsagePercent": {
                    "type": "number"
                },
                "storageUsage": {
                    "type": "number"
                }
            }
        },
        "valueObject.DatabaseType": {
            "type": "string",
            "enum": [
                "mysql",
                "postgres"
            ],
            "x-enum-varnames": [
                "mysql",
                "postgres"
            ]
        },
        "valueObject.HardwareSpecs": {
            "type": "object",
            "properties": {
                "cpuCores": {
                    "type": "integer"
                },
                "cpuFrequency": {
                    "type": "number"
                },
                "cpuModel": {
                    "type": "string"
                },
                "memoryTotal": {
                    "type": "integer"
                },
                "storageTotal": {
                    "type": "integer"
                }
            }
        },
        "valueObject.RuntimeContext": {
            "type": "string",
            "enum": [
                "container",
                "vm",
                "bareMetal"
            ],
            "x-enum-varnames": [
                "container",
                "vm",
                "bareMetal"
            ]
        },
        "valueObject.ServiceStatus": {
            "type": "string",
            "enum": [
                "running",
                "stopped",
                "uninstalled",
                "installed"
            ],
            "x-enum-varnames": [
                "running",
                "stopped",
                "uninstalled",
                "installed"
            ]
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "Type \"Bearer\" + JWT token or API key.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "localhost:10000",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "SamApi",
	Description:      "Speedia AppManager API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
