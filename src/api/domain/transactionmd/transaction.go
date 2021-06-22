package transactionmd

import "time"

type Transaction struct {
	Merchant string
	Amount float64
	Time time.Time
}
