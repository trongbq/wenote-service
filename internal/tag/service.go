package tag

import "wetodo/internal/storage"

type Repository interface {
	GetAllTagsByUser(userID int) []storage.Tag
}

type Service struct {
	r Repository
}

func NewService(r Repository) *Service {
	return &Service{r}
}

func (s *Service) GetAllTagsByUser(userID int) []storage.Tag {
	return s.r.GetAllTagsByUser(userID)
}
