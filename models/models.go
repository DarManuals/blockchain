package models

import "time"

//192.168.44.88

const ID string = "88"

type (
	Block struct {
		PreviousBlockHash string
		Rows              []Tx
		Timestamp         time.Time
		BlockHash         string
	}

	Data struct {
		SomeData string
	}

	Tx struct {
		From   string
		To     string
		Amount int
	}
)
