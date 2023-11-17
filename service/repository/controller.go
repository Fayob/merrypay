package repository

import "merrypay/port"

type Server struct {
	Server port.Store
}

func NewServer(server port.Store) *Server {
	return &Server{
		Server: server,
	}
}
