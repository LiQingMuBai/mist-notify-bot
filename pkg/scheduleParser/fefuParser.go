package scheduleParser

import (
	"encoding/json"
	"homework_bot/internal/domain"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type FefuParser struct {
}

func NewFefuParser() *FefuParser {
	return &FefuParser{}
}

func ByteToSchedule(body []byte) (domain.Schedule, error) {
	var schedule domain.Schedule
	err := json.Unmarshal(body, &schedule)
	return schedule, err
}

func (p *FefuParser) ParseSchedule(link string, studyGroup int) (domain.Schedule, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return domain.Schedule{}, err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 YaBrowser/24.7.0.0 Safari/537.36")
	req.Header.Add("Referer", "https://univer.dvfu.ru/schedule")
	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Cookie", os.Getenv("COOKIE"))

	resp, err := client.Do(req)
	if err != nil {
		return domain.Schedule{}, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return domain.Schedule{}, err
	}

	schedule, err := ByteToSchedule(body)
	if err != nil {
		return domain.Schedule{}, err
	}

	group := strconv.Itoa(studyGroup)
	var resultInput domain.Schedule
	for _, subject := range schedule.Subjects {
		if subject.Subgroup == group || subject.Subgroup == "" {
			resultInput.Subjects = append(resultInput.Subjects, subject)
		}
	}

	return resultInput, nil
}
