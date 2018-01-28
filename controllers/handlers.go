package controllers

import (
	"encoding/json"
	"github.com/darmanuals/blockchain/db"
	"github.com/darmanuals/blockchain/models"
	"github.com/darmanuals/blockchain/service"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
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
	blocks := service.GetBlocks(int(n))
	json.NewEncoder(w).Encode(blocks)
}

func GetStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	stat := models.Status{
		Id:       88,
		Name:     "Bogdan",
		URL:      "192.168.44.88:3000",
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
	var lastUrl string

	for _, url := range db.Hosts {
		b, _ := http.Get("http://" + url + "/blockchain/get_blocks/100")
		body, err := ioutil.ReadAll(b.Body)
		if err != nil {
			continue
		}
		b.Body.Close()

		err = json.Unmarshal(body, &tmpBlocks)
		if err != nil {
			log.Println("Unmarshal: ", err)
		}
		if len(blocks) < len(tmpBlocks) {
			blocks = tmpBlocks
			lastUrl = url
		}
	}

	if len(blocks) < 1 {
		log.Println("nothing to do")
		return
	}

	db.Blocks = make(map[string]models.Block)
	for _, val := range blocks {
		db.Blocks[val.BlockHash] = val
	}

	b, _ := http.Get("http://" + lastUrl + "/management/status")
	bodyBytes, err := ioutil.ReadAll(b.Body)
	if err != nil {
		return
	}
	b.Body.Close()
	var status models.Status
	json.Unmarshal(bodyBytes, &status)
	if len(status.LastHash) > 1 {
		log.Println(status.LastHash)
		service.LastBlockHash = status.LastHash
	}

	log.Println("Sync: Blocks: ", db.Blocks)
	json.NewEncoder(w).Encode(db.Blocks)
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var up models.Updates
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	defer r.Body.Close()

	json.Unmarshal(body, &up)
	log.Println(string(body))
	log.Println("struct: ", up)

	resp := models.UpdateResp{
		Success: true,
		ErrCode: "0x0000",
		Message: "Ok",
	}

	json.NewEncoder(w).Encode(resp)
}
