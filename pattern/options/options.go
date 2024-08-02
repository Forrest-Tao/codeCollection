package options

import (
	"time"
)

type Server struct {
	Address      string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Option func(*Server)

func NewServer(opts ...Option) *Server {
	server := &Server{
		Address:      "localhost",
		Port:         8080,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	for _, opt := range opts {
		opt(server)
	}

	return server
}

// WithAddress配置服务器地址
func WithAddress(address string) Option {
	return func(s *Server) {
		s.Address = address
	}
}

func WithPort(port int) Option {
	return func(s *Server) {
		s.Port = port
	}
}

func WithReadTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.ReadTimeout = timeout
	}
}

func WithWriteTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.WriteTimeout = timeout
	}
}
