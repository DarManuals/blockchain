package db

import (
	"encoding/json"
	"github.com/borudar/blockchain/models"
	"github.com/gocql/gocql"
)

var Session *gocql.Session
var Blocks = make(map[string]models.Block)
var Balances = make(map[string]int)
var Hosts = make(map[string]string)

func Load() {
	var err error
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "blockchain"
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}

	var key, val string
	iter := Session.Query(`SELECT * FROM blocks`).Iter()
	for iter.Scan(&key, &val) {
		var block models.Block
		json.Unmarshal([]byte(val), &block)
		Blocks[key] = block
	}
	//log.Println(Blocks)
}
