package service

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/borudar/blockchain/db"
	"github.com/borudar/blockchain/models"
	"log"
	"strconv"
	"time"
)

var (
	LastBlockHash = "0"
)

func AddBlock(data []models.Tx) {
	hasher := sha256.New()
	block := models.Block{
		PreviousBlockHash: LastBlockHash,
		Rows:              append(data),
		Timestamp:         time.Now(),
	}
	hasher.Write(BlockBytes(block))
	block.BlockHash = hex.EncodeToString(hasher.Sum(nil))

	db.Blocks[block.BlockHash] = block
	LastBlockHash = block.BlockHash

	val, _ := json.Marshal(block)

	if err := db.Session.Query(`INSERT INTO blocks (key, val) VALUES (?, ?)`,
		block.BlockHash, string(val)).Exec(); err != nil {
		log.Fatal(err)
	}
}

func BlockBytes(block models.Block) (b []byte) {
	b = append([]byte(block.PreviousBlockHash))
	for _, row := range block.Rows {
		b = append(b, []byte(row.From)...)
		b = append(b, []byte(row.To)...)
		b = append(b, []byte(strconv.Itoa(row.Amount))...)
	}
	b = append(b, []byte(strconv.Itoa(block.Timestamp.Second()))...)
	return
}

func GetBlocks(count int) []models.Block {
	result := make([]models.Block, 0)
	key := LastBlockHash

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
