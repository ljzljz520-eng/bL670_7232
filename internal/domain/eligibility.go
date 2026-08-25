package domain

type Eligibility struct {
	CourseID     string
	Department   string
	RequiredTags []string
	MaxVersion   int
}

func (e Eligibility) Check(item Registration) []string {
	issues := make([]string, 0)
	if e.CourseID != "" && item.CourseID != e.CourseID {
		issues = append(issues, "course mismatch")
	}
	if e.Department != "" && item.Department != e.Department {
		issues = append(issues, "department mismatch")
	}
	if e.MaxVersion > 0 && item.Version > e.MaxVersion {
		issues = append(issues, "version exceeded")
	}
	for _, required := range e.RequiredTags {
		if !hasTag(item.Tags, required) {
			issues = append(issues, "missing tag: "+required)
		}
	}
	return issues
}

func hasTag(tags []string, expected string) bool {
	for _, tag := range tags {
		if tag == expected {
			return true
		}
	}
	return false
}

func (e Eligibility) Accepts(item Registration) bool { return len(e.Check(item)) == 0 }

func (e Eligibility) Description() string {
	result := ""
	if e.CourseID != "" {
		result += "course=" + e.CourseID
	}
	if e.Department != "" {
		result += " department=" + e.Department
	}
	return result
}

func (e Eligibility) WithTag(tag string) Eligibility {
	e.RequiredTags = append(append([]string(nil), e.RequiredTags...), tag)
	return e
}

func (e Eligibility) WithCourse(courseID string) Eligibility { e.CourseID = courseID; return e }

func (e Eligibility) WithDepartment(department string) Eligibility {
	e.Department = department
	return e
}
