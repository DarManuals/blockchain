package main

import (
	"github.com/borudar/blockchain/controllers"
	"github.com/borudar/blockchain/db"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	db.Load()

	router := mux.NewRouter()
	router.HandleFunc("/blockchain/get_blocks/{count}", controllers.GetBlocks).Methods("GET")
	router.HandleFunc("/management/add_transaction", controllers.AddTransaction).Methods("POST")
	router.HandleFunc("/management/add_link", controllers.AddLink).Methods("POST")
	router.HandleFunc("/management/status", controllers.GetStatus).Methods("GET")
	router.HandleFunc("/management/sync", controllers.Sync).Methods("GET")

	http.ListenAndServe(":3000", router)
}
