package repository_test

import (
	"fmt"
	"testing"

	"github.com/atomgunlk/prime-vote/cmd/prime-vote/model"
)

func TestIntegration_voteRepository_Vote(t *testing.T) {
	tests := []struct {
		name    string
		args    model.VoteRequest
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: model.VoteRequest{
				ID: 4,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if response, err := repo.Vote(&tt.args); (err != nil) != tt.wantErr {
				t.Errorf("voteRepository.Vote() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				fmt.Printf("RESP : %+v\r\n", response.ResponseStatus)
			}
		})
	}
}
