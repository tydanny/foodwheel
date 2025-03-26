package cuisines

import "fmt"

func NotFound(name string) error {
	return NotFoundError{Name: name}
}

// NotFoundError is returned if an item cannot be found.
type NotFoundError struct {
	Name string
}

func (err NotFoundError) Error() string {
	return fmt.Sprintf("cuisine '%s' not found", err.Name)
}

func AlreadyExists(name string) error {
	return AlreadyExistsError{Name: name}
}

// AlreadyExistsError is returned if an item already exists.
type AlreadyExistsError struct {
	Name string
}

func (err AlreadyExistsError) Error() string {
	return fmt.Sprintf("cuisine '%s' already exists", err.Name)
}
