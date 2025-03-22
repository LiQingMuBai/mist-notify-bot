package services

import (
	"homework_bot/internal/application/interfaces"
	"homework_bot/internal/domain"
	"time"
)

type HomeworkService struct {
	homeworkRepo interfaces.IHomeworkRepository
}

func NewHomeworkService(repos interfaces.IHomeworkRepository) *HomeworkService {
	return &HomeworkService{
		homeworkRepo: repos,
	}
}

func (s *HomeworkService) Create(homework domain.Homework) (int, error) {
	homework.CreatedAt = time.Now()
	homework.UpdatedAt = time.Now()
	return s.homeworkRepo.Create(homework)
}

func (s *HomeworkService) GetByTags(tags []string) ([]domain.HomeworkToGet, error) {
	return s.homeworkRepo.GetByTags(tags)
}

func (s *HomeworkService) GetById(id int) (domain.HomeworkToGet, error) {
	return s.homeworkRepo.GetById(id)
}

func (s *HomeworkService) GetAll() ([]domain.HomeworkToGet, error) {
	return s.homeworkRepo.GetAll()
}

func (s *HomeworkService) GetByName(name string) ([]domain.HomeworkToGet, error) {
	return s.homeworkRepo.GetByName(name)
}

func (s *HomeworkService) GetByWeek() ([]domain.HomeworkToGet, error) {
	return s.homeworkRepo.GetByWeek()
}

func (s *HomeworkService) GetByToday() ([]domain.HomeworkToGet, error) {
	return s.homeworkRepo.GetByToday()
}

func (s *HomeworkService) GetByTomorrow() ([]domain.HomeworkToGet, error) {
	return s.homeworkRepo.GetByTomorrow()
}

func (s *HomeworkService) GetByDate(date time.Time) ([]domain.HomeworkToGet, error) {
	return s.homeworkRepo.GetByDate(date)
}

func (s *HomeworkService) Update(homeworkToUpdate domain.HomeworkToUpdate) (domain.Homework, error) {
	return s.homeworkRepo.Update(homeworkToUpdate)
}

func (s *HomeworkService) Delete(id int) error {
	return s.homeworkRepo.Delete(id)
}
