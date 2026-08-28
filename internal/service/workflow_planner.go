package service

import (
	"fmt"
	"training-review/internal/domain"
)

type WorkflowPlanner struct{ definitions []domain.WorkflowDefinition }

func NewWorkflowPlanner() *WorkflowPlanner {
	return &WorkflowPlanner{definitions: domain.DefaultDefinitions()}
}

func (p *WorkflowPlanner) Definition(id string) (domain.WorkflowDefinition, error) {
	for _, definition := range p.definitions {
		if definition.ID == id {
			return definition, nil
		}
	}
	return domain.WorkflowDefinition{}, fmt.Errorf("unknown workflow %s", id)
}

func (p *WorkflowPlanner) ValidateStep(workflowID, stepID, state string) error {
	definition, err := p.Definition(workflowID)
	if err != nil {
		return err
	}
	if !domain.CanRunStep(definition, stepID, state) {
		return fmt.Errorf("step %s cannot run from %s", stepID, state)
	}
	return nil
}

func (p *WorkflowPlanner) Advance(workflowID, stepID string) (string, error) {
	definition, err := p.Definition(workflowID)
	if err != nil {
		return "", err
	}
	next := domain.NextWorkflowState(definition, stepID)
	if next == "" {
		return "", fmt.Errorf("unknown workflow step %s", stepID)
	}
	return next, nil
}

func (p *WorkflowPlanner) Progress(workflowID, state string) (int, bool) {
	definition, err := p.Definition(workflowID)
	if err != nil {
		return 0, false
	}
	return domain.WorkflowProgress(definition, state), domain.WorkflowComplete(definition, state)
}

func (p *WorkflowPlanner) Names() []string {
	result := make([]string, 0, len(p.definitions))
	for _, definition := range p.definitions {
		result = append(result, definition.Name)
	}
	return result
}
