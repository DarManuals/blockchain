package service

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/darmanuals/blockchain/db"
	"github.com/darmanuals/blockchain/models"
	"log"
	"net/http"
	"strconv"
	"time"
)

func AddBlock(data []models.Tx) {
	hasher := sha256.New()
	block := models.Block{
		PreviousBlockHash: db.LastBlockHash,
		Tx:                append(data),
		Timestamp:         int(time.Now().UTC().Unix()),
	}
	hasher.Write(BlockBytes(block))
	block.BlockHash = hex.EncodeToString(hasher.Sum(nil))

	db.SaveBlock(block)
	db.SaveLastHash(block.BlockHash)
	db.LastBlockHash = block.BlockHash

	for _, url := range db.Hosts {
		up := models.Updates{
			Sender: 88,
			Block:  block,
		}
		b, _ := json.Marshal(up)
		r := bytes.NewReader(b)
		http.Post("http://"+url+"/blockchain/receive_update", "application/json", r)
		log.Println("send: ", up)
	}
}

func BlockBytes(block models.Block) (b []byte) {
	b = append([]byte(block.PreviousBlockHash))
	for _, row := range block.Tx {
		b = append(b, []byte(row.From)...)
		b = append(b, []byte(row.To)...)
		b = append(b, []byte(strconv.Itoa(row.Amount))...)
	}
	b = append(b, []byte(strconv.Itoa(block.Timestamp))...)
	return
}

func GetBlocks(count int) []models.Block {
	result := make([]models.Block, 0)
	key := db.LastBlockHash

	for i := 0; i < count; i++ {
		block, err := db.GetBlock(key)
		if err != nil {
			log.Println("get blocks err: ", err)
			break
		}

		key = block.PreviousBlockHash
		if key == "0" {break}

		result = append(result, block)
	}
	return result
}
