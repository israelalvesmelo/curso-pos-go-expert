package product

type ProductUsecase struct {
	repository ProductRepositoryInterface
}

func NewProductUsecase(repository ProductRepositoryInterface) *ProductUsecase {
	return &ProductUsecase{repository: repository}
}

// GetProduct returns a product by ID.
// This product was not supposed to be returned. We should return a DTO instead.
// However, we will return it for now to keep the example simple.
func (u *ProductUsecase) GetProduct(id int) (Product, error) {
	return u.repository.GetProduct(id)
}
