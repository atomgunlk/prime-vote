package model

import "time"

type VoteItem struct {
	ID          uint64     `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	VoteCount   int64      `json:"vote_count"`
}

type CreateVoteItemRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}
type CreateVoteItemResponse struct {
	ResponseStatus ResponseStatus `json:"responseStatus"`
}

type ListVoteItemRequest struct {
	Page uint `query:"page" validate:"required"`
	Size uint `query:"size" validate:"required"`
}
type ListVoteItemResponse struct {
	ResponseStatus ResponseStatus `json:"responseStatus"`
	Items          []VoteItem     `json:"items"`
}

type GetVoteItemRequest struct {
	ID uint64 `param:"id" validate:"required"`
}
type GetVoteItemResponse struct {
	ResponseStatus ResponseStatus `json:"responseStatus"`
	Item           VoteItem       `json:"item"`
}

type UpdateVoteItemRequest struct {
	ID          uint64 `param:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}
type UpdateVoteItemResponse struct {
	ResponseStatus ResponseStatus `json:"responseStatus"`
}

type ClearVoteItemRequest struct {
	ID uint64 `param:"id" validate:"required"`
}
type ClearVoteItemResponse struct {
	ResponseStatus ResponseStatus `json:"responseStatus"`
}

type DeleteVoteItemRequest struct {
	ID uint64 `param:"id" validate:"required"`
}
type DeleteVoteItemResponse struct {
	ResponseStatus ResponseStatus `json:"responseStatus"`
}

type GetVoteResultRequest struct {
}
type GetVoteResultResponse struct {
	ResponseStatus ResponseStatus `json:"responseStatus"`
	Items          []VoteItem     `json:"items"`
}

type ExportVoteResultRequest struct {
}
type ExportVoteResultResponse struct {
}
