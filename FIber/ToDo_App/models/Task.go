package tasks

import "time"

type Task struct {
	ID          int        `json:"id" gorm:"primaryKey"`
	Description string     `json:"description"`
	Completed   bool       `json:"completed"`
	CreatedOn   time.Time  `json:"created_on"`
	FinishedOn  *time.Time `json:"finished_on,omitempty" gorm:"default:null"`
}
