package obligationmodel

import "time"

type ObligationSettlementsModel struct {
	Id            int
	ObligationId  int
	TransactionId int
	Amount        float64
	CreatedAt     time.Time
	CanceledAt    *time.Time
}
