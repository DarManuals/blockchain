package db

import (
	"encoding/json"
	"github.com/darmanuals/blockchain/models"
	"github.com/go-redis/redis"
	"log"
)

var (
	Blocks        = make(map[string]models.Block)
	Balances      = make(map[string]int)
	Hosts         = make(map[string]string)
	LastBlockHash = "0"
	redisCli         *redis.Client
)

func Load() {
	redisCli = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	st := redisCli.Ping()
	if st.String() != "ping: PONG" {
		log.Fatal(st.String())
	}

	lastHash := redisCli.Get("last_hash")
	if lastHash.Val() == "" {
		LastBlockHash = "0"
		err := SaveLastHash("0")
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Got last block hash: ", LastBlockHash)
	} else {
		LastBlockHash = lastHash.Val()
		log.Println("Got last block hash: ", LastBlockHash)
	}
}

func SaveBlock(block models.Block) error {
	b, err := json.Marshal(block)
	if err != nil {
		return err
	}
	return redisCli.Set(block.BlockHash, b, 0).Err()
}

func GetBlock(hash string) (models.Block, error) {
	res := redisCli.Get(hash)
	if res.Err() != nil {
		return models.Block{}, res.Err()
	}

	b, err := res.Bytes()
	if err != nil {
		return models.Block{}, err
	}

	var block models.Block
	err = json.Unmarshal(b, &block)
	if err != nil {
		return models.Block{}, err
	}

	return block, nil
}

func SaveLastHash(hash string) error {
	return redisCli.Set("last_hash", hash, 0).Err()
}
