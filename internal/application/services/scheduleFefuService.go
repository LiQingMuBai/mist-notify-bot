package services

import (
	"fmt"
	"homework_bot/internal/application"
	"homework_bot/internal/domain"
	"homework_bot/pkg/scheduleParser"
	"time"
)

type ScheduleFefuService struct {
	parser scheduleParser.IFefuParser
	sorter application.Sorter
}

func NewScheduleFefuService() *ScheduleFefuService {
	return &ScheduleFefuService{
		parser: scheduleParser.NewParser(),
		sorter: application.NewSorter(),
	}
}

func generateLink(numberGroup, typeSchedule string, firstDate, secondDate time.Time) string {
	groups := map[string]string{
		"Б9122-09.03.04": "5560",
		"Б3122-08.03.01": "5623",
	}

	link := fmt.Sprintf("https://univer.dvfu.ru/schedule/get?type=%s&start=%s", typeSchedule, firstDate.Format("2006-01-02"))
	link += "T14%3A00%3A00.000Z&" + fmt.Sprintf("end=%s", secondDate.Format("2006-01-02")) + "T14%3A00%3A00.000Z&"
	link += "groups%5B%5D=" + fmt.Sprintf("%s&ppsId=&facilityId=0", groups[numberGroup])
	return link
}

func (s *ScheduleFefuService) GetOnDate(user domain.User, date time.Time) domain.Schedule {
	//link := generateLink(user.Associates, "agendaDay", date.Add(-24*time.Hour), date)
	//schedule, err := s.parser.ParseSchedule(link, user.Associates)
	//if err != nil {
	//	return domain.Schedule{}
	//}
	//
	//s.sorter.SortSchedule(&schedule)
	//return schedule
	return domain.Schedule{}
}

func getDatesForWeek() (time.Time, time.Time) {
	date := time.Now()
	dayOfWeek := int(date.Weekday())
	lastSunday := date.AddDate(0, 0, -dayOfWeek)
	saturday := date.AddDate(0, 0, 6-dayOfWeek)
	return lastSunday, saturday
}

func (s *ScheduleFefuService) GetOnWeek(user domain.User, lastSunday, saturday time.Time) domain.Schedule {
	//link := generateLink(user.Associates, "agendaWeek", lastSunday, saturday)
	//
	//schedule, err := s.parser.ParseSchedule(link, user.Id)
	//if err != nil {
	//	logrus.Errorf("Error in parse schedule, %v", err)
	//	return domain.Schedule{}
	//}
	//
	//s.sorter.SortSchedule(&schedule)
	//return schedule
	//return
	return domain.Schedule{}
}

func (s *ScheduleFefuService) GetOnThisWeek(user domain.User) domain.Schedule {
	lastSunday, saturday := getDatesForWeek()
	schedule := s.GetOnWeek(user, lastSunday, saturday)
	return schedule
}

func (s *ScheduleFefuService) GetOnNextWeek(user domain.User) domain.Schedule {
	lastSunday, saturday := getDatesForWeek()
	lastSunday = lastSunday.AddDate(0, 0, 7)
	saturday = lastSunday.AddDate(0, 0, 6)

	schedule := s.GetOnWeek(user, lastSunday, saturday)
	return schedule
}

func (s *ScheduleFefuService) GetOnToday(user domain.User) domain.Schedule {
	return s.GetOnDate(user, time.Now())
}

func (s *ScheduleFefuService) GetOnTomorrow(user domain.User) domain.Schedule {
	schedule := s.GetOnDate(user, time.Now().AddDate(0, 0, 1))
	s.sorter.SortSchedule(&schedule)
	return schedule
}
