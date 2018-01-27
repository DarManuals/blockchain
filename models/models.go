package models

import "time"

//const ID string = "88"

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

	Host struct {
		Id string
		URL string
	}

	Status struct {
		Id, Name, LastHash string
		Neighbours []string
		URL string
	}

)

