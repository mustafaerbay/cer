package modals

// Label represents a GitLab label.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/labels.html
type Label struct {
	ID                     int    `json:"id"`
	Name                   string `json:"name"`
	Color                  string `json:"color"`
	TextColor              string `json:"text_color"`
	Description            string `json:"description"`
	OpenIssuesCount        int    `json:"open_issues_count"`
	ClosedIssuesCount      int    `json:"closed_issues_count"`
	OpenMergeRequestsCount int    `json:"open_merge_requests_count"`
	Subscribed             bool   `json:"subscribed"`
	Priority               int    `json:"priority"`
	IsProjectLabel         bool   `json:"is_project_label"`
}
