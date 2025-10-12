package entities

type Professional struct {
	ID       string  `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	City     string `json:"city"`
	Contact  string `json:"contact"`
}
