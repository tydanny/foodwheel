package server

// import (
// 	"context"
// 	"errors"
// 	"fmt"
// 	"math/rand"
// 	"sync"
// 	"time"
// )

// type FoodwheelServer struct {
// 	// foodwheel.FoodwheelServer

// 	mu       sync.Mutex
// 	Cuisines map[string]*foodwheel.Cuisine
// }

// func (s *FoodwheelServer) GetCuisines(e *foodwheel.Empty, stream foodwheel.Foodwheel_GetCuisinesServer) error {
// 	s.mu.Lock()
// 	for _, v := range s.Cuisines {
// 		if err := stream.Send(v); err != nil {
// 			return err
// 		}
// 	}
// 	s.mu.Unlock()
// 	return nil
// }

// func (s *FoodwheelServer) GetCuisineByName(
// 	ctx context.Context,
// 	req *foodwheel.CuisineRequest,
// ) (*foodwheel.Cuisine, error) {
// 	if c, ok := s.Cuisines[req.GetName()]; ok {
// 		return c, nil
// 	}
// 	return nil, fmt.Errorf("requested cuisine \"%s\" not found", req.GetName())
// }

// func (s *FoodwheelServer) AddCuisine(ctx context.Context, c *foodwheel.Cuisine) (*foodwheel.Cuisine, error) {
// 	s.Cuisines[c.GetName()] = c
// 	return c, nil
// }

// func (s *FoodwheelServer) Spin(context.Context, *foodwheel.Empty) (*foodwheel.Cuisine, error) {
// 	// It doen't matter if this random number is insecure or not
// 	randNum := rand.New(rand.NewSource(time.Now().Unix())).Intn(len(s.Cuisines)) //nolint:gosec

// 	for _, v := range s.Cuisines {
// 		if randNum == 0 {
// 			return v, nil
// 		}
// 		randNum--
// 	}
// 	return nil, errors.New("failed to retrieve a random cuisine")
// }

// func ExampleServer() *FoodwheelServer {
// 	s := FoodwheelServer{Cuisines: make(map[string]*foodwheel.Cuisine)}
// 	s.Cuisines["North_American"] = &foodwheel.Cuisine{
// 		Name:   "North_American",
// 		Dishes: []string{"Burgers", "Fried Chicken"},
// 	}
// 	s.Cuisines["South_American"] = &foodwheel.Cuisine{
// 		Name:   "South_American",
// 		Dishes: []string{"Burritos", "Tacos", "Quesadillas"},
// 	}
// 	s.Cuisines["Chinese"] = &foodwheel.Cuisine{
// 		Name:   "Chinese",
// 		Dishes: []string{"Orange Chicken", "General Tso's Chicken", "Beijing Beef", "Ham Fried Rice"},
// 	}
// 	s.Cuisines["Indian"] = &foodwheel.Cuisine{
// 		Name:   "Indian",
// 		Dishes: []string{"Chicken Tikka Masala", "Naan", "Kofta"},
// 	}
// 	return &s
// }

// // var exampleCuisines = []byte(`[{
// // 	"Name": "North_American",
// // 	"Dishes": ["Burgers", "Fired Chicken"]
// // }, {
// // 	"Name": "South_American",
// // 	"Dishes": ["Burritos", "Tacos", "Quesadillas"]
// // }, {
// // 	"Name": "Chinese",
// // 	"Dishes": ["Orange Chicken", "General Tso's Chicken", "Beijing Beef", "Ham Fried Rice"]
// // }, {
// // 	"Name": "Indian",
// // 	"Dishes": ["Chicken Tikka Masala", "Naan", "Kofta"]
// // }]`)
