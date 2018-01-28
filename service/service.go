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

	db.Blocks[block.BlockHash] = block
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

	if count > len(db.Blocks) || count == -1 {
		count = len(db.Blocks)
	}

	for i := 0; i < count; i++ {
		block, ok := db.Blocks[key]

		if !ok {
			break
		}
		key = block.PreviousBlockHash
		result = append(result, block)
	}
	return result
}
