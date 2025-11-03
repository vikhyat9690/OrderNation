package domain

import (
	"database/sql"
	"time"
)

type Product struct {
	ID                int64           `json:"id"`
	Name              string          `json:"name"`
	Sku               string          `json:"sku"`
	Short_Description sql.NullString  `json:"short_description"`
	Long_Description  sql.NullString  `json:"long_description"`
	Price             float32         `json:"price"`
	Special_Price     sql.NullFloat64 `json:"special_price"`
	Created_At        time.Time       `json:"created_at"`
	Updated_At        time.Time       `json:"updated_at"`
}

type ProductRepository interface {
	FindAll() ([]Product, error)
	FindById(id int64) (Product, error)
	// Add(product Product) (Product, error)
	// Update(product Product) (Product, error)
}

type Stock struct {
	ID               int64
	PRODUCT_ID       int64
	OUTLET_ID        int64
	IS_IN_STOCK      int
	Quantity         int64
	Salable_Quantity int64
}

type OrderItem struct {
	ID         int64
	Order_Id   int64
	Quantity   int64
	Product_Id int64
}

type Order struct {
	ID               int64
	Outlet_Id        int64
	Item_Quantity    int16
	Status           string
	Items            []OrderItem
	Total            float32
	Base_Grand_Total float32
	Base_SubTotal    float32
	Grand_Total      float32
	SubTotal         float32
	Base_Tax         float32
	Tax              float32
	Base_Discount    float32
	Discount         float32
	Created_At       time.Time
}
