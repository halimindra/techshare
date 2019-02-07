package models

type Person struct {
	ID      int      `json:"id,omitempty"`
	Name    string   `json:"name,omitempty"`
	Address *Address `json:"address,omitempty"`
}

type Address struct {
	City    string `json:"city,omitempty"`
	Country string `json:"country,omitempty"`
}
