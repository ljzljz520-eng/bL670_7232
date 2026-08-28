package report

import (
	"fmt"
	"strings"
	"training-review/internal/domain"
)

func AuditText(events []domain.AuditEvent) string {
	var out strings.Builder
	for _, event := range events {
		out.WriteString(fmt.Sprintf("%s %s %s %s\n", event.At, event.Action, event.Actor, event.Detail))
	}
	return out.String()
}

func ReviewText(reviews []domain.Review) string {
	var out strings.Builder
	for _, review := range reviews {
		out.WriteString(fmt.Sprintf("%s %s %s %s\n", review.ReviewedAt, review.Reviewer, domain.StatusLabel(review.Decision), review.Note))
	}
	return out.String()
}

func WorkflowText(workflows []domain.Workflow) string {
	var out strings.Builder
	for _, workflow := range workflows {
		out.WriteString(fmt.Sprintf("%s %s %s\n", workflow.UpdatedAt, workflow.Name, workflow.State))
	}
	return out.String()
}

func ArchiveText(records []domain.Record) string {
	var out strings.Builder
	for _, record := range records {
		out.WriteString(fmt.Sprintf("%s %s %s\n", record.CreatedAt, record.Kind, record.Payload))
	}
	return out.String()
}

func AttachmentText(items []domain.Attachment) string {
	var out strings.Builder
	for _, item := range items {
		out.WriteString(fmt.Sprintf("%s %s %d %s\n", item.Name, item.ContentType, item.Size, item.Digest))
	}
	return out.String()
}

func Timeline(events []domain.AuditEvent, reviews []domain.Review) []string {
	result := make([]string, 0, len(events)+len(reviews))
	for _, event := range events {
		result = append(result, event.At+" "+event.Action)
	}
	for _, review := range reviews {
		result = append(result, review.ReviewedAt+" "+string(review.Decision))
	}
	return result
}
