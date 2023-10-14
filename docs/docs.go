// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/v1/departments": {
            "post": {
                "description": "Get department by range",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "department"
                ],
                "summary": "Get department by range",
                "parameters": [
                    {
                        "description": "Department range request",
                        "name": "departmentData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.DepartmentRangeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Department"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable entity",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/departments/favourite": {
            "get": {
                "description": "Get favourite departments",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "department"
                ],
                "summary": "Get favourite departments",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Department"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/departments/favourite/{id}": {
            "post": {
                "description": "Add department to favourites",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "department"
                ],
                "summary": "Add department to favourites",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Department ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Added to favourites",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable entity",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes department from favourites",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "department"
                ],
                "summary": "Deletes department from favourites",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Department ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Deleted from favourites",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/departments/rating": {
            "post": {
                "description": "Add department rating",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "department"
                ],
                "summary": "Add department rating",
                "parameters": [
                    {
                        "description": "Department rating",
                        "name": "ratingData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.DepartmentRating"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Rating added",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable entity",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/departments/{id}": {
            "get": {
                "description": "Get department by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "department"
                ],
                "summary": "Get department by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Department ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Start latitude",
                        "name": "startLatitude",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Start longitude",
                        "name": "startLongitude",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Department"
                        }
                    },
                    "404": {
                        "description": "Department not found",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/search": {
            "post": {
                "description": "Create a new search record",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "search"
                ],
                "summary": "Create a new search record",
                "parameters": [
                    {
                        "description": "Search",
                        "name": "search",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SearchCreate"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Search",
                        "schema": {
                            "$ref": "#/definitions/model.Search"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/search/user": {
            "get": {
                "description": "Get search records for user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "search"
                ],
                "summary": "Get search records for user",
                "responses": {
                    "200": {
                        "description": "Searches",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Search"
                            }
                        }
                    },
                    "404": {
                        "description": "Searches not found",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/search/{searchId}": {
            "get": {
                "description": "Get search record by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "search"
                ],
                "summary": "Get search record by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Search id",
                        "name": "searchId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Search",
                        "schema": {
                            "$ref": "#/definitions/model.Search"
                        }
                    },
                    "404": {
                        "description": "Search not found",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/tickets": {
            "post": {
                "description": "Create a new ticket",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tickets"
                ],
                "summary": "Create a new ticket",
                "parameters": [
                    {
                        "description": "Ticket",
                        "name": "ticket",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TicketCreate"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Ticket id",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/tickets/department/{departmentId}": {
            "get": {
                "description": "Get all tickets for department",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tickets"
                ],
                "summary": "Get all tickets for department",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Department id",
                        "name": "departmentId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Tickets",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Ticket"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/tickets/user": {
            "get": {
                "description": "Get all tickets for user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tickets"
                ],
                "summary": "Get all tickets for user",
                "responses": {
                    "200": {
                        "description": "Tickets",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Ticket"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/tickets/{ticketId}": {
            "get": {
                "description": "Get ticket by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tickets"
                ],
                "summary": "Get ticket by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ticket id",
                        "name": "ticketId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Ticket",
                        "schema": {
                            "$ref": "#/definitions/model.Ticket"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Cancel ticket",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tickets"
                ],
                "summary": "Cancel ticket",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ticket id",
                        "name": "ticketId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Coordinates": {
            "type": "object",
            "properties": {
                "latitude": {
                    "type": "number",
                    "example": 55.892334
                },
                "longitude": {
                    "type": "number",
                    "example": 37.44055
                }
            }
        },
        "model.Department": {
            "type": "object",
            "properties": {
                "Biskvit_id": {
                    "type": "string",
                    "example": "5010"
                },
                "_id": {
                    "type": "string",
                    "example": "65298f171d9eaf1f3125fc41"
                },
                "address": {
                    "type": "string",
                    "example": "Московская область, г. Химки, ул. Пролетарская, д. 8, стр. 1"
                },
                "city": {
                    "type": "string",
                    "example": "Химки"
                },
                "coordinates": {
                    "$ref": "#/definitions/model.Coordinates"
                },
                "estimatedTimeCar": {
                    "type": "number"
                },
                "estimatedTimeWalk": {
                    "type": "number"
                },
                "favourite": {
                    "type": "boolean",
                    "default": false,
                    "example": true
                },
                "id": {
                    "type": "integer",
                    "example": 29000262
                },
                "location": {
                    "$ref": "#/definitions/model.Location"
                },
                "scheduleFl": {
                    "type": "string",
                    "example": "пн-пт: 10:00-20:00 сб: 10:00-17:00 вс: выходной"
                },
                "scheduleJurL": {
                    "type": "string",
                    "example": "пн-чт: 10:00-19:00 пт: 10:00-18:00 сб, вс: выходной"
                },
                "shortName": {
                    "type": "string",
                    "example": "ДО «ЦИК «Химки-Правобережный» Филиала № 7701 Банка ВТБ (ПАО)"
                },
                "special": {
                    "$ref": "#/definitions/model.Special"
                },
                "workload": {
                    "description": "историческое",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Workload"
                    }
                }
            }
        },
        "model.DepartmentRangeRequest": {
            "type": "object",
            "required": [
                "latitude",
                "longitude",
                "radius"
            ],
            "properties": {
                "latitude": {
                    "type": "number",
                    "example": 55.892334
                },
                "longitude": {
                    "type": "number",
                    "example": 37.44055
                },
                "radius": {
                    "description": "in km",
                    "type": "number",
                    "example": 10
                }
            }
        },
        "model.DepartmentRating": {
            "type": "object",
            "required": [
                "departmentId",
                "rating",
                "text"
            ],
            "properties": {
                "departmentId": {
                    "type": "string"
                },
                "rating": {
                    "type": "number"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "model.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "model.HourWorkload": {
            "type": "object",
            "properties": {
                "hour": {
                    "type": "string",
                    "example": "10:0-11:0"
                },
                "load": {
                    "type": "number",
                    "example": 0.3256373598976446
                }
            }
        },
        "model.Location": {
            "type": "object",
            "properties": {
                "coordinates": {
                    "$ref": "#/definitions/model.Coordinates"
                },
                "type": {
                    "type": "string",
                    "example": "Point"
                }
            }
        },
        "model.Search": {
            "type": "object",
            "required": [
                "atm",
                "coordinates",
                "createdAt",
                "online",
                "special",
                "text",
                "userId"
            ],
            "properties": {
                "_id": {
                    "type": "string",
                    "example": "5f9e9b9b9b9b9b9b9b9b9b9b"
                },
                "atm": {
                    "type": "boolean"
                },
                "coordinates": {
                    "$ref": "#/definitions/model.Coordinates"
                },
                "createdAt": {
                    "type": "string",
                    "example": "2021-01-01T00:00:00Z"
                },
                "online": {
                    "type": "boolean"
                },
                "special": {
                    "$ref": "#/definitions/model.SearchSpecial"
                },
                "text": {
                    "type": "string",
                    "example": "текст запроса"
                },
                "userId": {
                    "type": "string",
                    "example": "5f9e9b9b9b9b9b9b9b889b9b"
                }
            }
        },
        "model.SearchCreate": {
            "type": "object",
            "required": [
                "coordinates",
                "text"
            ],
            "properties": {
                "coordinates": {
                    "$ref": "#/definitions/model.Coordinates"
                },
                "test": {
                    "type": "boolean",
                    "example": true
                },
                "text": {
                    "type": "string",
                    "example": "текст запроса"
                }
            }
        },
        "model.SearchSpecial": {
            "type": "object",
            "properties": {
                "Prime": {
                    "type": "boolean",
                    "example": false
                },
                "juridical": {
                    "type": "boolean",
                    "example": true
                },
                "person": {
                    "type": "boolean",
                    "example": true
                },
                "ramp": {
                    "type": "boolean",
                    "example": true
                },
                "vipOffice": {
                    "type": "boolean",
                    "example": false
                },
                "vipZone": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "model.Special": {
            "type": "object",
            "properties": {
                "Prime": {
                    "type": "integer",
                    "example": 0
                },
                "juridical": {
                    "type": "integer",
                    "example": 1
                },
                "person": {
                    "type": "integer",
                    "example": 1
                },
                "ramp": {
                    "type": "integer",
                    "example": 1
                },
                "vipOffice": {
                    "type": "integer",
                    "example": 0
                },
                "vipZone": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "model.Ticket": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string",
                    "example": "5f9e3b4e1d9eaf1f3asdfc3f"
                },
                "createdAt": {
                    "type": "string",
                    "example": "2021-01-01T00:00:00Z"
                },
                "departmentId": {
                    "type": "string",
                    "example": "5f9e3b4eknjeaf1f3125fc3f"
                },
                "timeSlot": {
                    "type": "string",
                    "example": "2020-11-02T10:00:00.000Z"
                },
                "userId": {
                    "type": "string",
                    "example": "5f9e3b4e1d9jnh1f3125fc3f"
                }
            }
        },
        "model.TicketCreate": {
            "type": "object",
            "required": [
                "departmentId",
                "timeSlot"
            ],
            "properties": {
                "departmentId": {
                    "type": "string",
                    "example": "5f9e3b4e1d9eaf1f3125fc3f"
                },
                "timeSlot": {
                    "type": "string",
                    "example": "2020-11-02T10:00:00.000Z"
                }
            }
        },
        "model.Workload": {
            "type": "object",
            "properties": {
                "day": {
                    "type": "string",
                    "example": "пн"
                },
                "loadHours": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.HourWorkload"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "More Tech API",
	Description:      "More Tech API server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
