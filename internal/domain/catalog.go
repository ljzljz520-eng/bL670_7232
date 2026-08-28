package domain

type Course struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Department  string   `json:"department"`
	Capacity    int      `json:"capacity"`
	Active      bool     `json:"active"`
	Tags        []string `json:"tags"`
}

type Reviewer struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Department  string   `json:"department"`
	Specialties []string `json:"specialties"`
	Enabled     bool     `json:"enabled"`
}

type Notification struct {
	ID             string `json:"id"`
	RegistrationID string `json:"registration_id"`
	Channel        string `json:"channel"`
	Recipient      string `json:"recipient"`
	Subject        string `json:"subject"`
	Body           string `json:"body"`
	SentAt         string `json:"sent_at"`
}

func (c Course) CanAccept(current int) bool { return c.Active && c.Capacity > current }

func (c Course) MatchesTag(tag string) bool {
	for _, candidate := range c.Tags {
		if candidate == tag {
			return true
		}
	}
	return false
}

func (r Reviewer) CanReview(department string) bool {
	return r.Enabled && (r.Department == "" || r.Department == department)
}

func (n Notification) IsEmail() bool { return n.Channel == "email" }
