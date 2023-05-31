package repository_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/atomgunlk/golang-common/pkg/env"
	"github.com/atomgunlk/golang-common/pkg/logger"
	"github.com/atomgunlk/prime-vote/cmd/prime-vote/repository"
)

var repo repository.Repository

func setup() {
	var err error
	dbHost := env.RequiredEnv("DB_HOST")
	dbPort, err := strconv.Atoi(env.RequiredEnv("DB_PORT"))
	if err != nil {
		logger.WithError(err).Panic("[test.repository]: Invalid DB_PORT")
	}
	dbUsername := env.RequiredEnv("DB_USERNAME")
	dbPassword := env.RequiredEnv("DB_PASSWORD")
	dbName := env.RequiredEnv("DB_NAME")

	repo, err = repository.New(
		&repository.Config{
			Host:     dbHost,
			Port:     dbPort,
			Database: dbName,

			Username: dbUsername,
			Password: dbPassword,

			OperationTimeout: 10,
		},
	)
	if err != nil {
		logger.WithError(err).Panic("[main]: Unable to New repository")
	}
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}
