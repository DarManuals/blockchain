package db

import (
	"github.com/darmanuals/blockchain/models"
	"github.com/go-redis/redis"
	"log"
)

var (
	Blocks   = make(map[string]models.Block)
	Balances = make(map[string]int)
	Hosts    = make(map[string]string)
	LastBlockHash = "0"
	Redis 	*redis.Client
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
	if len(lastHash.Val()) < 1 {
		LastBlockHash = "0"
		err := Redis.Set("last_hash", "0", 0).Err()
		if err != nil {log.Fatal(err)}
	}
}
