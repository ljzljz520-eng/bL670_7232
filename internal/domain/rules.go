package domain

import (
	"fmt"
	"strings"
)

type ValidationIssue struct {
	Field   string
	Message string
}

func (i ValidationIssue) Error() string { return i.Field + ": " + i.Message }

func ValidateRegistration(item Registration) []ValidationIssue {
	issues := make([]ValidationIssue, 0)
	if strings.TrimSpace(item.ID) == "" {
		issues = append(issues, ValidationIssue{"id", "required"})
	}
	if strings.TrimSpace(item.CourseID) == "" {
		issues = append(issues, ValidationIssue{"course_id", "required"})
	}
	if strings.TrimSpace(item.Participant) == "" {
		issues = append(issues, ValidationIssue{"participant", "required"})
	}
	if !strings.Contains(item.Email, "@") {
		issues = append(issues, ValidationIssue{"email", "invalid"})
	}
	if item.Version < 1 {
		issues = append(issues, ValidationIssue{"version", "must be positive"})
	}
	if item.Status == "" {
		issues = append(issues, ValidationIssue{"status", "required"})
	}
	return issues
}

func ValidateReview(review Review) []ValidationIssue {
	issues := make([]ValidationIssue, 0)
	if review.RegistrationID == "" {
		issues = append(issues, ValidationIssue{"registration_id", "required"})
	}
	if review.Reviewer == "" {
		issues = append(issues, ValidationIssue{"reviewer", "required"})
	}
	if review.Decision != StatusApproved && review.Decision != StatusRejected {
		issues = append(issues, ValidationIssue{"decision", "must approve or reject"})
	}
	return issues
}

func FormatIssues(issues []ValidationIssue) string {
	if len(issues) == 0 {
		return ""
	}
	parts := make([]string, 0, len(issues))
	for _, issue := range issues {
		parts = append(parts, issue.Error())
	}
	return fmt.Sprintf("validation failed: %s", strings.Join(parts, "; "))
}

func AllowedTargets(from Status) []Status {
	targets := make([]Status, 0, 2)
	for _, candidate := range []Status{StatusPending, StatusApproved, StatusRejected, StatusArchived} {
		if CanTransition(from, candidate) {
			targets = append(targets, candidate)
		}
	}
	return targets
}

func StatusRank(status Status) int {
	switch status {
	case StatusDraft:
		return 1
	case StatusPending:
		return 2
	case StatusApproved:
		return 3
	case StatusRejected:
		return 4
	case StatusArchived:
		return 5
	default:
		return 0
	}
}
