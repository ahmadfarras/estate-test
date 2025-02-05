package model

type Tree struct {
	ID        string `json:"id"`
	EstateID  string `json:"estate_id"`
	Height    int    `json:"height"`
	X         int    `json:"x"`
	Y         int    `json:"y"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
