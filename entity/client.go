package entity

type Client struct {
	ID        int64  `json:"ID"`
	Nombres   string `json:"Nombres"`
	Apellidos string `json:"Apellidos"`
	DNI       string `json:"DNI"`
	Anio      string `json:"Anio"`
	Mes       string `json:"Mes"`
	Dia       string `json:"Dia"`
	Ciudad    string `json:"Ciudad"`
}
