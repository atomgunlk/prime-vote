package repository

import (
	"errors"
	"fmt"

	"github.com/atomgunlk/prime-vote/cmd/prime-vote/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// autoMigrator Use GORM ONLY For migrate
func autoMigrator(cfg *Config) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Bangkok",
		cfg.Host, cfg.Username, cfg.Password, cfg.Database, cfg.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.Join(err, errors.New("[repository.AutoMigrator]: unable to open connection"))
	}

	err = db.AutoMigrate(
		&model.User{},
		&model.VoteItem{},
	)
	if err != nil {
		return errors.Join(err, errors.New("[repository.AutoMigrator]: unable to migrate"))
	}

	return nil
}
