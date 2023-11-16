package rpc

import (
	db "go-webapp/db/sqlc"
	"go-webapp/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.PaymentUser) *pb.User {
	return &pb.User{
		Username:         user.Username,
		FullName:         user.FullName,
		Email:            user.Email,
		PasswordChangeAt: timestamppb.New(user.PasswordChangedAt),
		CreateAt:         timestamppb.New(user.CreatedAt),
	}
}
