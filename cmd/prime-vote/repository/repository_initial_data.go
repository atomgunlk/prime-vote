package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/atomgunlk/golang-common/pkg/logger"
	"github.com/atomgunlk/prime-vote/cmd/prime-vote/model"
	"github.com/atomgunlk/prime-vote/internal/encrypt"
)

func (r *voteRepository) PrepareUserData() error {
	bkk, _ := time.LoadLocation("Asia/Bangkok")
	// insert 100 user
	for i := 0; i < 100; i++ {
		hashed, _ := encrypt.HashPassword(fmt.Sprintf("test%d", i))
		user := model.User{
			Username: fmt.Sprintf("test%d", i),
			Password: hashed,
		}

		_, err := r.DB.Exec(
			context.Background(),
			"INSERT INTO users (created_at, updated_at, username, password, is_voted) VALUES ($1, $2, $3, $4, false)",
			time.Now().In(bkk), time.Now().In(bkk), user.Username, user.Password,
		)
		if err != nil {
			logger.WithError(err).Errorf("[Repository.PrepareUserData]")
		}
	}

	return nil
}

func (r *voteRepository) PrepareVoteItem() error {
	bkk, _ := time.LoadLocation("Asia/Bangkok")
	for i := 0; i < 10; i++ {
		item := model.VoteItem{
			Name:        fmt.Sprintf("testVoteItem%d", i),
			Description: fmt.Sprintf("testVoteItem%d Description", i),
		}
		_, err := r.DB.Exec(
			context.Background(),
			"INSERT INTO vote_items (created_at, updated_at, name, description, vote_count) VALUES ($1, $2, $3, $4, 0)",
			time.Now().In(bkk), time.Now().In(bkk), item.Name, item.Description,
		)
		if err != nil {
			logger.WithError(err).Errorf("[Repository.PrepareVoteItem]")
		}
	}

	return nil
}
