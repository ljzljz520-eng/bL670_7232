package domain

func CanTransition(from, to Status) bool {
	if from == StatusDraft && to == StatusPending {
		return true
	}
	if from == StatusPending && (to == StatusApproved || to == StatusRejected) {
		return true
	}
	if (from == StatusApproved || from == StatusRejected) && to == StatusArchived {
		return true
	}
	if from == StatusArchived && to == StatusApproved {
		return true
	}
	return false
}

func IsTerminal(status Status) bool {
	return status == StatusArchived || status == StatusRejected
}

func NormalizeStatus(value string) Status {
	for _, status := range []Status{StatusDraft, StatusPending, StatusApproved, StatusRejected, StatusArchived} {
		if string(status) == value {
			return status
		}
	}
	return StatusDraft
}

func StatusLabel(status Status) string {
	switch status {
	case StatusDraft:
		return "草稿"
	case StatusPending:
		return "待审核"
	case StatusApproved:
		return "已通过"
	case StatusRejected:
		return "已驳回"
	case StatusArchived:
		return "已归档"
	default:
		return "未知"
	}
}
