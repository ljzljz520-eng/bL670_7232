package service

import (
	"fmt"
	"training-review/internal/domain"
	"training-review/internal/store"
)

type WorkflowService struct {
	store *store.Store
	clock Clock
	ids   *IDGenerator
}

func NewWorkflowService(st *store.Store, clock Clock) *WorkflowService {
	return &WorkflowService{store: st, clock: clock, ids: NewIDGenerator("flow")}
}

func (s *WorkflowService) Start(registrationID, name string) (domain.Workflow, error) {
	if _, err := s.store.GetRegistration(registrationID); err != nil {
		return domain.Workflow{}, err
	}
	if name == "" {
		return domain.Workflow{}, fmt.Errorf("workflow name is required")
	}
	flow := domain.Workflow{ID: s.ids.Next(), RegistrationID: registrationID, Name: name, State: "started", UpdatedAt: s.clock.Now()}
	if err := s.store.SaveWorkflow(flow); err != nil {
		return domain.Workflow{}, err
	}
	return flow, nil
}

func (s *WorkflowService) Advance(flow domain.Workflow, state string) (domain.Workflow, error) {
	if state == "" {
		return domain.Workflow{}, fmt.Errorf("state is required")
	}
	if flow.State == "archived" {
		return domain.Workflow{}, fmt.Errorf("workflow archived")
	}
	flow.State, flow.UpdatedAt = state, s.clock.Now()
	return flow, s.store.SaveWorkflow(flow)
}
