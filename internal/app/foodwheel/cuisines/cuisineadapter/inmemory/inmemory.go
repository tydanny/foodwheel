package inmemory

import (
	"context"
	"iter"
	"maps"
	"sync"

	cuisinesv1 "github.com/tydanny/foodwheel/gen/foodwheel/cuisines/v1"
	"github.com/tydanny/foodwheel/internal/app/foodwheel/cuisines"
)

func NewStore() *Store {
	store := &Store{}

	store.init()

	return store
}

type Store struct {
	cuisines  map[string]*cuisinesv1.Cuisine
	storeOnce sync.Once
	lock      sync.RWMutex
}

var _ cuisines.Service = (*Store)(nil)

func (s *Store) init() {
	s.storeOnce.Do(func() {
		s.cuisines = make(map[string]*cuisinesv1.Cuisine)
	})
}

func (s *Store) Get(ctx context.Context, name string) (*cuisinesv1.Cuisine, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	cuisine, ok := s.cuisines[name]
	if !ok {
		return nil, cuisines.NotFound(name)
	}

	return cuisine, nil
}

func (s *Store) List(ctx context.Context, offset int, pageSize int) ([]*cuisinesv1.Cuisine, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	nextCuisine, stopIter := iter.Pull(maps.Values(s.cuisines))
	defer stopIter()

	var cuisines []*cuisinesv1.Cuisine

	cuisinesRemaining := len(s.cuisines) - offset
	if cuisinesRemaining > pageSize {
		cuisines = make([]*cuisinesv1.Cuisine, pageSize)
	} else {
		cuisines = make([]*cuisinesv1.Cuisine, cuisinesRemaining)
	}

	// Skip to the offset.
	for range offset {
		_, ok := nextCuisine()
		if !ok {
			break
		}
	}

	// Pull the next cuisines.
	for index := range cuisines {
		cuisine, ok := nextCuisine()
		if !ok {
			break
		}

		cuisines[index] = cuisine
	}

	return cuisines, nil
}

func (s *Store) Create(
	ctx context.Context,
	cuisine *cuisinesv1.Cuisine,
) (*cuisinesv1.Cuisine, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if _, ok := s.cuisines[cuisine.GetName()]; ok {
		return nil, cuisines.AlreadyExists(cuisine.GetName())
	}

	s.cuisines[cuisine.GetName()] = cuisine

	return cuisine, nil
}

func (s *Store) Delete(ctx context.Context, name string) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if _, ok := s.cuisines[name]; !ok {
		return cuisines.NotFound(name)
	}

	delete(s.cuisines, name)

	return nil
}

func (s *Store) Update(
	ctx context.Context,
	cuisine *cuisinesv1.Cuisine,
) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if _, ok := s.cuisines[cuisine.GetName()]; !ok {
		return cuisines.NotFound(cuisine.GetName())
	}

	s.cuisines[cuisine.GetName()] = cuisine

	return nil
}
