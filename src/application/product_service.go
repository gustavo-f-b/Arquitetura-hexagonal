package application

type ProductService struct {
	ProductPersistence ProductPersistenceInterface
}

func (productService *ProductService) Get(id string) (ProductInterface, error) {
	product, err := productService.ProductPersistence.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}
