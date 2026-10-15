package financialresponse

type FinancialCurrentBalanceResponse struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}
