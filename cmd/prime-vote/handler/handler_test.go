package handler_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/atomgunlk/golang-common/pkg/env"
)

// var h *handler.Handler

// var f *fiber.App
var (
	jwtSecret string
)

func setup() {
	env.RequiredEnv("DB_HOST")
	strconv.Atoi(env.RequiredEnv("DB_PORT"))
	env.RequiredEnv("DB_USERNAME")
	env.RequiredEnv("DB_PASSWORD")
	env.RequiredEnv("DB_NAME")
	jwtSecret = env.RequiredEnv("JWT_SECRET")
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}
