package db

import (
	"github.com/gocql/gocql"
	"github.com/borudar/blockchain/models"
	"encoding/json"
)

var Session *gocql.Session
var Blocks       = make(map[string]models.Block)


func Load()  {
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
