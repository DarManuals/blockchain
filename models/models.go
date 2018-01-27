package models

import "time"

//const ID string = "88"

type (
	Block struct {
		PreviousBlockHash string    `json:"prev_hash"`
		Rows              []Tx      `json:"tx"`
		Timestamp         time.Time `json:"ts"`
		BlockHash         string    `json:"hash"`
	}

	Data struct {
		SomeData string
	}

	Tx struct {
		From   string
		To     string
		Amount int
	}

	Host struct {
		Id  string
		URL string
	}

	Status struct {
		Id         int      `json:"id"`
		Name       string   `json:"name"`
		LastHash   string   `json:"last_hash"`
		Neighbours []string `json:"neighbours"`
		URL        string   `json:"url"`
	}
)
