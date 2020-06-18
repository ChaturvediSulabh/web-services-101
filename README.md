# WEB Services - 101

## INTRODCTION

This repo is part of my committment to code more and learn/imporve at Go/My programming skills.

In this project I'm trying to build a Donut store by implementing REST API.

## BEHAVIOUR

1. Any User should be able to list all available toppings at `/api/toppings`
2. Any User should be able to get the name of the donut and its cost at `/api/topping/{TOPPING_NAME}`
   - if the topping name isn't available in the menu, user shall receive with a `HTTP 404 Status`.
3. Any User should be able to POST a new menu item in a valid JSON format at `/api/toppings`. See example below.
   - A successful HTTP POST must receive `HTTP 201 Status`.

```
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
```
