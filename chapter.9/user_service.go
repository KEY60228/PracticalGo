package main

import (
	"context"
	"database/sql"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (s *UserService) UpdateName(ctx context.Context, userID string, userName string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, "UPDATE users SET user_name = $1 WHERE user_id = $2", userName, userID); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
