package models

type Productos struct{
	ID int `json:"id"`
	Nombre string `json:"Nombre"`
	Cant int `json:"Cant"`
	CodigoBarras string `json:"CodigoBarras"`
}

