package usecase

import "github.com/israelalvesmelo/CleanArch/internal/entity"

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(OrderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrdersUseCase) Execute() ([]OrderOutputDTO, error) {
	orders, err := c.OrderRepository.FindAll()
	if err != nil {
		return nil, err
	}

	dtoList := []OrderOutputDTO{}
	for _, order := range orders {
		dtoList = append(dtoList, OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return dtoList, nil
}
