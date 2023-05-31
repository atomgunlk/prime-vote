package model

type VoteRequest struct {
	ID     uint64 `json:"id" validate:"required"`
	UserId uint64
}
type VoteResponse struct {
	ResponseStatus ResponseStatus `json:"responseStatus"`
}
