package model

// PostVote structure for read json from vote request
type PostVote struct {
	PostID int64 `json:"post_id"`
	UserID int64 `json:"user_id"`
	Vote   int8  `json:"vote"`
}
