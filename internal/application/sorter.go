package application

import (
	"homework_bot/internal/domain"
	"sort"
)

type Sorter struct{}

func NewSorter() Sorter {
	return Sorter{}
}

func (s *Sorter) SortSchedule(schedule *domain.Schedule) {
	sort.Slice(schedule.Subjects, func(i, j int) bool {
		return schedule.Subjects[i].Start.Before(schedule.Subjects[j].Start)
	})
}
