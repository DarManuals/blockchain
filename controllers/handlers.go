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
	"github.com/borudar/blockchain/db"
	"log"
)

var (
	tmpData []models.Tx
	mu      = &sync.Mutex{}
)

func AddTransaction(w http.ResponseWriter, r *http.Request) {
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

func GetStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	stat := models.Status{
		Id: 88,
		Name: "Bogdan",
		URL: "192.168.44.88:3000",
		LastHash: service.LastBlockHash,
	}

	neighbours := []string{}
	for host, _ := range db.Hosts {
		neighbours = append(neighbours, host)
	}
	stat.Neighbours = neighbours
	json.NewEncoder(w).Encode(stat)
	log.Println("got request")
	log.Println("send: ", stat)
}

func AddLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var host models.Host
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	defer r.Body.Close()
	json.Unmarshal(body, &host)

	if len(host.Id) > 1 && len(host.URL) > 1 {
		db.Hosts[host.Id] = host.URL
	}
	json.NewEncoder(w).Encode(host)
}

func Sync(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var blocks []models.Block
	var tmpBlocks []models.Block

	for _, url := range db.Hosts {
		b, _ := http.Get("http://"+url+"/blockchain/get_blocks/10000")
		body, err := ioutil.ReadAll(b.Body)
		if err != nil {
			return
		}
		b.Body.Close()

		err = json.Unmarshal(body,&tmpBlocks)
		if err != nil {
			log.Println("Unmarshal: ", err)
		}
		if len(blocks) < len(tmpBlocks) {
			blocks = tmpBlocks
		}
	}

	db.Blocks = make(map[string]models.Block)
	for _, val := range blocks {
		db.Blocks[val.BlockHash] = val
	}
	log.Println("Sync: Blocks: ", db.Blocks)
}