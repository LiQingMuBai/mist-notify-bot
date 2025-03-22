package converter

import "homework_bot/internal/domain"

type HomeworkConv struct {
}

func NewHomeworkConv() *HomeworkConv {
	return &HomeworkConv{}
}

func (h *HomeworkConv) HomeworkToText(homework domain.HomeworkToGet) string {
	text := "Название: " + homework.Name +
		"\n" + "Описание: " + homework.Description + "\n" +
		"Дедлайн: " + homework.Deadline.String() + "\n"
	for _, tag := range homework.Tags {
		text += "#" + tag + "\n"
	}
	return text
}
