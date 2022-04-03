package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type cuisine struct {
	Name string
	// Continent string
	// Region    string
	Dishes foodItems
}

type foodItems []string

// TODO: replace this with a database
// this is the default list of cuisines or styles of food
var cuisines = []cuisine{
	{Name: "North American", Dishes: foodItems{"Burgers", "Fired Chicken"}},
	{Name: "South American", Dishes: foodItems{"Burritos", "Tacos", "Quesadillas"}},
	{Name: "Chinese", Dishes: foodItems{"Orange Chicken", "General Tso's Chicken", "Beijing Beef", "Ham Fried Rice"}},
	{Name: "Indian", Dishes: foodItems{"Chicken Tikka Masala", "Naan", "Kofta"}},
}

func main() {
	router := gin.Default()
	router.GET("/cuisines", getCuisines)

	router.Run("localhost:8080")
}

// getCuisines responds with the list of cuisines and
// their dishes as JSON
func getCuisines(c *gin.Context) {
	// Note that you can replace Context.IndentedJSON with a
	// call to Context.JSON to send more compact JSON. In
	// practice, the indented form is much easier to work
	// with when debugging and the size difference is
	// usually small.
	c.IndentedJSON(http.StatusOK, cuisines)
}
