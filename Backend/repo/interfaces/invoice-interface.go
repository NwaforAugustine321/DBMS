package interfaces

type Invoices struct {
	ItemName string    `json:"Name"`
	UserEmail string    `json:"userEmail"`
	Rate     int64     `json:"Rate"`
	Hour     int64     `json:"hour"`
}