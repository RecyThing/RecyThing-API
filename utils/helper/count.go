package helper

type CountMissionApproval struct {
	All         int64 `json:"all"`
	NeedReview  int64 `json:"need_review"`
	NotApproved int64 `json:"not_approved"`
	Approved    int64 `json:"approved"`
	Search      int64 `json:"search,omitempty"`
}