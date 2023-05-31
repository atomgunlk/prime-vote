package handler

import (
	"net/http"

	"github.com/atomgunlk/golang-common/pkg/logger"
	"github.com/atomgunlk/prime-vote/cmd/prime-vote/be_error"
	"github.com/atomgunlk/prime-vote/cmd/prime-vote/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func (h *Handler) Vote(c *fiber.Ctx) error {
	req := new(model.VoteRequest)

	if err := c.BodyParser(req); err != nil {
		logger.WithError(err).Error("[Handler.Vote]: cannot parse request")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse request",
		})
	}

	// validate req
	if err := h.Validate.Struct(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// get user from jwt
	userjwt, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid auth",
		})
	}
	claims := userjwt.Claims.(jwt.MapClaims)
	userID, ok := claims["id"].(float64)
	if !ok {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid user id",
		})
	}

	// get user from db
	user, err := h.Deps.Repo.GetUserByID(uint64(userID))
	if err != nil {
		logger.WithError(err).Error("[Handler.Vote]: repo.GetUserByID failed")
		return c.Status(http.StatusInternalServerError).
			JSON(model.VoteResponse{
				ResponseStatus: model.ResponseStatus{
					// Code:    be_error.HandlerUpdateVoteAlreadyVote.Code(),
					// Message: be_error.HandlerUpdateVoteAlreadyVote.Message(),
				},
			})
	}
	if user.IsVoted {
		return c.Status(http.StatusOK).
			JSON(model.VoteResponse{
				ResponseStatus: model.ResponseStatus{
					Code:    be_error.HandlerVoteUserAlreadyVoted.Code(),
					Message: be_error.HandlerVoteUserAlreadyVoted.Message(),
				},
			})
	}

	req.UserId = user.ID
	response, err := h.Deps.Repo.Vote(req)
	if err != nil {
		logger.WithError(err).Error("[Handler.Vote]: repo.Vote failed")
		return c.Status(http.StatusInternalServerError).JSON(response)
	}
	return c.JSON(response)
}
