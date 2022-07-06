package server

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"

	"github.com/tydanny/foodwheel/pkg/foodwheel"
	"google.golang.org/grpc"
)

type FoodwheelServer struct {
	foodwheel.UnimplementedFoodwheelServer

	mu       sync.Mutex
	Cuisines map[string]*foodwheel.Cuisine
}

var (
	// tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	// certFile   = flag.String("cert_file", "", "The TLS cert file")
	// keyFile    = flag.String("key_file", "", "The TLS key file")
	// jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of Cuisines")
	port = flag.Int("port", 50051, "The server port")
)

func (s *FoodwheelServer) GetCuisines(e *foodwheel.Empty, stream foodwheel.Foodwheel_GetCuisinesServer) error {
	s.mu.Lock()
	for _, v := range s.Cuisines {
		if err := stream.Send(v); err != nil {
			return err
		}
	}
	return nil
}

func (s *FoodwheelServer) GetCuisineByName(ctx context.Context, req *foodwheel.CuisineRequest) (*foodwheel.Cuisine, error) {
	if c, ok := s.Cuisines[req.GetName()]; ok {
		return c, nil
	}
	return nil, fmt.Errorf("requested cuisine \"%s\" not found", req.GetName())
}

func (s *FoodwheelServer) AddCuisine(ctx context.Context, c *foodwheel.Cuisine) (*foodwheel.Empty, error) {
	s.Cuisines[c.GetName()] = c
	return &foodwheel.Empty{}, nil
}

func (s *FoodwheelServer) Spin(context.Context, *foodwheel.Empty) (*foodwheel.Cuisine, error) {
	// It doen't matter if this random number is insecure or not
	randNum := rand.New(rand.NewSource(time.Now().Unix())).Intn(len(s.Cuisines)) //nolint:gosec

	for _, v := range s.Cuisines {
		if randNum == 0 {
			return v, nil
		}
		randNum--
	}
	return nil, errors.New("failed to retrieve a random cuisine")
}

func NewServer() *FoodwheelServer {
	s := FoodwheelServer{Cuisines: make(map[string]*foodwheel.Cuisine)}
	s.Cuisines["North_American"] = &foodwheel.Cuisine{
		Name:   "North_American",
		Dishes: []string{"Burgers", "Fried Chicken"},
	}
	s.Cuisines["South_American"] = &foodwheel.Cuisine{
		Name:   "South_American",
		Dishes: []string{"Burritos", "Tacos", "Quesadillas"},
	}
	s.Cuisines["Chinese"] = &foodwheel.Cuisine{
		Name:   "Chinese",
		Dishes: []string{"Orange Chicken", "General Tso's Chicken", "Beijing Beef", "Ham Fried Rice"},
	}
	s.Cuisines["Indian"] = &foodwheel.Cuisine{
		Name:   "Indian",
		Dishes: []string{"Chicken Tikka Masala", "Naan", "Kofta"},
	}
	return &s
}

func Main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	// if *tls {
	// 	if *certFile == "" {
	// 		// *certFile = data.Path("x509/server_cert.pem")
	// 	}
	// 	if *keyFile == "" {
	// 		// *keyFile = data.Path("x509/server_key.pem")
	// 	}
	// 	creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
	// 	if err != nil {
	// 		log.Fatalf("Failed to generate credentials %v", err)
	// 	}
	// 	opts = []grpc.ServerOption{grpc.Creds(creds)}
	// }

	grpcServer := grpc.NewServer(opts...)
	foodwheel.RegisterFoodwheelServer(grpcServer, NewServer())
	if servErr := grpcServer.Serve(lis); servErr != nil {
		log.Fatalf("server exited unexpectedly: %v", servErr)
	}
}

// var exampleCuisines = []byte(`[{
// 	"Name": "North_American",
// 	"Dishes": ["Burgers", "Fired Chicken"]
// }, {
// 	"Name": "South_American",
// 	"Dishes": ["Burritos", "Tacos", "Quesadillas"]
// }, {
// 	"Name": "Chinese",
// 	"Dishes": ["Orange Chicken", "General Tso's Chicken", "Beijing Beef", "Ham Fried Rice"]
// }, {
// 	"Name": "Indian",
// 	"Dishes": ["Chicken Tikka Masala", "Naan", "Kofta"]
// }]`)
