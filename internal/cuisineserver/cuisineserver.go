package cuisineserver

import cuisinesv1 "github.com/tydanny/foodwheel/gen/foodwheel/cuisines/v1"

type CuisineServer struct {
	cuisinesv1.UnimplementedCuisineServiceServer
}
