package be_error

import (
	"errors"
	"fmt"
)

type BeError struct {
	code string
	Err  error
}

func (r *BeError) Code() string {
	return r.code
}
func (r *BeError) Error() string {
	return fmt.Sprintf("%v", r.Err)
}
func (r *BeError) Message() string {
	return fmt.Sprintf("%v", r.Err)
}
func (r *BeError) WrapMessage(m string) string {
	return fmt.Sprintf("%v, %s", r.Err, m)
}
func (r *BeError) FullError() string {
	return fmt.Sprintf("code %s: err %v", r.code, r.Err)
}

var (
	Success = BeError{"0000", errors.New("success")}

	// Common 1xxx

	ParamInvalid = BeError{"1001", errors.New("invalid param")}
	TypeInvalid  = BeError{"1002", errors.New("invalid type")}

	///////////////////////
	// Handler 2xxx
	///////////////////////
	HandlerUpdateVoteAlreadyVote = BeError{"2000", errors.New("vote item already vote")}
	HandlerDeleteVoteAlreadyVote = BeError{"2001", errors.New("vote item already vote")}
	HandlerVoteUserAlreadyVoted  = BeError{"2010", errors.New("user already voted")}

	///////////////////////
	// Repository 3xxx
	///////////////////////
	RepositoryCommon       = BeError{"3000", errors.New("repository common error")}
	RecordNotfound         = BeError{"3001", errors.New("record not found in repository")}
	RepositoryCannotCreate = BeError{"3002", errors.New("can not create a record")}
	RepositoryCannotGet    = BeError{"3003", errors.New("can not get a record")}
	RepositoryCannotUpdate = BeError{"3004", errors.New("can not update a record")}
	RepositoryCannotDelete = BeError{"3005", errors.New("can not delete a record")}
)
