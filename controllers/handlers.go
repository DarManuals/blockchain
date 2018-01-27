package controllers

import (
	"encoding/json"
	"github.com/borudar/blockchain/models"
	"github.com/borudar/blockchain/service"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
)

var (
	tmpData []models.Tx
	mu      = &sync.Mutex{}
)

func AddData(w http.ResponseWriter, r *http.Request) {
	var data models.Tx
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	defer r.Body.Close()
	json.Unmarshal(body, &data)

	if data.To != "" {
		mu.Lock()
		tmpData = append(tmpData, data)
		if len(tmpData) > 4 {
			service.AddBlock(tmpData)
			tmpData = make([]models.Tx, 0)
		}
		mu.Unlock()
	}
}

func GetBlocks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	count := mux.Vars(r)["count"]
	n, err := strconv.ParseInt(count, 10, 32)
	if err != nil {
		return
	}
	mu.Lock()
	blocks := service.GetBlocks(int(n))
	mu.Unlock()
	json.NewEncoder(w).Encode(blocks)
}

func AddTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}