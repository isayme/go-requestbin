package router

import (
	gdig "github.com/isayme/go-gdig"
)

func init() {
	constructors := []interface{}{
		NewRequest,
	}

	for _, constructor := range constructors {
		gdig.Provide(constructor)
	}
}
