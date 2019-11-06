package conf

import (
	gdig "github.com/isayme/go-gdig"
)

func init() {
	constructors := []interface{}{
		Get,
	}

	for _, constructor := range constructors {
		gdig.Provide(constructor)
	}
}
