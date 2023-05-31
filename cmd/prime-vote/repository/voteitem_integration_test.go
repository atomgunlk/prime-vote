package repository_test

import (
	"fmt"
	"testing"

	"github.com/atomgunlk/prime-vote/cmd/prime-vote/model"
)

func TestIntegration_voteRepository_CreateVoteItem(t *testing.T) {
	tests := []struct {
		name    string
		args    model.CreateVoteItemRequest
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: model.CreateVoteItemRequest{
				Name:        "Pizza 1111",
				Description: "4 cheese dip",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := repo.CreateVoteItem(&tt.args); (err != nil) != tt.wantErr {
				t.Errorf("voteRepository.CreateVoteItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIntegration_voteRepository_ListVoteItem(t *testing.T) {
	tests := []struct {
		name    string
		args    model.ListVoteItemRequest
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: model.ListVoteItemRequest{
				Page: 1,
				Size: 20,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if response, err := repo.ListVoteItem(&tt.args); (err != nil) != tt.wantErr {
				t.Errorf("voteRepository.ListVoteItem() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				fmt.Printf("RESP : %+v\r\n", response.ResponseStatus)
				for i := 0; i < len(response.Items); i++ {
					fmt.Printf("ITEM : %+v\r\n", response.Items[i])
				}
			}
		})
	}
}

func TestIntegration_voteRepository_GetVoteItem(t *testing.T) {
	tests := []struct {
		name    string
		args    model.GetVoteItemRequest
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: model.GetVoteItemRequest{
				ID: 4,
			},
			wantErr: false,
		},
		{
			name: "fail_not_found",
			args: model.GetVoteItemRequest{
				ID: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if response, err := repo.GetVoteItem(&tt.args); (err != nil) != tt.wantErr {
				t.Errorf("voteRepository.GetVoteItem() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				fmt.Printf("RESP : %+v\r\n", response.ResponseStatus)
				fmt.Printf("ITEM : %+v\r\n", response.Item)
			}
		})
	}
}

func TestIntegration_voteRepository_UpdateVoteItem(t *testing.T) {
	tests := []struct {
		name    string
		args    model.UpdateVoteItemRequest
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: model.UpdateVoteItemRequest{
				ID:          3,
				Name:        "test1",
				Description: "desc#1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if response, err := repo.UpdateVoteItem(&tt.args); (err != nil) != tt.wantErr {
				t.Errorf("voteRepository.UpdateVoteItem() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				fmt.Printf("RESP : %+v\r\n", response.ResponseStatus)
			}
		})
	}
}

func TestIntegration_voteRepository_ClearVoteItem(t *testing.T) {
	tests := []struct {
		name    string
		args    model.ClearVoteItemRequest
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: model.ClearVoteItemRequest{
				ID: 3,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if response, err := repo.ClearVoteItem(&tt.args); (err != nil) != tt.wantErr {
				t.Errorf("voteRepository.ClearVoteItem() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				fmt.Printf("RESP : %+v\r\n", response.ResponseStatus)
			}
		})
	}
}

func TestIntegration_voteRepository_DeleteVoteItem(t *testing.T) {
	tests := []struct {
		name    string
		args    model.DeleteVoteItemRequest
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: model.DeleteVoteItemRequest{
				ID: 3,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if response, err := repo.DeleteVoteItem(&tt.args); (err != nil) != tt.wantErr {
				t.Errorf("voteRepository.DeleteVoteItem() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				fmt.Printf("RESP : %+v\r\n", response.ResponseStatus)
			}
		})
	}
}
