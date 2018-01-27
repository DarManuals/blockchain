package main

import (
	"github.com/borudar/blockchain/controllers"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/borudar/blockchain/db"
)

func main() {
	db.Load()

	router := mux.NewRouter()
	router.HandleFunc("/add_data", controllers.AddData).Methods("POST")
	router.HandleFunc("/last_blocks/{count}", controllers.GetBlocks).Methods("GET")
	router.HandleFunc("/management/add_transaction", controllers.AddTransaction).Methods("POST")

	http.ListenAndServe(":3000", router)
}
