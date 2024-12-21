package cuisine

import (
	"slices"
)

type Cuisine struct {
	Name   string
	Dishes []Dish
}

// NewCuisine creates a new cuisine
func NewCuisine(name string, dishes ...Dish) Cuisine {
	return Cuisine{name, dishes}
}

type Dish struct {
	Name string
	Tags []string
}

// NewDish creates a new dish
func NewDish(name string, tags ...string) Dish {
	slices.Sort(tags)

	return Dish{name, tags}
}

// HasTag checks if a dish has a specific tag
// The tags field should be sorted at creation
// so we can perform a binary search without sorting
// first
func (d *Dish) HasTag(tag string) bool {
	_, hasTag := slices.BinarySearch(d.Tags, tag)

	return hasTag
}
