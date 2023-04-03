package entity

type Client struct {
	ID      int64  `json:"ID"`
	Nombres string `json:"Nombres"`
	Email   string `json:"Email"`
}
