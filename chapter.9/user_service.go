package main

import (
	"context"
	"database/sql"
	"fmt"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		db: db,
	}
}

type User struct {
	UserID   string
	UserName string
	// CreatedAt time.Time
}

func (s *UserService) FetchUser(ctx context.Context, userID string) (*User, error) {
	row := s.db.QueryRowContext(ctx, `SELECT user_id, user_name FROM users WHERE user_id = $1;`, userID)
	user, err := s.scanUser(row)
	if err != nil {
		return nil, fmt.Errorf("scan user: %w", err)
	}
	return user, nil
}

func (s *UserService) scanUser(row *sql.Row) (*User, error) {
	var u User
	err := row.Scan(&u.UserID, &u.UserName)
	if err != nil {
		return nil, fmt.Errorf("row scan: %w", err)
	}
	return &u, nil
}
