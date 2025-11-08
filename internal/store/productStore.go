package store

import (
	"database/sql"
	"errors"
	"fmt"
	"ordernationn/internal/domain"
	"time"
)

type ProductStore struct {
	products map[int64]domain.Product
	nextId   int64
}

type SQLProductRepository struct {
	DB *sql.DB
}

func NewSQLProductRepository(db *sql.DB) *SQLProductRepository {
	return &SQLProductRepository{
		DB: db,
	}
}

// func NewProductStore() *ProductStore {
// 	repo := &ProductStore{
// 		products: make(map[int64]domain.Product),
// 		nextId:   1,
// 	}
// 	// Seed initial data
// 	repo.Add(domain.Product{
// 		Name: "Go T-Shirt", Sku: "GT-001", Short_Description: "A cool Gopher shirt.",
// 		Long_Description: "Official Go programming language t-shirt.",
// 		Price:            25.00, Special_Price: 20.00,
// 	})
// 	return repo
// }

// Find all implements the ProductRepository interface
func (r *SQLProductRepository) FindAll() ([]domain.Product, error) {
	columnList := `id, sku, name, short_description, long_description, price, special_price, created_at, updated_at, base_image_url`
	query := fmt.Sprintf(`SELECT %s FROM "Product"`, columnList)
	rows, err := r.DB.Query(query)
	if err != nil {
		return []domain.Product{}, err
	}
	defer rows.Close()

	products := []domain.Product{}
	for rows.Next() {
		var p domain.Product
		err := rows.Scan(
			&p.ID, &p.Sku, &p.Name, &p.Short_Description, &p.Long_Description, &p.Price, &p.Special_Price,
			&p.Created_At, &p.Updated_At, &p.Base_Image_Url,
		)
		if err != nil {
			return []domain.Product{}, err
		}
		products = append(products, p)
	}
	return products, rows.Err()
}

// Find by id implements the ProductRepository interface.
func (r *SQLProductRepository) FindById(id int64) (domain.Product, error) {
	columnList := `id, sku, name, short_description, long_description, price, special_price, created_at, updated_at, base_image_url`
	query := fmt.Sprintf(`SELECT %s FROM "Product" WHERE id=$1`, columnList)
	resultProduct := r.DB.QueryRow(query, id)
	var p domain.Product
	scanErr := resultProduct.Scan(&p.ID, &p.Sku, &p.Name, &p.Short_Description, &p.Long_Description, &p.Price, &p.Special_Price,
		&p.Created_At, &p.Updated_At, &p.Base_Image_Url)
	if errors.Is(scanErr, sql.ErrNoRows) {
		return domain.Product{}, nil
	}
	if scanErr != nil {
		return domain.Product{}, scanErr
	}
	return p, nil
}

// // Add product implementation
func (r *SQLProductRepository) Add(product domain.Product) (domain.Product, error) {
	now := time.Now()
	query := `
		INSERT INTO "Product"
			(sku, name, short_description, long_description, price, special_price, created_at, updated_at, base_image_url)
		VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, created_at, updated_at
	`
	err := r.DB.QueryRow(
		query,
		product.Sku,
		product.Name,
		product.Short_Description.String,
		product.Long_Description.String,
		product.Price,
		product.Special_Price.Float64,
		now,
		now,
		product.Base_Image_Url.String,
	).Scan(&product.ID, &product.Created_At, &product.Updated_At)

	if err != nil {
		return domain.Product{}, err
	}
	product.Created_At = now
	product.Updated_At = now

	return product, nil
}

// // Update implements the ProductRepository interface
// func (r *ProductStore) Update(product domain.Product) (domain.Product, error) {
// 	_, exists := r.products[product.ID]
// 	if !exists {
// 		return domain.Product{}, nil
// 	}
// 	product.Created_At = r.products[product.ID].Created_At
// 	product.Updated_At = time.Now()
// 	r.products[product.ID] = product

// 	return product, nil
// }
