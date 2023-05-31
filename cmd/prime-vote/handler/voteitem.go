package handler

import (
	"net/http"

	"github.com/atomgunlk/golang-common/pkg/logger"
	"github.com/atomgunlk/prime-vote/cmd/prime-vote/be_error"
	"github.com/atomgunlk/prime-vote/cmd/prime-vote/model"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreateVoteItem(c *fiber.Ctx) error {
	req := new(model.CreateVoteItemRequest)

	if err := c.BodyParser(req); err != nil {
		logger.WithError(err).Error("[CreateVoteItem]: cannot parse request")
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

	response, err := h.Deps.Repo.CreateVoteItem(req)
	if err != nil {
		logger.WithError(err).Error("[CreateVoteItem]: repo.CreateVoteItem failed")
		return c.Status(http.StatusInternalServerError).JSON(response)
	}
	return c.JSON(response)
}

func (h *Handler) ListVoteItem(c *fiber.Ctx) error {
	req := new(model.ListVoteItemRequest)

	if err := c.QueryParser(req); err != nil {
		logger.WithError(err).Error("[ListVoteItem]: cannot parse request")
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

	response, err := h.Deps.Repo.ListVoteItem(req)
	if err != nil {
		logger.WithError(err).Error("[ListVoteItem]: repo.ListVoteItem failed")
		return c.Status(http.StatusInternalServerError).JSON(response)
	}
	return c.JSON(response)
}

func (h *Handler) GetVoteItem(c *fiber.Ctx) error {
	req := new(model.GetVoteItemRequest)
	if err := c.ParamsParser(req); err != nil {
		logger.WithError(err).Error("[GetVoteItem]: cannot parse request")
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

	response, err := h.Deps.Repo.GetVoteItem(req)
	if err != nil {
		logger.WithError(err).Error("[GetVoteItem]: repo.GetVoteItem failed")
		return c.Status(http.StatusInternalServerError).JSON(response)
	}
	return c.JSON(response)
}

func (h *Handler) UpdateVoteItem(c *fiber.Ctx) error {
	req := new(model.UpdateVoteItemRequest)
	if err := c.ParamsParser(req); err != nil {
		logger.WithError(err).Error("[UpdateVoteItem]: cannot parse request")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse request",
		})
	}

	if err := c.BodyParser(req); err != nil {
		logger.WithError(err).Error("[UpdateVoteItem]: cannot parse request")
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

	// check votecount
	getResponse, err := h.Deps.Repo.GetVoteItem(&model.GetVoteItemRequest{ID: req.ID})
	if err != nil {
		logger.WithError(err).Error("[UpdateVoteItem]: repo.GetVoteItem failed")
		return c.Status(http.StatusInternalServerError).JSON(getResponse)
	}

	if getResponse.Item.VoteCount > 0 {
		response := model.UpdateVoteItemResponse{
			ResponseStatus: model.ResponseStatus{
				Code:    be_error.HandlerUpdateVoteAlreadyVote.Code(),
				Message: be_error.HandlerUpdateVoteAlreadyVote.Message(),
			},
		}
		return c.JSON(response)
	}

	response, err := h.Deps.Repo.UpdateVoteItem(req)
	if err != nil {
		logger.WithError(err).Error("[UpdateVoteItem]: repo.UpdateVoteItem failed")
		return c.Status(http.StatusInternalServerError).JSON(response)
	}
	return c.JSON(response)
}

func (h *Handler) ClearVoteItem(c *fiber.Ctx) error {
	req := new(model.ClearVoteItemRequest)
	if err := c.ParamsParser(req); err != nil {
		logger.WithError(err).Error("[ClearVoteItemItem]: cannot parse request")
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

	response, err := h.Deps.Repo.ClearVoteItem(req)
	if err != nil {
		logger.WithError(err).Error("[ClearVoteItemItem]: repo.ClearVoteItemItem failed")
		return c.Status(http.StatusInternalServerError).JSON(response)
	}
	return c.JSON(response)
}

func (h *Handler) DeleteVoteItem(c *fiber.Ctx) error {
	req := new(model.DeleteVoteItemRequest)
	if err := c.ParamsParser(req); err != nil {
		logger.WithError(err).Error("[DeleteVoteItem]: cannot parse request")
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

	// check votecount
	getResponse, err := h.Deps.Repo.GetVoteItem(&model.GetVoteItemRequest{ID: req.ID})
	if err != nil {
		logger.WithError(err).Error("[DeleteVoteItem]: repo.GetVoteItem failed")
		return c.Status(http.StatusInternalServerError).JSON(getResponse)
	}

	if getResponse.Item.VoteCount > 0 {
		response := model.DeleteVoteItemResponse{
			ResponseStatus: model.ResponseStatus{
				Code:    be_error.HandlerDeleteVoteAlreadyVote.Code(),
				Message: be_error.HandlerDeleteVoteAlreadyVote.Message(),
			},
		}
		return c.JSON(response)
	}

	response, err := h.Deps.Repo.DeleteVoteItem(req)
	if err != nil {
		logger.WithError(err).Error("[DeleteVoteItem]: repo.DeleteVoteItem failed")
		return c.Status(http.StatusInternalServerError).JSON(response)
	}
	return c.JSON(response)
}
