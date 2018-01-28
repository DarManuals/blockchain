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
	Redis         *redis.Client
)

func Load() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	st := Redis.Ping()
	if st.String() != "ping: PONG" {
		log.Fatal(st.String())
	}

	lastHash := Redis.Get("last_hash")
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
	return Redis.Set(block.BlockHash, b, 0).Err()
}

func GetBlock(hash string) (models.Block, error) {
	res := Redis.Get(hash)
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
	return Redis.Set("last_hash", hash, 0).Err()
}
