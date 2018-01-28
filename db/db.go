package db

import (
	"github.com/borudar/blockchain/models"
)

var (
	Blocks   = make(map[string]models.Block)
	Balances = make(map[string]int)
	Hosts    = make(map[string]string)
)
