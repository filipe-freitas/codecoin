package dto

import "time"

type Transaction struct {
	ID              string
	Name            string
	Number          string
	ExpirationMonth uint16
	ExpirationYear  uint16
	CVV             uint16
	Amount          float64
	Store           string
	Description     string
	CreatedAt       time.Time
}
