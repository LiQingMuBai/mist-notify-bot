package domain

import (
	"github.com/lib/pq"
	"time"
)

type HomeworkToGet struct {
	Id          int            `db:"id"`
	Name        string         `db:"name"`
	Description string         `db:"description"`
	Images      pq.StringArray `db:"images"`
	Tags        pq.StringArray `db:"tags"`
	CreatedAt   time.Time      `db:"create_at"`
	Deadline    time.Time      `db:"deadline"`
	UpdatedAt   time.Time      `db:"update_at"`
}
