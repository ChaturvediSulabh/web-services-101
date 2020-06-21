# WEB Services - 101

## INTRODCTION

This repo is part of my committment to code more and learn/imporve at Go/My programming skills.

In this project I'm trying to build a Donut store by implementing REST API.

### DEV STAGES

- Pass Data as JSON File
- Implement Data Persistence : Use PostGres DB
- Refractor Use Advance Go techniques sunc as context.

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

## USAGE

`docker run -d --name go-web-services-101 -p 5000:5000 chaturvedisulabh/go-web-services-101:latest -PORT=5000 -DB_CONN_STR=$DB_CONN_STR`

### Configurable parameters

- PORT = <YOUR_HOST_PORT> [update `docker run -p` param accordingly]
- DB_CONN_STR = <POSTGRES_DB_CONNECTION_STRING>