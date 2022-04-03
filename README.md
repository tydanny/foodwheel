# Food Wheel

This is an RESTful API build with GoLang and the [Gin Web Framework](https://gin-gonic.com/docs/) to determine what kind of food to eat.

## Motivation

This is designed to help relieve users of the [analysis paralysis](https://en.wikipedia.org/wiki/Analysis_paralysis) when deciding what to eat for lunch or dinner.

## Endpoints

### /cuisine

- `GET` - Get a randomly determined style of food from the default list, returned as JSON
- `POST` - Add a new cuisine or style of food from request data sent as JSON
