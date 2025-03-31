package models

import "time"

type Productos struct {
	ID           int       `json:"id"`
	Nombre       string    `json:"Nombre"`
	Cant         int       `json:"Cant"`
	CodigoBarras string    `json:"CodigoBarras"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}
