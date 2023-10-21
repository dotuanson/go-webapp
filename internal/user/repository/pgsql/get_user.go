package pgsql

import (
	"context"
	"go-webapp/internal/user/entity"
)

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	u := entity.User{}
	query := "SELECT id, email, username, password FROM users WHERE email = $1"
	err := r.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.Email, &u.Username, &u.Password)
	if err != nil {
		return &entity.User{}, nil
	}

	return &u, nil
}
