package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"purchasesvc/internal/domain"
	"strconv"
)

type OrdersRepo struct {
	db *sqlx.DB
}

func NewOrdersRepo(db *sqlx.DB) *OrdersRepo {
	return &OrdersRepo{db: db}
}

func (r *OrdersRepo) Create(ctx context.Context, order domain.Order, items map[string]string) (domain.Order, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return domain.Order{}, err
	}

	row := tx.QueryRowContext(ctx, "INSERT INTO orders (user_id, full_name, address, total_price) VALUES ($1, $2, $3, $4) RETURNING id", order.UserID, order.FullName, order.Address, order.TotalPrice)

	err = row.Scan(&order.ID)
	if err != nil {
		tx.Rollback()
		return domain.Order{}, err
	}

	for item, countString := range items {
		count, err := strconv.Atoi(countString)
		if err != nil {
			tx.Rollback()
			return domain.Order{}, err
		}
		for i := 0; i < count; i++ {
			_, err := tx.ExecContext(ctx, "INSERT INTO order_items (order_id, product_id) VALUES ($1, $2)", order.ID, item)
			if err != nil {
				tx.Rollback()
				return domain.Order{}, err
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return domain.Order{}, err
	}

	return order, nil
}
