package handler

import (
	"net/http"
	"time"

	"github.com/atomgunlk/golang-common/pkg/logger"
	"github.com/atomgunlk/prime-vote/cmd/prime-vote/be_error"
	"github.com/atomgunlk/prime-vote/cmd/prime-vote/model"
	"github.com/atomgunlk/prime-vote/internal/encrypt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(user *model.User, jwtSecret string) (string, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"username": user.Username,
		"id":       user.ID,
		"issuer":   "prime-vote",
		"issuedAt": time.Now().Unix(),
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (h *Handler) Login(c *fiber.Ctx) error {
	req := new(model.LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// validate req
	if err := h.Validate.Struct(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// get user info from DB
	user, err := h.Deps.Repo.GetUserByUsername(req.Username)
	if err != nil {
		logger.WithError(err).Error("[Handler.Login]: repo.GetUserByUsername failed")
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	err = encrypt.CheckPassword(req.Password, user.Password)
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create the Claims
	t, err := GenerateJWT(user, h.Config.JWTSecret)
	if err != nil {
		logger.WithError(err).Error("[Handler.Login]: GenerateJWT failed")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	response := model.LoginReqponse{
		ResponseStatus: model.ResponseStatus{
			Code:    be_error.Success.Code(),
			Message: be_error.Success.Message(),
		},
		Token: t,
	}

	return c.JSON(response)
}
