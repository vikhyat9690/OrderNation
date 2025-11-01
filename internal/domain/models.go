package domain

import "time"

type Product struct {
	ID                int64
	Name              string
	Sku               string
	Short_Description string
	Long_Description  string
	Price             float32
	Special_Price     float32
	Created_At        time.Time
	Updated_At        time.Time
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
