package bun

import "fmt"

// Config represents the configuration settings for connecting to a database.
type Config struct {
	// Host is the database server hostname or IP address.
	Host string

	// Port is the port number on which the database server is listening.
	Port string

	// User is the username for authenticating with the database.
	User string

	// Password is the password for authenticating with the database.
	Password string

	// Name is the name of the database to connect to.
	Name string

	// Params is a map of additional connection parameters.
	Params map[string]string
}

func (c Config) Validate() error {
	if c.Host == "" {
		return ErrMissingHost
	}

	if c.Port == "" {
		return ErrMissingPort
	}

	if c.User == "" {
		return ErrMissingUser
	}

	if c.Name == "" {
		return ErrMissingName
	}

	return nil
}

func (c Config) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s %s",
		c.Host, c.Port, c.User, c.Password, c.Name, c.ParamsString(),
	)
}

func (c Config) ParamsString() string {
	var params string

	for k, v := range c.Params {
		params += fmt.Sprintf("%s=%s ", k, v)
	}

	return params
}
