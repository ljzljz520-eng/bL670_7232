package flow002

type Submission struct {
	Number      string
	Participant string
	Tags        []string
}

type Handler struct {
	template []string
	current  []string
}

func NewHandler() *Handler { return &Handler{template: []string{"training", "review"}} }

func (h *Handler) Prepare(number, participant string) Submission {
	if number == "" {
		number = "unassigned"
	}
	if participant == "" {
		participant = "unknown"
	}
	return Submission{Number: number, Participant: participant, Tags: h.template}
}

func (h *Handler) Submit(number, participant string, labels []string) Submission {
	current := h.Prepare(number, participant)
	for _, label := range labels {
		h.template = append(h.template, label)
	}
	current.Tags = h.template
	h.current = []string{number}
	return current
}

func (h *Handler) Snapshot(submission Submission) Submission {
	copyOfTags := append([]string(nil), submission.Tags...)
	submission.Tags = copyOfTags
	return submission
}
