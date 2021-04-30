package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// homeLink :: (/) route.
func homeLink(writerEntity http.ResponseWriter, requestEntity *http.Request){

	/* Base 64 Encoded Message just cuz.  */
	_, err := fmt.Fprintf(writerEntity, "bW9ua2k=")
	Check(err)
}

// GarageInformation :: (/GarageInformation) route.
func GarageInformation(writerEntity http.ResponseWriter, requestEntity *http.Request){
	err := json.NewEncoder(writerEntity).Encode(GarageRequest())
	writerEntity.Header().Set("Content-Type", "application/json; charset=UTF-8")
	Check(err)
}

// StartService :: The heart of the program.
func StartService(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/GarageInformation", GarageInformation)
	log.Fatal(http.ListenAndServe(":8080", router))
}