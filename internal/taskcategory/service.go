package taskcategory

import "wetodo/internal/storage"

type Repository interface {
	GetAllTaskCategoriesByUser(userID int) []storage.TaskCategory
}

type Service struct {
	r Repository
}

func NewService(r Repository) *Service {
	return &Service{r}
}

func (s *Service) GetAllTaskCategoriesByUser(userID int) []storage.TaskCategory {
	return s.r.GetAllTaskCategoriesByUser(userID)
}
