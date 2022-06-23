package main

import (
	"context"
	"database/sql"
	"fmt"
)

type txAdmin struct {
	*sql.DB
}

func NewTxAdmin(db *sql.DB) *txAdmin {
	return &txAdmin{db}
}

func (t *txAdmin) Transaction(ctx context.Context, f func(ctx context.Context) error) error {
	tx, err := t.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := f(ctx); err != nil {
		return fmt.Errorf("transaction query failed: %w", err)
	}

	return tx.Commit()
}
