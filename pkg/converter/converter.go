package converter

type Converter struct {
	IScheduleConv
	IHomeworkConv
}

func NewConverter() *Converter {
	return &Converter{
		IScheduleConv: NewScheduleConv(),
		IHomeworkConv: NewHomeworkConv(),
	}
}
