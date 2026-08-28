package service

import (
	"fmt"
	"strings"
	"training-review/internal/domain"
)

type ChecklistItem struct {
	Key      string
	Label    string
	Complete bool
}

func RegistrationChecklist(item domain.Registration) []ChecklistItem {
	return []ChecklistItem{
		{Key: "identity", Label: "课程和参与者", Complete: item.CourseID != "" && item.Participant != ""},
		{Key: "contact", Label: "邮箱", Complete: strings.Contains(item.Email, "@")},
		{Key: "department", Label: "所属部门", Complete: item.Department != ""},
		{Key: "tags", Label: "业务标签", Complete: len(item.Tags) > 0},
	}
}

func ChecklistComplete(items []ChecklistItem) bool {
	for _, item := range items {
		if !item.Complete {
			return false
		}
	}
	return true
}

func MissingChecklist(items []ChecklistItem) []string {
	missing := make([]string, 0)
	for _, item := range items {
		if !item.Complete {
			missing = append(missing, item.Label)
		}
	}
	return missing
}

func ChecklistMessage(item domain.Registration) string {
	missing := MissingChecklist(RegistrationChecklist(item))
	if len(missing) == 0 {
		return "报名资料完整"
	}
	return fmt.Sprintf("缺少：%s", strings.Join(missing, "、"))
}

func EligibleForSubmit(item domain.Registration) bool {
	return ChecklistComplete(RegistrationChecklist(item)) && item.Status == domain.StatusDraft
}

func EligibleForArchive(item domain.Registration) bool { return item.Status == domain.StatusApproved }
