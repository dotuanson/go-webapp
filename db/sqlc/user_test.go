package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"go-webapp/util"
	"testing"
	"time"
)

func createRandomUser(t *testing.T) PaymentUser {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)
	userResult, err := testQueries.GetUser(context.Background(), user.Username)
	require.NoError(t, err)
	require.NotEmpty(t, userResult)

	require.Equal(t, user.Username, userResult.Username)
	require.Equal(t, user.HashedPassword, userResult.HashedPassword)
	require.Equal(t, user.FullName, userResult.FullName)
	require.Equal(t, user.Email, userResult.Email)
	require.WithinDuration(t, user.PasswordChangedAt, userResult.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user.CreatedAt, userResult.CreatedAt, time.Second)
}
