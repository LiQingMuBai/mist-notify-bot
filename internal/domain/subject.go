package domain

import "time"

type Subject struct {
	Title     string    `json:"title"`
	Start     time.Time `json:"start"`
	End       time.Time `json:"end"`
	Group     string    `json:"group"`
	PPSLoad   string    `json:"pps_load"`
	Classroom string    `json:"classroom"`
	Teacher   string    `json:"teacher"`
	Subgroup  string    `json:"subgroup"`
}
