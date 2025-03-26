package cuisines

import (
	"context"

	cuisinesv1 "github.com/tydanny/foodwheel/gen/foodwheel/cuisines/v1"
)

// Service persists Cuisines.
type Service interface {
	Get(ctx context.Context, name string) (*cuisinesv1.Cuisine, error)
	List(ctx context.Context, offset, pageSize int) ([]*cuisinesv1.Cuisine, error)
	Create(ctx context.Context, cuisine *cuisinesv1.Cuisine) (*cuisinesv1.Cuisine, error)
	Update(ctx context.Context, cuisine *cuisinesv1.Cuisine) error
	Delete(ctx context.Context, name string) error
}
