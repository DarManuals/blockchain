package db

import (
	"github.com/darmanuals/blockchain/models"
)

var (
	Blocks   = make(map[string]models.Block)
	Balances = make(map[string]int)
	Hosts    = make(map[string]string)
)
