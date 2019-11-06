package service

import (
	gdig "github.com/isayme/go-gdig"
)

func init() {
	constructors := []interface{}{
		NewMongo,
	}

	for _, constructor := range constructors {
		gdig.Provide(constructor)
	}
}
