// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/addresses": {
            "get": {
                "description": "GetAddresses is an example controller that fetches addresses.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "addresses"
                ],
                "summary": "Get all addresses",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "Filter by active status",
                        "name": "active",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by location type",
                        "name": "location_type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully retrieved",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Address"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "CreateAddress is an example controller to create a new address.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "addresses"
                ],
                "summary": "Create a new address",
                "parameters": [
                    {
                        "description": "Address information",
                        "name": "address",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AddressRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Successfully created",
                        "schema": {
                            "$ref": "#/definitions/models.Address"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/addresses/{id}": {
            "get": {
                "description": "GetAddressByID is an example controller that fetches an address by its ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "addresses"
                ],
                "summary": "Get an address by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Address ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully retrieved",
                        "schema": {
                            "$ref": "#/definitions/models.Address"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Address not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "UpdateAddress is an example controller that updates an address by its ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "addresses"
                ],
                "summary": "Update an address by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Address ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated address information",
                        "name": "address",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Address"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Address updated successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Address not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "DeleteAddress is an example controller that deletes an address by its ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "addresses"
                ],
                "summary": "Delete an address by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Address ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Address deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Address not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Address": {
            "type": "object",
            "required": [
                "latitude",
                "longitude",
                "name"
            ],
            "properties": {
                "active": {
                    "type": "boolean"
                },
                "complementary_informations": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "door_code": {
                    "type": "string"
                },
                "floor": {
                    "$ref": "#/definitions/models.FloorType"
                },
                "id": {
                    "type": "integer"
                },
                "latitude": {
                    "type": "number"
                },
                "lift": {
                    "type": "string"
                },
                "loading_dock": {
                    "type": "boolean"
                },
                "location_type": {
                    "$ref": "#/definitions/models.LocationType"
                },
                "longitude": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                },
                "side_loading": {
                    "type": "boolean"
                },
                "time_zone": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "yard": {
                    "$ref": "#/definitions/models.YardType"
                }
            }
        },
        "models.AddressRequest": {
            "type": "object",
            "properties": {
                "latitude": {
                    "type": "number"
                },
                "longitude": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.FloorType": {
            "type": "string",
            "enum": [
                "basement",
                "sidewalk",
                "ground_floor",
                "first",
                "second",
                "third",
                "fourth",
                "fifth",
                "sixth",
                "seventh"
            ],
            "x-enum-varnames": [
                "Basement",
                "Sidewalk",
                "GroundFloor",
                "First",
                "Second",
                "Third",
                "Fourth",
                "Fifth",
                "Sixth",
                "Seventh"
            ]
        },
        "models.LocationType": {
            "type": "string",
            "enum": [
                "individual",
                "company",
                "retail_store",
                "event",
                "supermarket",
                "warehouse",
                "distribution_platform"
            ],
            "x-enum-varnames": [
                "Individual",
                "Company",
                "RetailStore",
                "Event",
                "Supermarket",
                "Warehouse",
                "DistributionPlatform"
            ]
        },
        "models.YardType": {
            "type": "string",
            "enum": [
                "none",
                "inf_10m",
                "between_10_30m",
                "sup_30m"
            ],
            "x-enum-varnames": [
                "None",
                "Inf_10m",
                "Between_10_30m",
                "Sup_30m"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3001",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Address Service API",
	Description:      "API for address management",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
