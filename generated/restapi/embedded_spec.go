// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "produces": [
    "application/json",
    "text/text"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "An API for providing identifiers.",
    "title": "identifier",
    "contact": {
      "name": "Christina Harlow",
      "email": "cmharlow@stanford.edu"
    },
    "version": "0.1.0"
  },
  "host": "identifiers.dlss.stanford.edu",
  "basePath": "/v1",
  "paths": {
    "/healthcheck": {
      "get": {
        "description": "The healthcheck endpoint provides information about the health of the service.",
        "summary": "Health Check",
        "operationId": "healthCheck",
        "responses": {
          "200": {
            "description": "The service is functioning nominally",
            "schema": {
              "$ref": "#/definitions/HealthCheckResponse"
            }
          },
          "503": {
            "description": "The service is not working correctly",
            "schema": {
              "$ref": "#/definitions/HealthCheckResponse"
            }
          }
        }
      }
    },
    "/identifiers": {
      "get": {
        "produces": [
          "application/json"
        ],
        "operationId": "getIdentifiersInfo",
        "responses": {
          "200": {
            "description": "Get a list of identifier types \u0026 sources supported",
            "schema": {
              "$ref": "#/definitions/Sources"
            }
          }
        }
      }
    },
    "/identifiers/all": {
      "get": {
        "produces": [
          "application/json"
        ],
        "operationId": "getIdentifiersList",
        "responses": {
          "200": {
            "description": "Get a list of all Identifiers minted across types",
            "schema": {
              "$ref": "#/definitions/Sources"
            }
          }
        }
      }
    },
    "/identifiers/druids": {
      "get": {
        "operationId": "getCurrentDRUIDs",
        "responses": {
          "200": {
            "description": "List of DRUIDs currently being used",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Identifiers"
              }
            }
          }
        }
      },
      "post": {
        "operationId": "mintNewDRUIDs",
        "parameters": [
          {
            "type": "integer",
            "description": "Number of DRUIDs to mint. Default is 1.",
            "name": "quantity",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "New DRUIDs Created",
            "schema": {
              "$ref": "#/definitions/Identifiers"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "HealthCheckResponse": {
      "type": "object",
      "properties": {
        "status": {
          "description": "The status of the service",
          "type": "string"
        }
      },
      "example": {
        "status": "OK"
      }
    },
    "Identifier": {
      "type": "object",
      "required": [
        "id"
      ],
      "properties": {
        "id": {
          "type": "string"
        }
      }
    },
    "Identifiers": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Identifier"
      }
    },
    "Source": {
      "type": "object",
      "required": [
        "name",
        "template"
      ],
      "properties": {
        "name": {
          "type": "string"
        },
        "template": {
          "type": "string"
        }
      }
    },
    "Sources": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Source"
      }
    }
  }
}`))
}
