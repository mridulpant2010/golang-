package domain

type Pokemon struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Type      string  `json:"type"`
	Height    float64 `json:"height"`
	Weight    float64 `json:"weight"`
	Abilities string  `json:"abilities"`
}
