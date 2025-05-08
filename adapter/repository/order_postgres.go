package repository

import (
	"context"
	"errors"

	"github.com/Vractos/ecoffe-go/entity"
	"github.com/Vractos/ecoffe-go/pkg/metrics"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type OrderPostgreSQL struct {
	db     *pgxpool.Pool
	logger metrics.Logger
}

func NewOrderPostgreSQL(db *pgxpool.Pool, logger metrics.Logger) *OrderPostgreSQL {
	return &OrderPostgreSQL{db: db, logger: logger}
}

// RegisterOrder implements order.Repository
func (r *OrderPostgreSQL) CreateOrder(o *entity.Order) error {
	ctx := context.Background()
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, `
  INSERT INTO orders(id, client, item, quantity, observation, status, created_at)
  VALUES($1,$2,$3,$4,$5, $6,$7)
  `, o.ID, o.Item, o.Client, o.Quantity, o.Observation, o.Status, o.CreatedAt)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			r.logger.Error(pgErr.Message, pgErr, zap.String("db_error_code", pgErr.Code))
		}
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		r.logger.Error("Error to commit order", err)
		return errors.New("error to commit order")
	}

	return nil
}

func (r *OrderPostgreSQL) RetrieveAllOrders() (*[]entity.Order, error) {
	panic("implement me")
}
