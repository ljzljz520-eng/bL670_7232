package service

import (
	"fmt"
	"strings"

	"training-review/internal/domain"
)

type ImportRow struct {
	CourseID, Participant, Email, Department string
	Tags                                     []string
}
type ImportResult struct {
	Accepted []domain.Registration
	Rejected []string
}

func (s *RegistrationService) Import(rows []ImportRow) (ImportResult, error) {
	result := ImportResult{Accepted: make([]domain.Registration, 0), Rejected: make([]string, 0)}
	for index, row := range rows {
		if strings.TrimSpace(row.CourseID) == "" || strings.TrimSpace(row.Participant) == "" {
			result.Rejected = append(result.Rejected, fmt.Sprintf("row %d: missing identity", index+1))
			continue
		}
		item, err := s.Register(row.CourseID, row.Participant, row.Email, row.Department, row.Tags)
		if err != nil {
			result.Rejected = append(result.Rejected, fmt.Sprintf("row %d: %v", index+1, err))
			continue
		}
		result.Accepted = append(result.Accepted, item)
	}
	return result, nil
}

func (s *RegistrationService) Export(filter domain.SearchFilter) (string, error) {
	items, err := s.Search(filter)
	if err != nil {
		return "", err
	}
	var builder strings.Builder
	builder.WriteString("id,course,participant,status,version\n")
	for _, item := range items {
		builder.WriteString(fmt.Sprintf("%s,%s,%s,%s,%d\n", item.ID, item.CourseID, item.Participant, item.Status, item.Version))
	}
	return builder.String(), nil
}
