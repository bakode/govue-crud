package main

import (
	"log"
	"net/http"

	"github.com/aasumitro/govue/server/cmd"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/examples", cmd.Create).Methods("POST")
	router.HandleFunc("/examples", cmd.FindAll).Methods("GET")
	router.HandleFunc("/examples/{id}", cmd.FindOne).Methods("GET")
	router.HandleFunc("/examples/{id}/update", cmd.Edit).Methods("PUT")
	// router.HandleFunc("/examples/{id}/delete", cmd.Delete).Methods("DELETE")
	router.HandleFunc("/examples/{id}/delete", cmd.Delete).Methods("GET")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
