package task

import "wetodo/internal/storage"

type Repository interface {
}

type Service struct {
	r Repository
}

func NewService(r Repository) *Service {
	return &Service{r}
}

func (s *Service) GetAllTasksByUser(userID int, completed bool, page int, limit int, sortField string, sortOrder string) []storage.Task {
	return []storage.Task{}
}
