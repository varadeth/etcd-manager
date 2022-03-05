// Code generated by go-swagger; DO NOT EDIT.

package restapi

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
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "ETCD-Manager",
    "title": "Etcd Manager",
    "version": "1.0.0"
  },
  "host": "localhost",
  "basePath": "/v1",
  "paths": {
    "/connect": {
      "post": {
        "tags": [
          "connect"
        ],
        "operationId": "connect",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "$ref": "#/definitions/Connection"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "Etcd Connection not established"
          },
          "500": {
            "description": "Internal Server Error"
          }
        }
      }
    }
  },
  "definitions": {
    "Connection": {
      "type": "object",
      "required": [
        "urls"
      ],
      "properties": {
        "urls": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "ETCD-Manager",
    "title": "Etcd Manager",
    "version": "1.0.0"
  },
  "host": "localhost",
  "basePath": "/v1",
  "paths": {
    "/connect": {
      "post": {
        "tags": [
          "connect"
        ],
        "operationId": "connect",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "$ref": "#/definitions/Connection"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "Etcd Connection not established"
          },
          "500": {
            "description": "Internal Server Error"
          }
        }
      }
    }
  },
  "definitions": {
    "Connection": {
      "type": "object",
      "required": [
        "urls"
      ],
      "properties": {
        "urls": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    }
  }
}`))
}
