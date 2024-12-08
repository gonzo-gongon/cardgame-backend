package app

import "go.uber.org/dig"

func BuildContainer() *dig.Container {
	container := dig.New()

	return container
}
