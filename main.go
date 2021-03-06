package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/darmanuals/blockchain/controllers"
	"github.com/darmanuals/blockchain/db"
)

func main() {
	db.Load()

	router := mux.NewRouter()
	router.HandleFunc("/blockchain/get_blocks/{count}", controllers.GetBlocks).Methods("GET")
	router.HandleFunc("/blockchain/receive_update", controllers.Update).Methods("POST")
	router.HandleFunc("/management/add_transaction", controllers.AddTransaction).Methods("POST")
	router.HandleFunc("/management/add_link", controllers.AddLink).Methods("POST")
	router.HandleFunc("/management/state", controllers.GetStatus).Methods("GET")
	router.HandleFunc("/management/sync", controllers.Sync).Methods("GET")

	http.ListenAndServe(":3000", router)
}
