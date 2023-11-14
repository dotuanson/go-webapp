package rpc

import (
	"fmt"
	db "go-webapp/db/sqlc"
	"go-webapp/pb"
	"go-webapp/token"
	"go-webapp/util"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedGoWebAppServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	return server, err
}
