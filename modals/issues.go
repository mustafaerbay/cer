package modals

// GET /issues
// GET /issues?assignee_id=5
// GET /issues?author_id=5
// GET /issues?confidential=true
// GET /issues?iids[]=42&iids[]=43
// GET /issues?labels=foo
// GET /issues?labels=foo,bar
// GET /issues?labels=foo,bar&state=opened
// GET /issues?milestone=1.0.0
// GET /issues?milestone=1.0.0&state=opened
// GET /issues?my_reaction_emoji=star
// GET /issues?search=foo&in=title
// GET /issues?state=closed
// GET /issues?state=opened

import (
	"time"
)

type ISOTime time.Time

// IssueAuthor represents a author of the issue.
type IssueAuthor struct {
	ID        int    `json:"id"`
	State     string `json:"state"`
	WebURL    string `json:"web_url"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
	Username  string `json:"username"`
}

// IssueAssignee represents a assignee of the issue.
type IssueAssignee struct {
	ID        int    `json:"id"`
	State     string `json:"state"`
	WebURL    string `json:"web_url"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
	Username  string `json:"username"`
}

// IssueReferences represents references of the issue.
type IssueReferences struct {
	Short    string `json:"short"`
	Relative string `json:"relative"`
	Full     string `json:"full"`
}

// IssueCloser represents a closer of the issue.
type IssueCloser struct {
	ID        int    `json:"id"`
	State     string `json:"state"`
	WebURL    string `json:"web_url"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
	Username  string `json:"username"`
}

// IssueLinks represents links of the issue.
type IssueLinks struct {
	Self       string `json:"self"`
	Notes      string `json:"notes"`
	AwardEmoji string `json:"award_emoji"`
	Project    string `json:"project"`
}

// LabelDetails represents detailed label information.
type LabelDetails struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Color           string `json:"color"`
	Description     string `json:"description"`
	DescriptionHTML string `json:"description_html"`
	TextColor       string `json:"text_color"`
}

type Labels []string

// Issue represents a GitLab issue.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/issues.html
type Issue struct {
	// ID                   int                    `json:"id"`
	IID int `json:"iid"`
	// ProjectID            int                    `json:"project_id"`
	Title    string         `json:"title"`
	Assignee *IssueAssignee `json:"assignee"`
	Labels   Labels         `json:"labels"`
	DueDate  *ISOTime       `json:"due_date"`
	// ExternalID           string                 `json:"external_id"`
	State string `json:"state"`
	// Description          string                 `json:"description"`
	Author *IssueAuthor `json:"author"`
	// Milestone            *Milestone             `json:"milestone"`
	// Assignees            []*IssueAssignee       `json:"assignees"`
	UpdatedAt time.Time `json:"updated_at"`
	// ClosedAt             *time.Time             `json:"closed_at"`
	// ClosedBy             *IssueCloser           `json:"closed_by"`
	CreatedAt time.Time `json:"created_at"`
	// MovedToID            int                    `json:"moved_to_id"`
	// // LabelDetails         []*LabelDetails        `json:"label_details"`
	// Upvotes              int                    `json:"upvotes"`
	// Downvotes            int                    `json:"downvotes"`
	WebURL string `json:"web_url"`
	// References           *IssueReferences       `json:"references"`
	// // TimeStats            *TimeStats       		`json:"time_stats"`
	// Confidential         bool                   `json:"confidential"`
	// Weight               int                    `json:"weight"`
	// DiscussionLocked     bool                   `json:"discussion_locked"`
	// IssueType            *string                `json:"issue_type,omitempty"`
	// Subscribed           bool                   `json:"subscribed"`
	// UserNotesCount       int                    `json:"user_notes_count"`
	// Links                *IssueLinks            `json:"_links"`
	// IssueLinkID          int                    `json:"issue_link_id"`
	// MergeRequestCount    int                    `json:"merge_requests_count"`
	// EpicIssueID          int                    `json:"epic_issue_id"`
	// // Epic                 *Epic                  `json:"epic"`
	// TaskCompletionStatus *TasksCompletionStatus `json:"task_completion_status"`
}

type IssueList struct {
	ManyIssues []*Issue `json:"data"`
}
