package models

import "time"

type (
	Block struct {
		PreviousBlockHash string
		Rows              []string
		Timestamp         time.Time
		BlockHash         string
	}

	Data struct {
		SomeData string
	}
)
