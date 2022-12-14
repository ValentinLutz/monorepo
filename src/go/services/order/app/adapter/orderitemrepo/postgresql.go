package orderitemrepo

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"monorepo/services/order/app/core/entity"
)

type PostgreSQL struct {
	db *sqlx.DB
}

func NewPostgreSQL(database *sqlx.DB) PostgreSQL {
	return PostgreSQL{db: database}
}

func (orderItemRepository *PostgreSQL) FindAllByOrderIds(orderIds []entity.OrderId) ([]entity.OrderItem, error) {
	rows, err := orderItemRepository.db.Query("SELECT order_item_id, order_id, creation_date, item_name FROM order_service.order_item WHERE order_id = ANY($1)", pq.Array(orderIds))
	if err != nil {
		return nil, err
	}

	return extractOrderItemEntities(rows)
}

func (orderItemRepository *PostgreSQL) FindAllByOrderId(orderId entity.OrderId) ([]entity.OrderItem, error) {
	rows, err := orderItemRepository.db.Query("SELECT order_item_id, order_id, creation_date, item_name FROM order_service.order_item WHERE order_id = $1", orderId)
	if err != nil {
		return nil, err
	}

	return extractOrderItemEntities(rows)
}

func (orderItemRepository *PostgreSQL) SaveAll(orderItemEntities []entity.OrderItem) error {
	_, err := orderItemRepository.db.NamedExec(
		"INSERT INTO order_service.order_item (order_id, creation_date, item_name) VALUES (:order_id, :creation_date, :item_name)", orderItemEntities)
	return err
}

func extractOrderItemEntities(rows *sql.Rows) ([]entity.OrderItem, error) {
	var orderItemEntities []entity.OrderItem
	for rows.Next() {
		var orderItemEntity entity.OrderItem

		err := rows.Scan(&orderItemEntity.OrderItemId, &orderItemEntity.OrderId, &orderItemEntity.CreationDate, &orderItemEntity.Name)
		if err != nil {
			return nil, err
		}

		orderItemEntities = append(orderItemEntities, orderItemEntity)
	}
	return orderItemEntities, nil
}
