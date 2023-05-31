package repository

import (
	"errors"
	"time"

	"github.com/atomgunlk/prime-vote/cmd/prime-vote/be_error"
	"github.com/atomgunlk/prime-vote/cmd/prime-vote/model"
	"github.com/jackc/pgx/v5"
)

func (r *voteRepository) Vote(req *model.VoteRequest) (*model.VoteResponse, error) {
	response := new(model.VoteResponse)

	ctx, cancel := r.defaultContext()
	defer cancel()

	var err error

	bkk, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		response.ResponseStatus = model.ResponseStatus{
			Code:    be_error.RepositoryCommon.Code(),
			Message: be_error.RepositoryCommon.Message(),
		}
		return response, errors.Join(err, errors.New("[Repository.Vote]"))
	}

	tx, err := r.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		response.ResponseStatus = model.ResponseStatus{
			Code:    be_error.RepositoryCommon.Code(),
			Message: be_error.RepositoryCommon.Message(),
		}
		return response, errors.Join(err, errors.New("[Repository.Vote]"))
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	// Count up
	result, err := tx.Exec(ctx,
		`UPDATE vote_items SET updated_at = $1, vote_count = vote_count + 1
		WHERE id = $2 AND deleted_at IS NULL`,
		time.Now().In(bkk), req.ID)
	if err != nil {
		response.ResponseStatus = model.ResponseStatus{
			Code:    be_error.RepositoryCannotUpdate.Code(),
			Message: be_error.RepositoryCannotUpdate.Message(),
		}
		return response, errors.Join(err, errors.New("[Repository.Vote]"))
	}
	if result.RowsAffected() <= 0 {
		err = errors.New("[Repository.Vote]: no vote_item update")
		response.ResponseStatus = model.ResponseStatus{
			Code:    be_error.RepositoryCannotUpdate.Code(),
			Message: be_error.RepositoryCannotUpdate.Message(),
		}
		return response, err
	}

	// Set user voted
	result, err = tx.Exec(ctx,
		`UPDATE users SET updated_at = $1, is_voted = true
		WHERE id = $2 AND deleted_at IS NULL`,
		time.Now().In(bkk), req.UserId)
	if err != nil {
		response.ResponseStatus = model.ResponseStatus{
			Code:    be_error.RepositoryCannotUpdate.Code(),
			Message: be_error.RepositoryCannotUpdate.Message(),
		}
		return response, errors.Join(err, errors.New("[Repository.Vote]"))
	}
	if result.RowsAffected() <= 0 {
		err = errors.New("[Repository.Vote]: no user update")
		response.ResponseStatus = model.ResponseStatus{
			Code:    be_error.RepositoryCannotUpdate.Code(),
			Message: be_error.RepositoryCannotUpdate.Message(),
		}
		return response, err
	}

	response.ResponseStatus = model.ResponseStatus{
		Code:    be_error.Success.Code(),
		Message: be_error.Success.Message(),
	}
	return response, nil
}
