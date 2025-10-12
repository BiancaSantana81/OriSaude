package entities

type Professional struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	City     string `json:"city"`
	Contact  string `json:"contact"`
}
