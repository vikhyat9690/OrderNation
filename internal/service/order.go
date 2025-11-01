package service

import (
	"context"
	"errors"
	"ordernationn/internal/domain"
	"ordernationn/internal/store"
)

type OrderService struct {
	Store *store.Store
}

func (s *OrderService) PlaceOrder(ctx context.Context, order domain.Order) (string, error) {
	for _, item := range order.Items {
		qty, err := s.Store.GetStock(ctx, order.Outlet_Id, item.Product_Id)
		if err != nil {
			return "", err
		}
		if qty[1] < item.Quantity {
			return "", errors.New("product quantity is not available")
		}
	}
	return "Order is Placed Successfully", nil
}
