package aplication

type ProductService struct {
	ProductRepository ProductRepository
}

func NewProductService(productRepository ProductRepository) *ProductService {
	return &ProductService{ProductRepository: productRepository}
}

func (s *ProductService) Get(Id string) (ProductInterface, error) {
	product, err := s.ProductRepository.FindById(Id)

	//se tiver um erro retorna  o nil
	if err != nil {
		return nil, err
	}

	return product, nil
}
func (s *ProductService) Create(name string, price float32) (ProductInterface, error) {
	product := NewProduct()
	product.Name = name
	product.Price = price

	_, err := product.IsValid()

	//se tiver erro retorn  um Producto vazio
	//nail e retornar um erro em branco

	if err != nil {
		return &Product{}, err
	}

	result, err := s.ProductRepository.Save(product)

	return result, nil
}

func (s *ProductService) Enable(product ProductInterface) (ProductInterface, error) {
	err := product.Enable()

	if err != nil {
		return &Product{}, err
	}

	result, err := s.ProductRepository.Save(product)

	return result, nil
}

func (s *ProductService) Disable(product ProductInterface) (ProductInterface, error) {
	err := product.Disable()

	if err != nil {
		return &Product{}, err
	}

	result, err := s.ProductRepository.Save(product)

	return result, nil
}
