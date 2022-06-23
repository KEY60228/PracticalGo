package main

import (
	"context"
)

type UserService struct {
	tx txAdmin
}

func NewUserService(tx txAdmin) *UserService {
	return &UserService{
		tx: tx,
	}
}

func (s *UserService) UpdateName(ctx context.Context, userID string, userName string) error {
	updateFunc := func(ctx context.Context) error {
		if _, err := s.tx.ExecContext(ctx, "UPDATE users SET user_name = $1 WHERE user_id = $2", userName, userID); err != nil {
			return err
		}
		return nil
	}
	return s.tx.Transaction(ctx, updateFunc)
}
