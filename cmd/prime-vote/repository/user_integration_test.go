package repository_test

import (
	"testing"

	"github.com/atomgunlk/prime-vote/cmd/prime-vote/model"
)

func Test_voteRepository_GetUserByUsername(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				username: "test0",
			},
			want: &model.User{
				ID:       1,
				Username: "test0",
			},
			wantErr: false,
		},
		{
			name: "fail_notfound",
			args: args{
				username: "test-1",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.GetUserByUsername(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("voteRepository.GetUserByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				if got != tt.want {
					t.Errorf("voteRepository.GetUserByUsername() = %v, want %v", got, tt.want)
					return
				} else {
					// true
					return
				}
			}
			if got.ID != tt.want.ID {
				t.Errorf("voteRepository.GetUserByUsername() = %v, want %v", got, tt.want)
			}
			if got.Username != tt.want.Username {
				t.Errorf("voteRepository.GetUserByUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_voteRepository_GetUserByID(t *testing.T) {
	type args struct {
		id uint64
	}
	tests := []struct {
		name    string
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				id: 1,
			},
			want: &model.User{
				ID:       1,
				Username: "test0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.GetUserByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("voteRepository.GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.ID != tt.want.ID {
				t.Errorf("voteRepository.GetUserByID() = %v, want %v", got, tt.want)
			}
			if got.Username != tt.want.Username {
				t.Errorf("voteRepository.GetUserByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
