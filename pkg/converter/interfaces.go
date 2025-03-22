package converter

import "homework_bot/internal/domain"

type IScheduleConv interface {
	subjectToText(subject domain.Subject) string
	ScheduleToText(schedule domain.Schedule) map[string]string
}

type IHomeworkConv interface {
	HomeworkToText(homework domain.HomeworkToGet) string
}
