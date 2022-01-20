package modals

// TasksCompletionStatus represents tasks of the issue/merge request.
type TasksCompletionStatus struct {
	Count          int `json:"count"`
	CompletedCount int `json:"completed_count"`
}
