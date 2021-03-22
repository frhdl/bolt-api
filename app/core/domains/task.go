package domains

import "time"

// Task represent a task in the application.
type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	UserID      int       `json:"user_id"`
	ProjectID   int       `json:"project_id"`
	CreateAt    time.Time `json:"create_at"`
	FinishAt    time.Time `json:"finish_at"`
	Done        bool      `json:"done"`
}
