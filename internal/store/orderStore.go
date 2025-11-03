package store

import (
	"context"
	"database/sql"
)

type Store struct {
	DB *sql.DB
}

func (s *Store) GetStock(ctx context.Context, outletId, productId int64) ([]int64, error) {
	var qty int64
	var salableQty int64
	err := s.DB.QueryRowContext(ctx,
		`SELECT quantity, salable_quantity FROM STOCK WHERE product_id=$1 and outlet_id = $2`,
		productId, outletId).Scan(&qty, &salableQty)

	if err == sql.ErrNoRows {
		return []int64{0, 0}, nil
	}
	if err != nil {
		return []int64{0, 0}, err
	}
	return []int64{qty, salableQty}, nil
}

func (s *Store) UpdateStock(ctx context.Context, outletId, productId, delta int64) (string, error) {
	_, err := s.DB.ExecContext(ctx,
		`UPDATE Stock SET salable_quantity = salable_quantity + $3 WHERE outlet_id=$1 AND product_id=$2`,
		outletId, productId, delta)
	if err != nil {
		return "", err
	}
	return "Stock Updated Successfully", nil
}
