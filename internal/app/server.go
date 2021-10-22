package app

import "github.com/abhinav812/cloudy-bookstore/internal/util/logger"

// Server - application representation wrapping a logger.Logger
type Server struct {
	logger *logger.Logger
}

// New - Creates a new Server with supplied logger.Logger
func New(l *logger.Logger) *Server {
	return &Server{
		logger: l,
	}
}

// Logger - return logger.Logger instance associated with the Server
func (s *Server) Logger() *logger.Logger {
	return s.logger
}
