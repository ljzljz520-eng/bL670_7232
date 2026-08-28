package report

import (
	"fmt"
	"strings"
	"training-review/internal/domain"
)

func RenderDepartmentMetrics(metrics []DepartmentMetric) string {
	var out strings.Builder
	out.WriteString("部门,总数,待审核,已通过,已驳回\n")
	for _, metric := range metrics {
		out.WriteString(fmt.Sprintf("%s,%d,%d,%d,%d\n", csvField(metric.Department), metric.Total, metric.Pending, metric.Approved, metric.Rejected))
	}
	return out.String()
}

func RenderStatusCounts(counts map[domain.Status]int) string {
	var out strings.Builder
	for _, status := range []domain.Status{domain.StatusDraft, domain.StatusPending, domain.StatusApproved, domain.StatusRejected, domain.StatusArchived} {
		out.WriteString(fmt.Sprintf("%s=%d\n", status, counts[status]))
	}
	return out.String()
}

func RenderCourseCounts(counts map[string]int) string {
	var out strings.Builder
	for course, count := range counts {
		out.WriteString(fmt.Sprintf("%s=%d\n", course, count))
	}
	return out.String()
}

func RenderRegistration(item domain.Registration) string {
	builder := NewBuilder("培训报名 " + item.ID)
	builder.Add("课程", item.CourseID)
	builder.Add("参与者", item.Participant)
	builder.Add("状态", domain.StatusLabel(item.Status))
	builder.Add("版本", fmt.Sprintf("%d", item.Version))
	builder.Add("资料", strings.Join(item.Tags, ","))
	return builder.Render()
}
