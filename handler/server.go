package handler

import "github.com/SawitProRecruitment/UserService/repository"

type Server struct {
	Repository repository.UserRepositoryInterface
}

type NewServerOptions struct {
	Repository repository.UserRepositoryInterface
}

func NewServer(opts NewServerOptions) *Server {
	return &Server{
		Repository: opts.Repository,
	}
}
