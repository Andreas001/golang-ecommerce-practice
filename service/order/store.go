package order

import (
	"database/sql"

	"github.com/Andreas001/golang-ecommerce/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateOrder(order types.Order) (int, error) {
	res, err := s.db.Exec("INSERT INTO orders (userId, total, status, address) VALUES (?,?,?,?)", order.UserID, order.Total, order.Status, order.Address)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *Store) CreateOrderItem(OrderItem types.OrderItem) error {
	_, err := s.db.Exec("INSERT INTO order_items (orderId, productId, quantity, price) VALUES (?,?,?,?)", OrderItem.OrderID, OrderItem.ProductID, OrderItem.Quantity, OrderItem.Price)
	if err != nil {
		return err
	}

	return nil
}
