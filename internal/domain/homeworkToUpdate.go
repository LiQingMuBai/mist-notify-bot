package domain

import "time"

type HomeworkToUpdate struct {
	Id          int        `db:"id"`
	Name        *string    `db:"name"`
	Description *string    `db:"description"`
	Images      *[]string  `db:"image"`
	Tags        *[]string  `db:"tags"`
	Deadline    *time.Time `db:"deadline"`
}
