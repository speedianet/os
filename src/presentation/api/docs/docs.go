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
        "/account/": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "List accounts.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "GetAccounts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Account"
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
                "description": "Update an account.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "UpdateAccount",
                "parameters": [
                    {
                        "description": "UpdateAccount",
                        "name": "updateAccountDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateAccount"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "AccountUpdated message or NewKeyString",
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
                "description": "Add a new account.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "AddNewAccount",
                "parameters": [
                    {
                        "description": "NewAccount",
                        "name": "addAccountDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AddAccount"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "AccountCreated",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/account/{accountId}/": {
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Delete an account.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "DeleteAccount",
                "parameters": [
                    {
                        "type": "string",
                        "description": "AccountId",
                        "name": "accountId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "AccountDeleted",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
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
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
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
        "/database/{dbType}/{dbName}/": {
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Delete a database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "database"
                ],
                "summary": "DeleteDatabase",
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
                        "type": "string",
                        "description": "DatabaseName",
                        "name": "dbName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "DatabaseDeleted",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/database/{dbType}/{dbName}/user/": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Add a new database user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "database"
                ],
                "summary": "AddDatabaseUser",
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
                        "type": "string",
                        "description": "DatabaseName",
                        "name": "dbName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "AddDatabaseUser",
                        "name": "addDatabaseUserDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AddDatabaseUser"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "DatabaseUserAdded",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/database/{dbType}/{dbName}/user/{dbUser}/": {
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Delete a database user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "database"
                ],
                "summary": "DeleteDatabaseUser",
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
                        "type": "string",
                        "description": "DatabaseName",
                        "name": "dbName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "DatabaseUsername",
                        "name": "dbUser",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "DatabaseUserDeleted",
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
        "/runtime/php/{hostname}/": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get php version, modules and settings for a hostname.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "runtime"
                ],
                "summary": "GetPhpConfigs",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Hostname",
                        "name": "hostname",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.PhpConfigs"
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
                "description": "Update php version, modules and settings for a hostname.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "runtime"
                ],
                "summary": "UpdatePhpConfigs",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Hostname",
                        "name": "hostname",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "UpdatePhpConfigs",
                        "name": "updatePhpConfigsDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdatePhpConfigs"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "PhpConfigsUpdated",
                        "schema": {
                            "type": "object"
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
        }
    },
    "definitions": {
        "dto.AddAccount": {
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
        "dto.AddDatabase": {
            "type": "object",
            "properties": {
                "dbName": {
                    "type": "string"
                }
            }
        },
        "dto.AddDatabaseUser": {
            "type": "object",
            "properties": {
                "dbName": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
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
        "dto.UpdateAccount": {
            "type": "object",
            "properties": {
                "accountId": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "shouldUpdateApiKey": {
                    "type": "boolean"
                }
            }
        },
        "dto.UpdatePhpConfigs": {
            "type": "object",
            "properties": {
                "hostname": {
                    "type": "string"
                },
                "modules": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.PhpModule"
                    }
                },
                "settings": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.PhpSetting"
                    }
                },
                "version": {
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
        "entity.Account": {
            "type": "object",
            "properties": {
                "groupId": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
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
        "entity.PhpConfigs": {
            "type": "object",
            "properties": {
                "hostname": {
                    "type": "string"
                },
                "modules": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.PhpModule"
                    }
                },
                "settings": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.PhpSetting"
                    }
                },
                "version": {
                    "$ref": "#/definitions/entity.PhpVersion"
                }
            }
        },
        "entity.PhpModule": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "entity.PhpSetting": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "options": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "entity.PhpVersion": {
            "type": "object",
            "properties": {
                "options": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "value": {
                    "type": "string"
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
                "accountApiKey"
            ],
            "x-enum-varnames": [
                "sessionToken",
                "accountApiKey"
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
