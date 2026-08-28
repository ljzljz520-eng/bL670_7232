package domain

type WorkflowStep struct {
	ID            string
	Label         string
	RequiredState string
	NextState     string
}

type WorkflowDefinition struct {
	ID    string
	Name  string
	Steps []WorkflowStep
}

func DefaultDefinitions() []WorkflowDefinition {
	return []WorkflowDefinition{
		{ID: "registration-review", Name: "登记审核", Steps: []WorkflowStep{{"create", "创建报名", "", "draft"}, {"submit", "提交审核", "draft", "pending"}, {"review", "完成审核", "pending", "approved"}, {"archive", "归档", "approved", "archived"}}},
		{ID: "registration-change", Name: "查询变更", Steps: []WorkflowStep{{"search", "查询报名", "", "found"}, {"select", "选择记录", "found", "selected"}, {"update", "变更资料", "selected", "updated"}, {"publish", "发布状态", "updated", "published"}}},
		{ID: "registration-import", Name: "导入报告", Steps: []WorkflowStep{{"import", "读取文件", "", "loaded"}, {"validate", "校验行", "loaded", "validated"}, {"persist", "持久化", "validated", "persisted"}, {"report", "输出报告", "persisted", "reported"}}},
	}
}

func FindDefinition(id string) (WorkflowDefinition, bool) {
	for _, definition := range DefaultDefinitions() {
		if definition.ID == id {
			return definition, true
		}
	}
	return WorkflowDefinition{}, false
}

func StepFor(definition WorkflowDefinition, stepID string) (WorkflowStep, bool) {
	for _, step := range definition.Steps {
		if step.ID == stepID {
			return step, true
		}
	}
	return WorkflowStep{}, false
}

func CanRunStep(definition WorkflowDefinition, stepID, state string) bool {
	step, ok := StepFor(definition, stepID)
	if !ok {
		return false
	}
	return step.RequiredState == "" || step.RequiredState == state
}

func NextWorkflowState(definition WorkflowDefinition, stepID string) string {
	step, ok := StepFor(definition, stepID)
	if !ok {
		return ""
	}
	return step.NextState
}

func WorkflowProgress(definition WorkflowDefinition, state string) int {
	for index, step := range definition.Steps {
		if step.NextState == state {
			return index + 1
		}
	}
	return 0
}

func WorkflowComplete(definition WorkflowDefinition, state string) bool {
	if len(definition.Steps) == 0 {
		return false
	}
	return definition.Steps[len(definition.Steps)-1].NextState == state
}
