// Package main for handling the signupsheet REST API requests
// and starting the HTTP listener service
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	signupsheet "signupsheetmem/rest-api"
	"strconv"

	"github.com/gorilla/mux"
)

func init() {
	// Initialize and setup logging
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	log.Println("init started")
}

func main() {
	port := flag.Int("port", 8000, "port number")
	flag.Parse()
	fmt.Println("port=", *port)
	fmt.Println("flag count =", flag.NFlag())
	// slice to get port number
	if flag.NFlag() != 1 {
		fmt.Println("Usage: signupsheet --port=number")
		os.Exit(1)
	}
	// REST API requests
	router := mux.NewRouter()
	router.HandleFunc("/signupsheet", signupsheet.GetSignupsheets).Methods("GET")
	router.HandleFunc("/signupsheet/{id}", signupsheet.GetSignupsheet).Methods("GET")
	router.HandleFunc("/signupsheet", signupsheet.UpdateOrCreateSignupsheet).Methods("POST")
	router.HandleFunc("/signupsheet/{id}", signupsheet.DeleteSignupsheet).Methods("DELETE")
	portAddress := ":" + strconv.Itoa(*port)
	// Start HTTP listener
	log.Fatal(http.ListenAndServe(portAddress, router))
}
