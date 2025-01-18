package pgcuisine

import (
	"github.com/tydanny/foodwheel/pkg/cuisine"
	"github.com/uptrace/bun"
)

var dsn = "postgres://postgres:@localhost:5432/test?sslmode=disable"

func Initialize() (*PGCuisineStore, error) {
	// sqldb := sql.OpenDB(pgdriver.NewConnector(
	// 	pgdriver.WithNetwork("tcp"),
	// 	pgdriver.WithAddr(viper.GetString("pg.dbaddr")),
	// 	pgdriver.WithUser(viper.GetString("pg.dbuser")),
	// 	pgdriver.WithPassword(viper.GetString("pg.dbpass")),
	// 	pgdriver.WithTimeout(5*time.Second),
	// ))

	// bun := bun.NewDB(sqldb, pgdialect.New())

	return nil, nil
}

type PGCuisineStore struct {
	bun.DB
}

var _ cuisine.Store = (*PGCuisineStore)(nil)

// CreateCuisine implements cuisine.Store.
func (p *PGCuisineStore) CreateCuisine(cuisine cuisine.Cuisine) error {
	panic("unimplemented")
}

// DeleteCuisine implements cuisine.Store.
func (p *PGCuisineStore) DeleteCuisine(name string) error {
	panic("unimplemented")
}

// GetCuisine implements cuisine.Store.
func (p *PGCuisineStore) GetCuisine(name string) (cuisine.Cuisine, error) {
	panic("unimplemented")
}

// ListCuisines implements cuisine.Store.
func (p *PGCuisineStore) ListCuisines(pageSize int, offset int) ([]cuisine.Cuisine, error) {
	panic("unimplemented")
}

// UpdateCuisine implements cuisine.Store.
func (p *PGCuisineStore) UpdateCuisine(cuisine cuisine.Cuisine, paths ...string) error {
	panic("unimplemented")
}
