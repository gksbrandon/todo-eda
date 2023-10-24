package config

import "errors"

type Environment struct {
	slug string
}

func (e Environment) String() string {
	return e.slug
}

var (
	Unknown = Environment{""}
	Dev     = Environment{"dev"}
	Stage   = Environment{"stage"}
	Prod    = Environment{"prod"}
)

func validateEnvironment(s string) error {
	switch s {
	case Dev.slug:
		return nil
	case Stage.slug:
		return nil
	case Prod.slug:
		return nil
	}

	return errors.New("Unknown Environment: " + s)
}
