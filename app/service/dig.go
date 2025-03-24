package service

import (
	gdig "github.com/isayme/go-gdig"
)

func init() {
	constructors := []interface{}{
		NewMongo,
		NewSseServer,
	}

	for _, constructor := range constructors {
		gdig.Provide(constructor)
	}
}
