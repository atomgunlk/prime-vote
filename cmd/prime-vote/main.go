package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/atomgunlk/golang-common/pkg/env"
	"github.com/atomgunlk/golang-common/pkg/graceful"
	"github.com/atomgunlk/golang-common/pkg/logger"
	"github.com/atomgunlk/prime-vote/cmd/prime-vote/handler"
	"github.com/atomgunlk/prime-vote/cmd/prime-vote/repository"
	"github.com/gofiber/fiber/v2"
)

func main() {
	f := fiber.New()

	appPort, err := strconv.Atoi(env.RequiredEnv("APP_PORT"))
	if err != nil {
		logger.WithError(err).Panic("[main]: Invalid APP_PORT")
	}

	dbHost := env.RequiredEnv("DB_HOST")
	dbPort, err := strconv.Atoi(env.RequiredEnv("DB_PORT"))
	if err != nil {
		logger.WithError(err).Panic("[main]: Invalid DB_PORT")
	}
	dbName := env.RequiredEnv("DB_NAME")
	dbUsername := env.RequiredEnv("DB_USERNAME")
	dbPassword := env.RequiredEnv("DB_PASSWORD")

	jwtSecret := env.RequiredEnv("JWT_SECRET")

	repo, err := repository.New(
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

	h, err := handler.New(
		&handler.Config{
			JWTSecret: jwtSecret,
		},
		&handler.Dependency{
			Repo: repo,
		},
	)
	if err != nil {
		logger.WithError(err).Panic("[main]: Unable to New handler")
	}

	err = h.InitRoutes(f)
	if err != nil {
		logger.WithError(err).Panic("[main]: Unable to init fiber routes")
	}

	go func() {
		if err := f.Listen(fmt.Sprintf(":%d", appPort)); err != nil {
			logger.Fatal(err)
		}
	}()

	if err := graceful.ListenSignal(func() error {

		err := f.ShutdownWithTimeout(30 * time.Second)
		if err != nil {
			logger.WithError(err).Error("[main]: Shutdown Fiber fail")
		}
		err = repo.Close()
		if err != nil {
			logger.WithError(err).Error("[main]: Shutdown repository fail")
			return err
		}

		logger.Info("[main]: Graceful shutdown success")
		return nil

	}); err != nil {
		logger.WithError(err).Panic("[main]: unable gracefully shutdown the service")
	}
}

// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		token := c.GetHeader("Authorization")
// 		if token != "Bearer YOUR_SECRET_TOKEN" {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 			c.Abort()
// 			return
// 		}
// 		c.Next()
// 	}
// }

// func createVoteItem(c *gin.Context) {
// 	var item VoteItem
// 	if err := c.ShouldBindJSON(&item); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if _, err := db.Exec(context.Background(), "INSERT INTO vote_items (name, description, vote_count) VALUES ($1, $2, $3)", item.Name, item.Description, item.VoteCount); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert vote item"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Vote item created successfully"})
// }

// func getVoteItem(c *gin.Context) {
// 	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
// 	row := db.QueryRow(context.Background(), "SELECT id, name, description, vote_count FROM vote_items WHERE id=$1", id)

// 	var item VoteItem
// 	err := row.Scan(&item.ID, &item.Name, &item.Description, &item.VoteCount)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Vote item not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, item)
// }

// func updateVoteItem(c *gin.Context) {
// 	// Add logic for update
// }

// func deleteVoteItem(c *gin.Context) {
// 	// Add logic for delete
// }
