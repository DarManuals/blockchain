package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"

	"github.com/darmanuals/blockchain/models"
	"github.com/darmanuals/blockchain/service"
)

var (
	tmpData []string
	mu      = &sync.Mutex{}
)

func AddData(w http.ResponseWriter, r *http.Request) {
	var data models.Data
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	defer r.Body.Close()
	json.Unmarshal(body, &data)

	if data.SomeData != "" {
		mu.Lock()
		tmpData = append(tmpData, data.SomeData)
		if len(tmpData) > 4 {
			service.AddBlock(tmpData)
			tmpData = make([]string, 0)
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
