// Package signupsheetmem contains functions for processing signupsheet requests
package signupsheetmem

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Signupsheet type structure for holding signupsheet data
type Signupsheet struct {
	ID          string       `json:"id"`
	EntityID    string       `json:"entityId"`
	EntityType  string       `json:"entityType"`
	FunctionSub *FunctionSub `json:"functionSub"`
}

// FunctionSub type structure for holding function subscription data
type FunctionSub struct {
	FunctionName string `json:"functionName"`
	FunctionType string `json:"functionType"`
}

// Signupsheets map for storing signupsheets in memory
var Signupsheets = make(map[string]Signupsheet)

// GetSignupsheet func for finding a specific signupsheet
func GetSignupsheet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	value, present := Signupsheets[id]
	if present {
		w.Write([]byte("[\n"))
		json.NewEncoder(w).Encode(value)
		w.Write([]byte("]"))
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Signupsheet with ID " + id + " not found."))
	}
}

// GetSignupsheets func for listing all the signupsheets
func GetSignupsheets(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("[\n"))
	for _, value := range Signupsheets {
		json.NewEncoder(w).Encode(value)
	}
	w.Write([]byte("]"))
}

// UpdateOrCreateSignupsheet func for updating or creating a signupsheet
func UpdateOrCreateSignupsheet(w http.ResponseWriter, r *http.Request) {
	var signupsheet []Signupsheet
	err := json.NewDecoder(r.Body).Decode(&signupsheet)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		for i := range signupsheet {
			s := signupsheet[i]
			log.Printf("%+v\n", signupsheet)
			Signupsheets[s.ID] = s
			log.Printf("%+v\n", Signupsheets)
		}
		w.Write([]byte("Signupsheets updated or created successfully."))
	}
}

// DeleteSignupsheet func for deleting a signupsheet
func DeleteSignupsheet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	_, present := Signupsheets[id]
	if present {
		delete(Signupsheets, id)
		w.Write([]byte("Signupsheet with ID " + id + " deleted successfully."))
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Signupsheet with ID " + id + " not found."))
	}
}
