package service

import (
	"crypto/sha256"
	"encoding/base64"
	"github.com/borudar/blockchain/models"
	"strconv"
	"time"
)

var (
	blocks        = make(map[string]models.Block)
	lastBlockHash = "0"
)

func AddBlock(data []string) {
	hasher := sha256.New()
	block := models.Block{
		PreviousBlockHash: lastBlockHash,
		Rows:              append(data),
		Timestamp:         time.Now(),
	}
	hasher.Write(BlockBytes(block))
	block.BlockHash = base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	blocks[block.BlockHash] = block
	lastBlockHash = block.BlockHash
}

func BlockBytes(block models.Block) (b []byte) {
	b = append([]byte(block.PreviousBlockHash))
	for _, row := range block.Rows {
		b = append(b, []byte(row)...)
	}
	b = append(b, []byte(strconv.Itoa(block.Timestamp.Second()))...)
	return
}

func GetBlocks(count int) []models.Block {
	result := make([]models.Block, 0)
	key := lastBlockHash

	if count > len(blocks) {
		count = len(blocks)
	}

	for i := 0; i < count; i++ {
		block, ok := blocks[key]
		if !ok {
			break
		}
		key = block.PreviousBlockHash
		result = append(result, block)
	}
	return result
}
