package service

import (
	"errors"
	"log"
	"ordernationn/internal/domain"
)

type ProductService interface {
	GetAll() ([]domain.Product, error)
	GetById(id int64) (domain.Product, error)
	Create(product domain.Product) (domain.Product, error)
	// Update(product domain.Product) (domain.Product, error)
}

type productService struct {
	repo domain.ProductRepository
}

func NewProductService(r domain.ProductRepository) ProductService {
	return &productService{
		repo: r,
	}
}

func (s *productService) GetAll() ([]domain.Product, error) {
	return s.repo.FindAll()
}

func (s *productService) GetById(id int64) (domain.Product, error) {
	if id <= 0 {
		return domain.Product{}, errors.New("invalid product id")
	}
	product, err := s.repo.FindById(id)
	if product.ID == 0 {
		log.Println(err)
		return domain.Product{}, errors.New("no valid product found")
	}
	return product, nil
}

func (s *productService) Create(p domain.Product) (domain.Product, error) {
	if p.Sku == "" || p.Name == "" || p.Price == 0 {
		return domain.Product{}, errors.New("sku, price and the name are required")
	}
	if p.Special_Price.Valid {
		if p.Special_Price.Float64 <= 0 {
			return domain.Product{}, errors.New("special price must be greater than 0")
		}
		if float32(p.Special_Price.Float64) >= p.Price {
			return domain.Product{}, errors.New("special price cannot be greater than or equal to price")
		}
	}
	createdProduct, err := s.repo.Add(p)
	if err != nil {
		return domain.Product{}, err
	}

	if createdProduct.ID == 0 {
		return domain.Product{}, errors.New("failed to create product")
	}

	return p, nil
}

// func (s *productService) GetById(id int64) (domain.Product, error) {
// 	if id <= 0 {
// 		return domain.Product{}, errors.New("invalid product id")
// 	}
// 	product, err := s.repo.FindById(id)
// 	if err != nil {
// 		return domain.Product{}, err
// 	}
// 	if product.ID == 0 {
// 		return domain.Product{}, errors.New("no valid product found")
// 	}
// 	return product, nil
// }

// func (s *productService) Create(product domain.Product) (domain.Product, error) {
// 	if product.Special_Price > 0 && product.Special_Price >= product.Price {
// 		return domain.Product{}, errors.New("special price cannot be greater than fixed price")
// 	}
// 	return s.repo.Add(product)
// }

// func (s *productService) Update(product domain.Product) (domain.Product, error) {
// 	exists, _ := s.repo.FindById(product.ID)
// 	if exists.ID == 0 {
// 		return domain.Product{}, errors.New("no product exists")
// 	}
// 	return s.repo.Update(product)
// }
