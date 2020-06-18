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

/*
{
			"id": "0003",
			"type": "donut",
			"name": "Old Fashioned",
			"ppu": 0.55,
			"topping": [
				{
					"id": "5001",
					"type": "None"
				},
				{
					"id": "5002",
					"type": "Glazed"
				},
				{
					"id": "5003",
					"type": "Chocolate"
				},
				{
					"id": "5004",
					"type": "Maple"
				}
			]
		}
*/
