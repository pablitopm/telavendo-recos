package registry

import (
	"github.com/sarulabs/di"
	"telavendo.com.ar/telavendo-recos/app/domain/service"
	"telavendo.com.ar/telavendo-recos/app/interface/persistence/dynamodb"
	"telavendo.com.ar/telavendo-recos/app/usecase"
)

type Container struct {
	ctn di.Container
}

func NewContainer() (*Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}

	if err := builder.Add([]di.Def{
		{
			Name:  "order-usecase",
			Build: buildOrderUsecase,
		},
		{
			Name:  "user-purchase-usecase",
			Build: buildUserPurchaseUsecase,
		},
	}...); err != nil {
		return nil, err
	}

	return &Container{
		ctn: builder.Build(),
	}, nil
}

func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

func (c *Container) Clean() error {
	return c.ctn.Clean()
}

func buildUserPurchaseUsecase(ctn di.Container) (interface{}, error) {
	repo := dynamodb.NewUserPurchaseRepository()
	service := service.NewUserPurchaseService(repo)
	return usecase.NewUserPurchaseUsecase(repo, service), nil
}
func buildOrderUsecase(ctn di.Container) (interface{}, error) {
	return usecase.NewOrderUsecase(), nil
}
