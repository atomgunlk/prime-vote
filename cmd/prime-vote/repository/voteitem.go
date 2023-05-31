package repository

import (
	"errors"
	"time"

	"github.com/atomgunlk/prime-vote/cmd/prime-vote/be_error"
	"github.com/atomgunlk/prime-vote/cmd/prime-vote/model"
	"github.com/jackc/pgtype"
)

func (r *voteRepository) CreateVoteItem(req *model.CreateVoteItemRequest) (*model.CreateVoteItemResponse, error) {
	response := new(model.CreateVoteItemResponse)

	ctx, cancel := r.defaultContext()
	defer cancel()

	bkk, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		response.ResponseStatus = model.ResponseStatus{
			Code:    be_error.RepositoryCommon.Code(),
			Message: be_error.RepositoryCommon.Message(),
		}
		return response, errors.Join(err, errors.New("[Repository.CreateVoteItem]"))
	}

	_, err = r.DB.Exec(ctx,
		`INSERT INTO vote_items (created_at, updated_at, name, description, vote_count) 
		VALUES ($1, $2, $3, $4, 0)`,
		time.Now().In(bkk), time.Now().In(bkk), req.Name, req.Description)
	if err != nil {
		response.ResponseStatus = model.ResponseStatus{
			Code:    be_error.RepositoryCannotCreate.Code(),
			Message: be_error.RepositoryCannotCreate.Message(),
		}
		return response, errors.Join(err, errors.New("[Repository.CreateVoteItem]"))
	}

	response.ResponseStatus = model.ResponseStatus{
		Code:    be_error.Success.Code(),
		Message: be_error.Success.Message(),
	}
	return response, nil
}

func (r *voteRepository) ListVoteItem(req *model.ListVoteItemRequest) (*model.ListVoteItemResponse, error) {
	response := new(model.ListVoteItemResponse)

	ctx, cancel := r.defaultContext()
	defer cancel()

	rows, err := r.DB.Query(ctx,
		`SELECT * FROM vote_items WHERE deleted_at IS null ORDER BY id OFFSET $1 LIMIT $2`,
		(req.Page-1)*(req.Size), req.Size,
	)
	if err != nil {
		response.ResponseStatus = model.ResponseStatus{
			Code:    be_error.RepositoryCannotGet.Code(),
			Message: be_error.RepositoryCannotGet.Message(),
		}
		return response, errors.Join(err, errors.New("[Repository.ListVoteItem]"))
	}
	defer rows.Close()

	// Iterate through the result set
	for rows.Next() {
		item := new(model.VoteItem)
		createdAt := pgtype.Timestamptz{}
		updatedAt := pgtype.Timestamptz{}
		deletedAt := pgtype.Timestamptz{}
		err = rows.Scan(&item.ID, &createdAt, &updatedAt, &deletedAt, &item.Name, &item.Description, &item.VoteCount)
		if err != nil {
			response.ResponseStatus = model.ResponseStatus{
				Code:    be_error.RepositoryCannotGet.Code(),
				Message: be_error.RepositoryCannotGet.Message(),
			}
			return response, errors.Join(err, errors.New("[Repository.ListVoteItem]"))
		}
		item.CreatedAt = createdAt.Time
		item.UpdatedAt = updatedAt.Time
		if deletedAt.Status != pgtype.Null {
			item.DeletedAt = &deletedAt.Time
		}

		response.Items = append(response.Items, *item)
	}

	response.ResponseStatus = model.ResponseStatus{
		Code:    be_error.Success.Code(),
		Message: be_error.Success.Message(),
	}
	return response, nil
}

func (r *voteRepository) GetVoteItem(req *model.GetVoteItemRequest) (*model.GetVoteItemResponse, error) {
	response := new(model.GetVoteItemResponse)

	ctx, cancel := r.defaultContext()
	defer cancel()

	row := r.DB.QueryRow(ctx, `SELECT * FROM vote_items WHERE id = $1 AND deleted_at IS null`, req.ID)
	item := new(model.VoteItem)
	createdAt := pgtype.Timestamptz{}
	updatedAt := pgtype.Timestamptz{}
	deletedAt := pgtype.Timestamptz{}
	err := row.Scan(&item.ID, &createdAt, &updatedAt, &deletedAt, &item.Name, &item.Description, &item.VoteCount)
	if err != nil {
		response.ResponseStatus = model.ResponseStatus{
			Code:    be_error.RepositoryCannotGet.Code(),
			Message: be_error.RepositoryCannotGet.Message(),
		}
		return response, errors.Join(err, errors.New("[Repository.ListVoteItem]"))
	}

	item.CreatedAt = createdAt.Time
	item.UpdatedAt = updatedAt.Time
	if deletedAt.Status != pgtype.Null {
		item.DeletedAt = &deletedAt.Time
	}

	response.Item = *item

	response.ResponseStatus = model.ResponseStatus{
		Code:    be_error.Success.Code(),
		Message: be_error.Success.Message(),
	}
	return response, nil
}

func (r *voteRepository) UpdateVoteItem(req *model.UpdateVoteItemRequest) (*model.UpdateVoteItemResponse, error) {
	response := new(model.UpdateVoteItemResponse)

	ctx, cancel := r.defaultContext()
	defer cancel()

	bkk, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		response.ResponseStatus = model.ResponseStatus{
			Code:    be_error.RepositoryCommon.Code(),
			Message: be_error.RepositoryCommon.Message(),
		}
		return response, errors.Join(err, errors.New("[Repository.UpdateVoteItem]"))
	}

	result, err := r.DB.Exec(ctx,
		`UPDATE vote_items SET updated_at = $1, name = $2, description = $3
		WHERE id = $4 AND vote_count = 0 AND deleted_at IS null `,
		time.Now().In(bkk), req.Name, req.Description, req.ID)
	if err != nil {
		response.ResponseStatus = model.ResponseStatus{
			Code:    be_error.RepositoryCannotUpdate.Code(),
			Message: be_error.RepositoryCannotUpdate.Message(),
		}
		return response, errors.Join(err, errors.New("[Repository.UpdateVoteItem]"))
	}

	if result.RowsAffected() <= 0 {
		response.ResponseStatus = model.ResponseStatus{
			Code:    be_error.RepositoryCannotUpdate.Code(),
			Message: be_error.RepositoryCannotUpdate.Message(),
		}
		return response, errors.New("[Repository.UpdateVoteItem]: no row update")
	}

	response.ResponseStatus = model.ResponseStatus{
		Code:    be_error.Success.Code(),
		Message: be_error.Success.Message(),
	}
	return response, nil
}

func (r *voteRepository) ClearVoteItem(req *model.ClearVoteItemRequest) (*model.ClearVoteItemResponse, error) {
	response := new(model.ClearVoteItemResponse)

	ctx, cancel := r.defaultContext()
	defer cancel()

	bkk, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		response.ResponseStatus = model.ResponseStatus{
			Code:    be_error.RepositoryCommon.Code(),
			Message: be_error.RepositoryCommon.Message(),
		}
		return response, errors.Join(err, errors.New("[Repository.ClearVoteItem]"))
	}

	_, err = r.DB.Exec(ctx,
		`UPDATE vote_items SET updated_at = $1, vote_count = 0
		WHERE id = $2 AND deleted_at IS null `,
		time.Now().In(bkk), req.ID)
	if err != nil {
		response.ResponseStatus = model.ResponseStatus{
			Code:    be_error.RepositoryCannotUpdate.Code(),
			Message: be_error.RepositoryCannotUpdate.Message(),
		}
		return response, errors.Join(err, errors.New("[Repository.ClearVoteItem]"))
	}

	response.ResponseStatus = model.ResponseStatus{
		Code:    be_error.Success.Code(),
		Message: be_error.Success.Message(),
	}
	return response, nil
}

func (r *voteRepository) DeleteVoteItem(req *model.DeleteVoteItemRequest) (*model.DeleteVoteItemResponse, error) {
	response := new(model.DeleteVoteItemResponse)

	ctx, cancel := r.defaultContext()
	defer cancel()

	bkk, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		response.ResponseStatus = model.ResponseStatus{
			Code:    be_error.RepositoryCommon.Code(),
			Message: be_error.RepositoryCommon.Message(),
		}
		return response, errors.Join(err, errors.New("[Repository.DeleteVoteItem]"))
	}

	_, err = r.DB.Exec(ctx, `UPDATE vote_items SET deleted_at = $1 WHERE id = $2 AND deleted_at IS null `, time.Now().In(bkk), req.ID)
	if err != nil {
		response.ResponseStatus = model.ResponseStatus{
			Code:    be_error.RepositoryCannotDelete.Code(),
			Message: be_error.RepositoryCannotDelete.Message(),
		}
		return response, errors.Join(err, errors.New("[Repository.DeleteVoteItem]"))
	}

	response.ResponseStatus = model.ResponseStatus{
		Code:    be_error.Success.Code(),
		Message: be_error.Success.Message(),
	}
	return response, nil
}

func (r *voteRepository) GetVoteResult(req *model.GetVoteResultRequest) (*model.GetVoteResultResponse, error) {
	response := new(model.GetVoteResultResponse)

	ctx, cancel := r.defaultContext()
	defer cancel()

	rows, err := r.DB.Query(ctx,
		`SELECT * FROM vote_items WHERE deleted_at IS null ORDER BY vote_count DESC , id ASC`,
	)
	if err != nil {
		response.ResponseStatus = model.ResponseStatus{
			Code:    be_error.RepositoryCannotGet.Code(),
			Message: be_error.RepositoryCannotGet.Message(),
		}
		return response, errors.Join(err, errors.New("[Repository.GetVoteResult]"))
	}
	defer rows.Close()

	// Iterate through the result set
	for rows.Next() {
		item := new(model.VoteItem)
		createdAt := pgtype.Timestamptz{}
		updatedAt := pgtype.Timestamptz{}
		deletedAt := pgtype.Timestamptz{}
		err = rows.Scan(&item.ID, &createdAt, &updatedAt, &deletedAt, &item.Name, &item.Description, &item.VoteCount)
		if err != nil {
			response.ResponseStatus = model.ResponseStatus{
				Code:    be_error.RepositoryCannotGet.Code(),
				Message: be_error.RepositoryCannotGet.Message(),
			}
			return response, errors.Join(err, errors.New("[Repository.GetVoteResult]"))
		}
		item.CreatedAt = createdAt.Time
		item.UpdatedAt = updatedAt.Time
		if deletedAt.Status != pgtype.Null {
			item.DeletedAt = &deletedAt.Time
		}

		response.Items = append(response.Items, *item)
	}

	response.ResponseStatus = model.ResponseStatus{
		Code:    be_error.Success.Code(),
		Message: be_error.Success.Message(),
	}
	return response, nil
}
