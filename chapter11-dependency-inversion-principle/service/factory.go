package service

type ServiceFactory interface {
	MakeService() Service
}

type ServiceFactoryImpl struct{}

func (f *ServiceFactoryImpl) MakeService() Service {
	return &ConcreteImpl{}
}
