package models

type (
	Block struct {
		PreviousBlockHash string `json:"prev_hash"`
		Tx                []Tx   `json:"tx"`
		Timestamp         int    `json:"ts"`
		BlockHash         string `json:"hash"`
	}

	Tx struct {
		From   string `json:"from"`
		To     string `json:"to"`
		Amount int    `json:"amount"`
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

	Updates struct {
		Sender int   `json:"sender_id"`
		Block  Block `json:"block"`
	}

	UpdateResp struct {
		Success bool   `json:"success"`
		ErrCode string `json:"err_code"`
		Message string `json:"message"`
	}
)
