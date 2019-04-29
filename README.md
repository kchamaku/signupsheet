# Signupsheet REST API demo using Go
A RESTful API to demo Signupsheets CRUD operations with Go using **gorilla/mux**
and Go standard lib map collection for storing signupsheet data in memory.

-- Covers some logging, validation and error handling
-- Does not cover security, fault-tolerance, resilence...

# Run Go process 
signupsheet 8888

# API Endpoint : http://localhost:8888

## Structure
signupsheet
 	|-- main.go		// Init logging, Requests Routing, HTTP listerner
	|-- rest-api	// REST API functions for signupsheet

#### /signupsheet
* HTTP Method: GET - Get all signupsheets
* HTTP Method: POST - Update or Create signupsheets (single or bulk request)

#### /signupsheet/{id}
* HTTP Method: GET - Get signupsheet based on parameter id
* HTTP Method: DELETE - Delete signupsheet based on parameter id