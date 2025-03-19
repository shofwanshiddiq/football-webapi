package models

type Player struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
	Club     string `json:"club"`
}
