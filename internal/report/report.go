package report

import (
	"fmt"
	"sort"
	"strings"

	"training-review/internal/domain"
)

type Builder struct {
	title string
	lines []string
}

func NewBuilder(title string) *Builder { return &Builder{title: title, lines: make([]string, 0)} }

func (b *Builder) Add(label, value string) { b.lines = append(b.lines, label+": "+value) }

func (b *Builder) Render() string {
	var out strings.Builder
	if b.title != "" {
		out.WriteString(b.title)
		out.WriteByte('\n')
	}
	for _, line := range b.lines {
		out.WriteString(line)
		out.WriteByte('\n')
	}
	return out.String()
}

func SummaryText(summary domain.EnrollmentSummary) string {
	b := NewBuilder("培训报名汇总")
	b.Add("总数", fmt.Sprintf("%d", summary.Total))
	statuses := make([]string, 0, len(summary.ByStatus))
	for status := range summary.ByStatus {
		statuses = append(statuses, string(status))
	}
	sort.Strings(statuses)
	for _, status := range statuses {
		b.Add(domain.StatusLabel(domain.Status(status)), fmt.Sprintf("%d", summary.ByStatus[domain.Status(status)]))
	}
	courses := make([]string, 0, len(summary.Courses))
	for course := range summary.Courses {
		courses = append(courses, course)
	}
	sort.Strings(courses)
	for _, course := range courses {
		b.Add("课程 "+course, fmt.Sprintf("%d", summary.Courses[course]))
	}
	return b.Render()
}

func CSV(items []domain.Registration) string {
	result := strings.Builder{}
	result.WriteString("id,course,participant,email,department,status,version,tags\n")
	for _, item := range items {
		result.WriteString(csvField(item.ID))
		result.WriteByte(',')
		result.WriteString(csvField(item.CourseID))
		result.WriteByte(',')
		result.WriteString(csvField(item.Participant))
		result.WriteByte(',')
		result.WriteString(csvField(item.Email))
		result.WriteByte(',')
		result.WriteString(csvField(item.Department))
		result.WriteByte(',')
		result.WriteString(csvField(string(item.Status)))
		result.WriteByte(',')
		result.WriteString(fmt.Sprintf("%d,", item.Version))
		result.WriteString(csvField(strings.Join(item.Tags, "|")))
		result.WriteByte('\n')
	}
	return result.String()
}

func csvField(value string) string {
	if strings.ContainsAny(value, ",\"\n") {
		return "\"" + strings.ReplaceAll(value, "\"", "\"\"") + "\""
	}
	return value
}

func GroupByDepartment(items []domain.Registration) map[string][]domain.Registration {
	groups := make(map[string][]domain.Registration)
	for _, item := range items {
		groups[item.Department] = append(groups[item.Department], item)
	}
	for key := range groups {
		sort.Slice(groups[key], func(i, j int) bool { return groups[key][i].ID < groups[key][j].ID })
	}
	return groups
}
