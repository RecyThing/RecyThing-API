package helper

type CountMissionApproval struct {
	TotalCount    int64 `json:"total_count"`
	CountPending  int64 `json:"count_pending"`
	CountApproved int64 `json:"count_approved"`
	CountRejected int64 `json:"count_rejected"`
}

type CountMission struct {
	TotalCount   int64 `json:"total_count"`
	CountActive  int64 `json:"count_active"`
	CountExpired int64 `json:"count_expired"`
}
