package tasklists

type Tasks struct {
	// The text of the task itself
	TaskText string `json:"taskText"`
	// Task List position
	TaskPosition int `json:"taskPosition"`
	// Bool representing if the task complete
	IsCompleted bool `json:"isCompleted"`
}
