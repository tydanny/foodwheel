package cuisine

type Store interface {
	CreateCuisine(cuisine Cuisine) error
	GetCuisine(name string) (Cuisine, error)
	ListCuisines(pageSize int, offset int) ([]Cuisine, error)
	UpdateCuisine(cuisine Cuisine, paths ...string) error
	DeleteCuisine(name string) error
}
