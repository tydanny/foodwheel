# Food Wheel

This is an RESTful API build with GoLang and the [Gin Web Framework](https://gin-gonic.com/docs/) to determine what kind of food to eat. This was built using the following [tutorial](https://go.dev/doc/tutorial/web-service-gin).

## Motivation

This is designed to help relieve users of the [analysis paralysis](https://en.wikipedia.org/wiki/Analysis_paralysis) when deciding what to eat for lunch or dinner.

## Endpoints

### /cuisines

- `GET` - Get the list of all cuisines and their dishes
- `POST` - Add a new cuisine or style of food from request data sent as JSON

### /cuisines/:name

- `GET` - Get a cuisine by its Name, returning its dishes as JSON

### /spin

- `GET` - Get a randomly determined style of food, returned as JSON

> **Note:** For now cuisines are stored in memory for simplicity. A better design would be to use a database.
