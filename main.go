package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/darmanuals/blockchain/controllers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/add_data", controllers.AddData).Methods("POST")
	router.HandleFunc("/last_blocks/{count}", controllers.GetBlocks).Methods("GET")
	http.ListenAndServe(":3000", router)
}
