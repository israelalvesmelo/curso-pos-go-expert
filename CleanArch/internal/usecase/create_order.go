package usecase

import (
	"github.com/israelalvesmelo/CleanArch/internal/entity"
	"github.com/israelalvesmelo/CleanArch/pkg/events"
)

type OrderInputDTO struct {
	ID    string  `json:"id"`
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

type OrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type CreateOrderUseCase struct {
	OrderRepository   entity.OrderRepositoryInterface
	EventOrderCreated events.EventInterface
	EventDispatcher   events.EventDispatcherInterface
}

func NewCreateOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	EventOrderCreated events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		OrderRepository:   OrderRepository,
		EventOrderCreated: EventOrderCreated,
		EventDispatcher:   EventDispatcher,
	}
}

func (c *CreateOrderUseCase) Execute(input OrderInputDTO) (OrderOutputDTO, error) {
	order, err := entity.NewOrder(input.ID, input.Price, input.Tax)
	if err != nil {
		return OrderOutputDTO{}, err
	}
	err = order.CalculateFinalPrice()
	if err != nil {
		return OrderOutputDTO{}, err
	}

	dto := OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}

	c.EventOrderCreated.SetPayload(dto)
	c.EventDispatcher.Dispatch(c.EventOrderCreated)

	return dto, nil
}
