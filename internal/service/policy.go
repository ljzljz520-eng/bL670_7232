package service

import (
	"fmt"
	"training-review/internal/domain"
)

func ValidateBeforePersist(item domain.Registration) error {
	issues := domain.ValidateRegistration(item)
	if len(issues) > 0 {
		return fmt.Errorf("%s", domain.FormatIssues(issues))
	}
	return nil
}

func ReviewEligibility(item domain.Registration, reviewer domain.Reviewer) error {
	if item.Status != domain.StatusPending {
		return fmt.Errorf("registration is not pending")
	}
	if !reviewer.CanReview(item.Department) {
		return fmt.Errorf("reviewer is not eligible")
	}
	return nil
}

func IsStale(item domain.Registration, expectedVersion int) bool {
	return item.Version != expectedVersion
}

func NextVersion(item domain.Registration) domain.Registration { item.Version++; return item }

func BuildStatusDetail(from, to domain.Status) string {
	return domain.StatusLabel(from) + " -> " + domain.StatusLabel(to)
}
