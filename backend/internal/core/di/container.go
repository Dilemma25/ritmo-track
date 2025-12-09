package di

import (
	"log"

	"go.uber.org/dig"
)

type Container struct {
	dig *dig.Container
}

func NewContainer() *Container {
	return &Container{
		dig: dig.New(),
	}
}

func (c *Container) Provide(constructor interface{}) *Container {
	if err := c.dig.Provide(constructor); err != nil {
		log.Fatalf("failed to provide constructor: %v", err)
	}

	return c
}

func (c *Container) Invoke(function interface{}) {
	if err := c.dig.Invoke(function); err != nil {
		log.Fatalf("failed to invoke function: %v", err)
	}
}
