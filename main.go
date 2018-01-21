package main

import (
	"github.com/borudar/blockchain/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/add_data", controllers.AddData).Methods("POST")
	router.HandleFunc("/last_blocks/{count}", controllers.GetBlocks).Methods("GET")
	http.ListenAndServe(":3000", router)
}
