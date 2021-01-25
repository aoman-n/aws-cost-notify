package domain

type Billing struct {
	CurrentMonthTotal string `json:"currentMonthTotal"`
	PreviousDay string `json:"previousDay"`
}
