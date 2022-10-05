package entity

type Transaction struct {
	TransactionId int     `json:"id"`
	Name          string  `json:"name"`
	Value         float32 `json:"value"`
	Type          uint8   `json:"type"`
	Category      string  `json:"category"`
}
