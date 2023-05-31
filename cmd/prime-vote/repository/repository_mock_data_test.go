package repository_test

import (
	"testing"
)

func TestIntegration_voteRepository_PrepareUserData(t *testing.T) {
	err := repo.PrepareUserData()
	if err != nil {
		t.Errorf("voteRepository.PrepareUserData() error = %v", err)
	}
}

func TestIntegration_voteRepository_PrepareVoteItem(t *testing.T) {
	err := repo.PrepareVoteItem()
	if err != nil {
		t.Errorf("voteRepository.PrepareVoteItem() error = %v", err)
	}
}
