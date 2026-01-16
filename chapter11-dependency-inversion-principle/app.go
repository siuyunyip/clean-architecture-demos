package main

import "siuyunyip.io/dip-demo/service"

type Application struct {
	factory service.ServiceFactory
	service service.Service
}

func NewApplication(factory service.ServiceFactory) *Application {
	return &Application{
		factory: factory,
		service: factory.MakeService(),
	}
}

func main() {
	app := NewApplication(&service.ServiceFactoryImpl{})
	if err := app.service.Serve(); err != nil {
		panic(err)
	}
}
