package pgsql

import (
	"context"
	"go-webapp/internal/user/entity"
)

func (r *repository) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	var lastInsertId int
	query := "INSERT INTO users(username, password, email) VALUES ($1, $2, $3) returning id"
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Password, user.Email).Scan(&lastInsertId)
	if err != nil {
		return &entity.User{}, err
	}

	user.ID = int64(lastInsertId)
	return user, nil
}
