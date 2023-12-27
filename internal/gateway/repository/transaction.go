package repository

import (
	"context"
	"fmt"
	"prel/internal/gateway/postgresql"
	"prel/pkg/logger"

	"github.com/cockroachdb/errors"
	"github.com/jackc/pgx/v5"
)

type TransactionManager struct{}

func NewTransactionManager() *TransactionManager {
	return &TransactionManager{}
}

func (tm *TransactionManager) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	conn := postgresql.GetConn()
	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	queries := postgresql.New(tx)
	ctxWithQueries := postgresql.SetQueries(ctx, queries)

	err = fn(ctxWithQueries)
	if err != nil {
		logger.Get(ctx).Error(fmt.Sprintf("tx err %s", err))
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			err = errors.Wrap(err, "failed to rollback transaction")
			return errors.Wrap(rbErr, err.Error())
		}
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}

	return nil
}
