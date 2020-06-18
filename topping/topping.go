package topping

//Sample ...
type Sample struct {
	ID      string  `json:"id"`
	Type    string  `json:"type"`
	Name    string  `json:"name"`
	Ppu     float64 `json:"ppu"`
	Topping Topping `json:"topping"`
}

//Topping ...
type Topping []struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

//ToppingResponse ...
type ToppingResponse struct {
	Name string
	Cost float64
}
