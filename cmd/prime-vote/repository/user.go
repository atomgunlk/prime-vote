package repository

import (
	"errors"

	"github.com/atomgunlk/prime-vote/cmd/prime-vote/model"
	"github.com/jackc/pgtype"
)

func (r *voteRepository) GetUserByUsername(username string) (*model.User, error) {
	user := model.User{
		Username: username,
	}
	ctx, cancel := r.defaultContext()
	defer cancel()

	row := r.DB.QueryRow(ctx, "SELECT id, created_at, updated_at, username, password, is_voted FROM users WHERE username = $1 AND deleted_at IS NULL", username)

	createdAt := pgtype.Timestamptz{}
	updatedAt := pgtype.Timestamptz{}
	err := row.Scan(&user.ID, &createdAt, &updatedAt, &user.Username, &user.Password, &user.IsVoted)
	if err != nil {
		return nil, errors.Join(err, errors.New("[Repository.GetUserByUsername]"))
	}

	user.CreatedAt = createdAt.Time
	user.UpdatedAt = updatedAt.Time

	return &user, nil
}

func (r *voteRepository) GetUserByID(id uint64) (*model.User, error) {
	user := model.User{
		ID: id,
	}
	ctx, cancel := r.defaultContext()
	defer cancel()

	row := r.DB.QueryRow(ctx, "SELECT id, created_at, updated_at, username, password, is_voted FROM users WHERE id = $1 AND deleted_at IS NULL", id)
	createdAt := pgtype.Timestamptz{}
	updatedAt := pgtype.Timestamptz{}
	err := row.Scan(&user.ID, &createdAt, &updatedAt, &user.Username, &user.Password, &user.IsVoted)
	if err != nil {
		return nil, errors.Join(err, errors.New("[Repository.GetUserByID]"))
	}

	user.CreatedAt = createdAt.Time
	user.UpdatedAt = updatedAt.Time

	return &user, nil
}
