package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/atomgunlk/prime-vote/cmd/prime-vote/model"
	"github.com/jackc/pgx/v5"
)

//go:generate mockery --name Repository
type Repository interface {
	// Vote
	Vote(*model.VoteRequest) (*model.VoteResponse, error)

	CreateVoteItem(*model.CreateVoteItemRequest) (*model.CreateVoteItemResponse, error)
	ListVoteItem(req *model.ListVoteItemRequest) (*model.ListVoteItemResponse, error)
	GetVoteItem(req *model.GetVoteItemRequest) (*model.GetVoteItemResponse, error)
	UpdateVoteItem(req *model.UpdateVoteItemRequest) (*model.UpdateVoteItemResponse, error)
	ClearVoteItem(req *model.ClearVoteItemRequest) (*model.ClearVoteItemResponse, error)
	DeleteVoteItem(req *model.DeleteVoteItemRequest) (*model.DeleteVoteItemResponse, error)
	GetVoteResult(req *model.GetVoteResultRequest) (*model.GetVoteResultResponse, error)

	// User
	GetUserByUsername(username string) (*model.User, error)
	GetUserByID(id uint64) (*model.User, error)

	// Close
	Close() error

	PrepareUserData() error
	PrepareVoteItem() error
}
type Config struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string

	OperationTimeout int // second
}

type voteRepository struct {
	DB               *pgx.Conn
	operationTimeout time.Duration
}

func New(cfg *Config) (Repository, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Bangkok",
		cfg.Host, cfg.Username, cfg.Password, cfg.Database, cfg.Port,
	)
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, errors.Join(err, errors.New("[repository.New]"))
	}

	err = conn.Ping(context.Background())
	if err != nil {
		return nil, errors.New("[repository.New]: can not connect to DB")
	}

	// Auto migration
	err = autoMigrator(cfg)
	if err != nil {
		return nil, err
	}

	operationTimeout := time.Duration(cfg.OperationTimeout) * time.Second
	if operationTimeout == 0 {
		return nil, errors.New("[repository.New]: operation timeout is invalid")
	}

	return &voteRepository{DB: conn, operationTimeout: operationTimeout}, nil
}

func (r *voteRepository) defaultContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), r.operationTimeout)
}

func (r *voteRepository) Close() error {
	return r.DB.Close(context.Background())
}
