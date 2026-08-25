package domain

type Status string

func (s Status) String() string { return string(s) }

const (
	StatusDraft    Status = "draft"
	StatusPending  Status = "pending"
	StatusApproved Status = "approved"
	StatusRejected Status = "rejected"
	StatusArchived Status = "archived"
)

type Registration struct {
	ID          string   `json:"id"`
	CourseID    string   `json:"course_id"`
	Participant string   `json:"participant"`
	Email       string   `json:"email"`
	Department  string   `json:"department"`
	SubmittedAt string   `json:"submitted_at"`
	Status      Status   `json:"status"`
	Version     int      `json:"version"`
	Tags        []string `json:"tags"`
}

type Review struct {
	ID             string `json:"id"`
	RegistrationID string `json:"registration_id"`
	Reviewer       string `json:"reviewer"`
	Decision       Status `json:"decision"`
	Note           string `json:"note"`
	ReviewedAt     string `json:"reviewed_at"`
}

type Record struct {
	ID             string `json:"id"`
	RegistrationID string `json:"registration_id"`
	Kind           string `json:"kind"`
	Payload        string `json:"payload"`
	CreatedAt      string `json:"created_at"`
}

type AuditEvent struct {
	ID             string `json:"id"`
	RegistrationID string `json:"registration_id"`
	Action         string `json:"action"`
	Actor          string `json:"actor"`
	Detail         string `json:"detail"`
	At             string `json:"at"`
}

type Workflow struct {
	ID             string `json:"id"`
	RegistrationID string `json:"registration_id"`
	Name           string `json:"name"`
	State          string `json:"state"`
	UpdatedAt      string `json:"updated_at"`
}

type Attachment struct {
	ID             string `json:"id"`
	RegistrationID string `json:"registration_id"`
	Name           string `json:"name"`
	ContentType    string `json:"content_type"`
	Size           int    `json:"size"`
	Digest         string `json:"digest"`
}

type SearchFilter struct {
	CourseID    string
	Participant string
	Department  string
	Status      Status
	Tag         string
}

type EnrollmentSummary struct {
	Total    int            `json:"total"`
	ByStatus map[Status]int `json:"by_status"`
	Courses  map[string]int `json:"courses"`
}
