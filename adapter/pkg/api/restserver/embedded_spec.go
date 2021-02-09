// Code generated by go-swagger; DO NOT EDIT.

// Copyright (c) 2021, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package restserver

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json",
    "multipart/form-data"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This document specifies a **RESTful API** for WSO2 **API Microgateway** - Adapter.\n",
    "title": "WSO2 API Microgateway - Adapter",
    "contact": {
      "name": "WSO2",
      "url": "http://wso2.com/products/api-manager/",
      "email": "architecture@wso2.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "v1.2"
  },
  "host": "apis.wso2.com",
  "basePath": "/api/mgw/adapter/0.1",
  "paths": {
    "/apis": {
      "get": {
        "security": [
          {
            "BasicAuth": []
          }
        ],
        "description": "This operation can be used to retrieve meta info about all APIs\n",
        "tags": [
          "API (Collection)"
        ],
        "summary": "Get a list of API metadata",
        "parameters": [
          {
            "type": "string",
            "description": "Optional - Condition to filter APIs. Currently only filtering \nby API type (HTTP or WebSocket) is supported.\n\"http\" for HTTP type\n\"ws\" for WebSocket type\n",
            "name": "apiType",
            "in": "query"
          },
          {
            "type": "integer",
            "description": "Number of APIs (APIMeta objects to return)\n",
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "An array of API Metadata",
            "schema": {
              "$ref": "#/definitions/APIMeta"
            }
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        },
        "x-wso2-curl": "curl -k -H \"Authorization: Bearer ae4eae22-3f65-387b-a171-d37eaa366fa8\" \n-X GET \"https://127.0.0.1:9443/api/mgw/adapter/0.1/apis\"\n",
        "x-wso2-request": "GET https://127.0.0.1:9443/api/mgw/adapter/0.1/apis?apiType=http\nAuthorization: Bearer ae4eae22-3f65-387b-a171-d37eaa366fa8\n",
        "x-wso2-response": "HTTP/1.1 200 OK"
      },
      "post": {
        "security": [
          {
            "BasicAuth": []
          }
        ],
        "description": "This operation can be used to import an API.\n",
        "consumes": [
          "multipart/form-data"
        ],
        "tags": [
          "API (Individual)"
        ],
        "summary": "Import an API",
        "parameters": [
          {
            "type": "file",
            "x-exportParamName": "File",
            "description": "Zip archive consisting on exported api configuration\n",
            "name": "file",
            "in": "formData",
            "required": true
          },
          {
            "type": "boolean",
            "x-exportParamName": "PreserveProvider",
            "x-optionalDataType": "Bool",
            "description": "Preserve Original Provider of the API. This is the user choice to keep or replace the API provider.\n",
            "name": "preserveProvider",
            "in": "query"
          },
          {
            "type": "boolean",
            "x-exportParamName": "Overwrite",
            "x-optionalDataType": "Bool",
            "description": "Whether to update the API or not. This is used when updating already existing APIs.\n",
            "name": "overwrite",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful.\nAPI deployed or updated Successfully.\n",
            "schema": {
              "$ref": "#/definitions/DeployResponse"
            }
          },
          "403": {
            "description": "Forbidden.\nNot Authorized to deploy or update.\n",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Not Found. \nRequested API to update not found (when overwrite parameter is included).\n",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "409": {
            "description": "Conflict.\nAPI to import already exists (when overwrite parameter is not included).\n",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        },
        "x-wso2-curl": "curl -k -F \"file=@exported.zip\" -X POST -H \"Authorization: Bearer ae4eae22-3f65-387b-a171-d37eaa366fa8\" https://localhost:9443/api/mgw/adapter/0.1/apis?preserveProvider=false\u0026overwrite=true",
        "x-wso2-request": "POST https://localhost:9443/api/mgw/adapter/0.1/apis\nAuthorization: Bearer ae4eae22-3f65-387b-a171-d37eaa366fa8\n",
        "x-wso2-response": "HTTP/1.1 200 OK\nAPI imported successfully."
      }
    },
    "/apis/delete": {
      "post": {
        "security": [
          {
            "BasicAuth": []
          }
        ],
        "description": "This operation can be used to delete a API that was deployed\n",
        "tags": [
          "API (Individual)"
        ],
        "summary": "Delete deployed API",
        "parameters": [
          {
            "type": "string",
            "description": "Name of the API\n",
            "name": "apiName",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "version of the API\n",
            "name": "version",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "Virtual Host of the API\n",
            "name": "vhost",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK.\nAPI successfully undeployed from the Microgateway.\n",
            "schema": {
              "$ref": "#/definitions/DeployResponse"
            }
          },
          "400": {
            "description": "Bad Request.\nInvalid request or validation error\n",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Not Found.\nRequested API does not exist.\n",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        },
        "x-wso2-curl": "curl -k -H \"Authorization: Bearer ae4eae22-3f65-387b-a171-d37eaa366fa8\" \n-d '{\"apiName\":\"petstore\", \"version\":\"1.1\", \"vhost\":\"pets\"}'\n-X POST \"https://127.0.0.1:9443/api/mgw/adapter/0.1/apis/delete\"\n",
        "x-wso2-request": "POST https://127.0.0.1:9443/api/mgw/adapter/0.1/apis/delete\nAuthorization: Bearer ae4eae22-3f65-387b-a171-d37eaa366fa8\n{\"apiName\":\"petstore\", \"version\":\"1.1\", \"vhost\":\"pets\"}\n",
        "x-wso2-response": "HTTP/1.1 200 OK"
      }
    }
  },
  "definitions": {
    "APIMeta": {
      "properties": {
        "count": {
          "description": "Number of APIs returned in the response",
          "type": "integer"
        },
        "list": {
          "description": "All or sub set of info about APIs in the MGW",
          "type": "array",
          "items": {
            "$ref": "#/definitions/APIMetaListItem"
          }
        },
        "total": {
          "description": "Total number of APIs available in the MGW",
          "type": "integer"
        }
      }
    },
    "APIMetaListItem": {
      "type": "object",
      "properties": {
        "apiName": {
          "type": "string"
        },
        "apiType": {
          "type": "string"
        },
        "labels": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "version": {
          "type": "string"
        }
      }
    },
    "DeployResponse": {
      "type": "object",
      "properties": {
        "action": {
          "type": "string"
        },
        "info": {
          "type": "string"
        }
      }
    },
    "Error": {
      "title": "Error object returned with 4XX HTTP status",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "Error code",
          "type": "integer",
          "format": "int64"
        },
        "description": {
          "description": "A detail description about the error message.\n",
          "type": "string"
        },
        "error": {
          "description": "If there are more than one error list them out.\nFor example, list out validation errors by each field.\n",
          "type": "array",
          "items": {
            "$ref": "#/definitions/ErrorListItem"
          }
        },
        "message": {
          "description": "Error message.",
          "type": "string"
        },
        "moreInfo": {
          "description": "Preferably an url with more details about the error.\n",
          "type": "string"
        }
      }
    },
    "ErrorListItem": {
      "title": "Description of individual errors that may have occurred during a request.",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "Error code",
          "type": "string"
        },
        "message": {
          "description": "Description about individual errors occurred\n",
          "type": "string"
        }
      }
    },
    "Principal": {
      "type": "object",
      "properties": {
        "tenant": {
          "type": "string"
        },
        "token": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "BasicAuth": {
      "type": "basic"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json",
    "multipart/form-data"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This document specifies a **RESTful API** for WSO2 **API Microgateway** - Adapter.\n",
    "title": "WSO2 API Microgateway - Adapter",
    "contact": {
      "name": "WSO2",
      "url": "http://wso2.com/products/api-manager/",
      "email": "architecture@wso2.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "v1.2"
  },
  "host": "apis.wso2.com",
  "basePath": "/api/mgw/adapter/0.1",
  "paths": {
    "/apis": {
      "get": {
        "security": [
          {
            "BasicAuth": []
          }
        ],
        "description": "This operation can be used to retrieve meta info about all APIs\n",
        "tags": [
          "API (Collection)"
        ],
        "summary": "Get a list of API metadata",
        "parameters": [
          {
            "type": "string",
            "description": "Optional - Condition to filter APIs. Currently only filtering \nby API type (HTTP or WebSocket) is supported.\n\"http\" for HTTP type\n\"ws\" for WebSocket type\n",
            "name": "apiType",
            "in": "query"
          },
          {
            "type": "integer",
            "description": "Number of APIs (APIMeta objects to return)\n",
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "An array of API Metadata",
            "schema": {
              "$ref": "#/definitions/APIMeta"
            }
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        },
        "x-wso2-curl": "curl -k -H \"Authorization: Bearer ae4eae22-3f65-387b-a171-d37eaa366fa8\" \n-X GET \"https://127.0.0.1:9443/api/mgw/adapter/0.1/apis\"\n",
        "x-wso2-request": "GET https://127.0.0.1:9443/api/mgw/adapter/0.1/apis?apiType=http\nAuthorization: Bearer ae4eae22-3f65-387b-a171-d37eaa366fa8\n",
        "x-wso2-response": "HTTP/1.1 200 OK"
      },
      "post": {
        "security": [
          {
            "BasicAuth": []
          }
        ],
        "description": "This operation can be used to import an API.\n",
        "consumes": [
          "multipart/form-data"
        ],
        "tags": [
          "API (Individual)"
        ],
        "summary": "Import an API",
        "parameters": [
          {
            "type": "file",
            "x-exportParamName": "File",
            "description": "Zip archive consisting on exported api configuration\n",
            "name": "file",
            "in": "formData",
            "required": true
          },
          {
            "type": "boolean",
            "x-exportParamName": "PreserveProvider",
            "x-optionalDataType": "Bool",
            "description": "Preserve Original Provider of the API. This is the user choice to keep or replace the API provider.\n",
            "name": "preserveProvider",
            "in": "query"
          },
          {
            "type": "boolean",
            "x-exportParamName": "Overwrite",
            "x-optionalDataType": "Bool",
            "description": "Whether to update the API or not. This is used when updating already existing APIs.\n",
            "name": "overwrite",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful.\nAPI deployed or updated Successfully.\n",
            "schema": {
              "$ref": "#/definitions/DeployResponse"
            }
          },
          "403": {
            "description": "Forbidden.\nNot Authorized to deploy or update.\n",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Not Found. \nRequested API to update not found (when overwrite parameter is included).\n",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "409": {
            "description": "Conflict.\nAPI to import already exists (when overwrite parameter is not included).\n",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        },
        "x-wso2-curl": "curl -k -F \"file=@exported.zip\" -X POST -H \"Authorization: Bearer ae4eae22-3f65-387b-a171-d37eaa366fa8\" https://localhost:9443/api/mgw/adapter/0.1/apis?preserveProvider=false\u0026overwrite=true",
        "x-wso2-request": "POST https://localhost:9443/api/mgw/adapter/0.1/apis\nAuthorization: Bearer ae4eae22-3f65-387b-a171-d37eaa366fa8\n",
        "x-wso2-response": "HTTP/1.1 200 OK\nAPI imported successfully."
      }
    },
    "/apis/delete": {
      "post": {
        "security": [
          {
            "BasicAuth": []
          }
        ],
        "description": "This operation can be used to delete a API that was deployed\n",
        "tags": [
          "API (Individual)"
        ],
        "summary": "Delete deployed API",
        "parameters": [
          {
            "type": "string",
            "description": "Name of the API\n",
            "name": "apiName",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "version of the API\n",
            "name": "version",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "Virtual Host of the API\n",
            "name": "vhost",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK.\nAPI successfully undeployed from the Microgateway.\n",
            "schema": {
              "$ref": "#/definitions/DeployResponse"
            }
          },
          "400": {
            "description": "Bad Request.\nInvalid request or validation error\n",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Not Found.\nRequested API does not exist.\n",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        },
        "x-wso2-curl": "curl -k -H \"Authorization: Bearer ae4eae22-3f65-387b-a171-d37eaa366fa8\" \n-d '{\"apiName\":\"petstore\", \"version\":\"1.1\", \"vhost\":\"pets\"}'\n-X POST \"https://127.0.0.1:9443/api/mgw/adapter/0.1/apis/delete\"\n",
        "x-wso2-request": "POST https://127.0.0.1:9443/api/mgw/adapter/0.1/apis/delete\nAuthorization: Bearer ae4eae22-3f65-387b-a171-d37eaa366fa8\n{\"apiName\":\"petstore\", \"version\":\"1.1\", \"vhost\":\"pets\"}\n",
        "x-wso2-response": "HTTP/1.1 200 OK"
      }
    }
  },
  "definitions": {
    "APIMeta": {
      "properties": {
        "count": {
          "description": "Number of APIs returned in the response",
          "type": "integer"
        },
        "list": {
          "description": "All or sub set of info about APIs in the MGW",
          "type": "array",
          "items": {
            "$ref": "#/definitions/APIMetaListItem"
          }
        },
        "total": {
          "description": "Total number of APIs available in the MGW",
          "type": "integer"
        }
      }
    },
    "APIMetaListItem": {
      "type": "object",
      "properties": {
        "apiName": {
          "type": "string"
        },
        "apiType": {
          "type": "string"
        },
        "labels": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "version": {
          "type": "string"
        }
      }
    },
    "DeployResponse": {
      "type": "object",
      "properties": {
        "action": {
          "type": "string"
        },
        "info": {
          "type": "string"
        }
      }
    },
    "Error": {
      "title": "Error object returned with 4XX HTTP status",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "Error code",
          "type": "integer",
          "format": "int64"
        },
        "description": {
          "description": "A detail description about the error message.\n",
          "type": "string"
        },
        "error": {
          "description": "If there are more than one error list them out.\nFor example, list out validation errors by each field.\n",
          "type": "array",
          "items": {
            "$ref": "#/definitions/ErrorListItem"
          }
        },
        "message": {
          "description": "Error message.",
          "type": "string"
        },
        "moreInfo": {
          "description": "Preferably an url with more details about the error.\n",
          "type": "string"
        }
      }
    },
    "ErrorListItem": {
      "title": "Description of individual errors that may have occurred during a request.",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "Error code",
          "type": "string"
        },
        "message": {
          "description": "Description about individual errors occurred\n",
          "type": "string"
        }
      }
    },
    "Principal": {
      "type": "object",
      "properties": {
        "tenant": {
          "type": "string"
        },
        "token": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "BasicAuth": {
      "type": "basic"
    }
  }
}`))
}
