package taskgoal

import "wetodo/internal/storage"

type Repository interface {
	GetTaskGoalByID(userID int, id string) (storage.TaskGoal, bool)
}

type Service struct {
	r Repository
}

func NewService(r Repository) *Service {
	return &Service{r}
}

func (s *Service) GetTaskGoalByID(userID int, taskGoalID string) (storage.TaskGoal, bool) {
	return s.r.GetTaskGoalByID(userID, taskGoalID)
}
