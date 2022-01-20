package modals

import "time"

// Milestone represents a GitLab milestone.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/milestones.html
type Milestone struct {
	ID          int    `json:"id"`
	IID         int    `json:"iid"`
	ProjectID   int    `json:"project_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	// StartDate   *ISOTime   `json:"start_date"`
	// DueDate     *ISOTime   `json:"due_date"`
	State     string     `json:"state"`
	WebURL    string     `json:"web_url"`
	UpdatedAt *time.Time `json:"updated_at"`
	CreatedAt *time.Time `json:"created_at"`
	Expired   *bool      `json:"expired"`
}
