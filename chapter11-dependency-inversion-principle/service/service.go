package service

import "fmt"

type Service interface {
	Serve() error
}

type ConcreteImpl struct{}

func (c *ConcreteImpl) Serve() error {
	fmt.Println("This is a concrete implementation of service interface")
	return nil
}
