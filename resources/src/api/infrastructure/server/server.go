package server

import (
	"encoding/json"
	"fmt"
	"log"
)

type Envelope struct {
	Type string
	Msg  map[string]interface{}
}

type Account struct {
	ActiveCard 		bool `json:"active-card"`
	AvailableLimit  float64  `json:"available-limit"`
}

type Transaction struct {
	Merchant string //merchant
	Amount int
	Time string
}

func Start() {
	message := `{"account": {"active-card": true, "available-limit": 100}}`
	var result map[string]map[string]interface{}
	json.Unmarshal([]byte(message), &result)
	fmt.Println(result)

	for k := range result {
		if k == "account" {
			fmt.Println(k)
			account := Account{
				ActiveCard:     result["account"]["active-card"].(bool),
				AvailableLimit: result["account"]["available-limit"].(float64),
			}

			buf, err := json.Marshal(account)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s\n", buf)
		}
	}
}