package cuisineadapter

import "github.com/uptrace/bun"

func NewBunStore(bunDB *bun.DB) bunStore {
	return bunStore{bunDB}
}

type bunStore struct {
	bunDB *bun.DB
}
