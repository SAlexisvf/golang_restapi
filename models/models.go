package models

type Task struct {
	ID          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	Duedate     string `json:"duedate,omitempty"`
	Class       string `json:"class,omitempty"`
}
