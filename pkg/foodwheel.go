package foodwheel

import (
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type Cuisine struct {
	Name string `json:"name"`
	// Continent string
	// Region    string
	Dishes []string `json:"dishes"`
}

type Cuisines []Cuisine

var Lock sync.Mutex

// TODO: replace this with a database
// this is the default list of cuisines or styles of food
var cuisines = []Cuisine{
	{Name: "North_American", Dishes: []string{"Burgers", "Fired Chicken"}},
	{Name: "South_American", Dishes: []string{"Burritos", "Tacos", "Quesadillas"}},
	{Name: "Chinese", Dishes: []string{"Orange Chicken", "General Tso's Chicken", "Beijing Beef", "Ham Fried Rice"}},
	{Name: "Indian", Dishes: []string{"Chicken Tikka Masala", "Naan", "Kofta"}},
}

// getCuisines responds with the list of cuisines and
// their dishes as JSON
func getCuisines(c *gin.Context) {
	// Note that you can replace Context.IndentedJSON with a
	// call to Context.JSON to send more compact JSON. In
	// practice, the indented form is much easier to work
	// with when debugging and the size difference is
	// usually small.
	Lock.Lock()
	c.IndentedJSON(http.StatusOK, cuisines)
	Lock.Unlock()
}

func getCuisineByName(c *gin.Context) {
	name := c.Param("name")

	// TODO: there is a better way to do this with a map
	// Loop over cuisines and find cuisine
	// whose name matches parameter
	Lock.Lock()
	defer Lock.Unlock()
	for _, cuisine := range cuisines {
		if cuisine.Name == name {
			c.IndentedJSON(http.StatusOK, cuisine)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "cuisine not found"})
}

// postCuisines add a new cuisines from JSON in the
// received request body
func postCuisines(c *gin.Context) {
	var newCuisine Cuisine

	// BindJSON binds the received JSON to
	// newCuisine
	if err := c.BindJSON(&newCuisine); err != nil {
		return
	}

	// Add the new cuisine to the list
	// remember, since this is stored in memory
	// changes are lost on restart of the container
	Lock.Lock()
	cuisines = append(cuisines, newCuisine)
	Lock.Unlock()

	c.IndentedJSON(http.StatusCreated, newCuisine)
}

func getSpin(c *gin.Context) {
	//nolint:gosec // doesn't need to be an unpredictable random number
	randNum := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(cuisines))

	Lock.Lock()
	c.IndentedJSON(http.StatusOK, cuisines[randNum])
	Lock.Unlock()
}
