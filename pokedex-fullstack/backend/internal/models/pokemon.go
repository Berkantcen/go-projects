package models

// Pokemon is a struct that represents a Pokemon
type Pokemon struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Image   string `json:"image"`
	Type    string `json:"type"`
	Creator string `json:"creator"`
}
