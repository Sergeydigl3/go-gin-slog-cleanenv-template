package auth

import "log/slog"

type Service struct {
	log *slog.Logger
}

func NewService(log *slog.Logger) *Service {
	return &Service{
		log: log,
	}
}

func (s *Service) Login() {
	s.log.Info("login")
}

func (s *Service) Logout() {
	s.log.Info("logout")
}
