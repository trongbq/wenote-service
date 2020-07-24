package operation

import "github.com/sirupsen/logrus"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) SaveOperations(ops []Operation) {
	logrus.Info(ops)
}
