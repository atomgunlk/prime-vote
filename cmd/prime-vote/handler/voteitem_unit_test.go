package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/atomgunlk/prime-vote/cmd/prime-vote/be_error"
	"github.com/atomgunlk/prime-vote/cmd/prime-vote/handler"
	"github.com/atomgunlk/prime-vote/cmd/prime-vote/model"
	"github.com/atomgunlk/prime-vote/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestUnitHandler_CreateVoteItem(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	// Set mock repo function
	mockRepo.On("CreateVoteItem", &model.CreateVoteItemRequest{Name: "Pita", Description: "Kaoklai_test_create_success"}).
		Return(
			&model.CreateVoteItemResponse{
				ResponseStatus: model.ResponseStatus{
					Code:    be_error.Success.Code(),
					Message: be_error.Success.Message(),
				},
			},
			nil,
		)
	mockRepo.On("CreateVoteItem", &model.CreateVoteItemRequest{Name: "Pita", Description: "Kaoklai_test_create_fail_cannot_create"}).
		Return(
			&model.CreateVoteItemResponse{
				ResponseStatus: model.ResponseStatus{
					Code:    be_error.RepositoryCannotCreate.Code(),
					Message: be_error.RepositoryCannotCreate.Message(),
				},
			},
			be_error.RepositoryCannotCreate.Err,
		)

	h, _ := handler.New(
		&handler.Config{
			JWTSecret: jwtSecret,
		},
		&handler.Dependency{
			Repo: mockRepo,
		},
	)
	f := fiber.New()
	h.InitRoutes(f)

	tests := []struct {
		name             string
		args             *model.CreateVoteItemRequest
		userArg          *model.User
		wantStatusCode   int
		wantResponseCode string
		wantErr          bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: &model.CreateVoteItemRequest{
				Name:        "Pita",
				Description: "Kaoklai_test_create_success",
			},
			userArg: &model.User{
				ID:       1,
				Username: "atom",
				IsVoted:  false,
			},
			wantStatusCode:   http.StatusOK,
			wantResponseCode: be_error.Success.Code(),
			wantErr:          false,
		},
		{
			name: "fail_cannot_create",
			args: &model.CreateVoteItemRequest{
				Name:        "Pita",
				Description: "Kaoklai_test_create_fail_cannot_create",
			},
			userArg: &model.User{
				ID:       1,
				Username: "atom",
				IsVoted:  false,
			},
			wantStatusCode:   http.StatusInternalServerError,
			wantResponseCode: be_error.RepositoryCannotCreate.Code(),
			wantErr:          false,
		},
		{
			name: "fail_missing_name",
			args: &model.CreateVoteItemRequest{
				Name:        "",
				Description: "Kaoklai_test_fail_missing_name",
			},
			userArg: &model.User{
				ID:       1,
				Username: "atom",
				IsVoted:  false,
			},
			wantStatusCode:   http.StatusBadRequest,
			wantResponseCode: "",
			wantErr:          false,
		},
		{
			name: "fail_missing_desc",
			args: &model.CreateVoteItemRequest{
				Name:        "Pita",
				Description: "",
			},
			userArg: &model.User{
				ID:       1,
				Username: "atom",
				IsVoted:  false,
			},
			wantStatusCode:   http.StatusBadRequest,
			wantResponseCode: "",
			wantErr:          false,
		},
		{
			name: "fail_missing_body",
			args: nil,
			userArg: &model.User{
				ID:       1,
				Username: "atom",
				IsVoted:  false,
			},
			wantStatusCode:   http.StatusBadRequest,
			wantResponseCode: "",
			wantErr:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqObj := tt.args
			var jsonStr []byte
			if reqObj != nil {
				var err error
				jsonStr, err = json.Marshal(reqObj)
				if err != nil {
					t.Error("unable to marshal json request")
				}
			} else {
				jsonStr = nil
			}

			// Create a new HTTP request.
			req := httptest.NewRequest(http.MethodPost, "/v1/voteitem", bytes.NewBuffer(jsonStr))
			req.Header.Set("Content-Type", "application/json")

			// add authorization header to the req
			jwtToken, err := handler.GenerateJWT(tt.userArg, h.Config.JWTSecret)
			if err != nil {
				t.Error("unable to GenerateJWT")
			}
			req.Header.Add("Authorization", "Bearer "+jwtToken)

			// Use the Fiber app to handle the request.
			resp, err := f.Test(req)
			utils.AssertEqual(t, tt.wantErr, err != nil, "fiber.Test")

			// Check the status code.
			utils.AssertEqual(t, tt.wantStatusCode, resp.StatusCode, "Invalid Status code")

			// Check the body.
			body, _ := io.ReadAll(resp.Body)
			responseObj := model.CreateVoteItemResponse{}
			jsonerr := json.Unmarshal(body, &responseObj)
			// fmt.Printf("RESP : %s\r\n", body)
			if tt.wantStatusCode == http.StatusOK {
				utils.AssertEqual(t, nil, jsonerr, "json.Unmarshal")

				utils.AssertEqual(t, tt.wantResponseCode, responseObj.ResponseStatus.Code, "Invalid response code")
			}
		})
	}
}

func TestUnitHandler_ListVoteItem(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	// Set mock repo function
	mockRepo.On("ListVoteItem", &model.ListVoteItemRequest{Page: 1, Size: 20}).
		Return(
			&model.ListVoteItemResponse{
				ResponseStatus: model.ResponseStatus{
					Code:    be_error.Success.Code(),
					Message: be_error.Success.Message(),
				},
				Items: []model.VoteItem{},
			},
			nil,
		)

	h, _ := handler.New(
		&handler.Config{
			JWTSecret: jwtSecret,
		},
		&handler.Dependency{
			Repo: mockRepo,
		},
	)
	f := fiber.New()
	h.InitRoutes(f)

	tests := []struct {
		name             string
		args             model.ListVoteItemRequest
		userArg          *model.User
		wantStatusCode   int
		wantResponseCode string
		wantErr          bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: model.ListVoteItemRequest{Page: 1, Size: 20},
			userArg: &model.User{
				ID:       1,
				Username: "atom",
				IsVoted:  false,
			},
			wantStatusCode:   http.StatusOK,
			wantResponseCode: be_error.Success.Code(),
			wantErr:          false,
		},
		{
			name: "fail_param",
			args: model.ListVoteItemRequest{Page: 0, Size: 0},
			userArg: &model.User{
				ID:       1,
				Username: "atom",
				IsVoted:  false,
			},
			wantStatusCode:   http.StatusBadRequest,
			wantResponseCode: "",
			wantErr:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new HTTP request.
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/v1/voteitem?page=%d&size=%d", tt.args.Page, tt.args.Size), nil)
			// req.Header.Set("Content-Type", "application/json")

			// add authorization header to the req
			jwtToken, err := handler.GenerateJWT(tt.userArg, h.Config.JWTSecret)
			if err != nil {
				t.Error("unable to GenerateJWT")
			}
			req.Header.Add("Authorization", "Bearer "+jwtToken)

			// Use the Fiber app to handle the request.
			resp, err := f.Test(req)
			utils.AssertEqual(t, tt.wantErr, err != nil, "fiber.Test")

			// Check the status code.
			utils.AssertEqual(t, tt.wantStatusCode, resp.StatusCode, "Invalid Status code")

			if tt.wantStatusCode == http.StatusOK {
				// Check the body.
				body, _ := io.ReadAll(resp.Body)
				responseObj := model.ListVoteItemResponse{}
				err = json.Unmarshal(body, &responseObj)
				utils.AssertEqual(t, nil, err, "json.Unmarshal")

				utils.AssertEqual(t, tt.wantResponseCode, responseObj.ResponseStatus.Code, "Invalid response code")
			}
		})
	}
}

func TestUnitHandler_GetVoteItem(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	// Set mock repo function
	mockRepo.On("GetVoteItem", &model.GetVoteItemRequest{ID: 1}).
		Return(
			&model.GetVoteItemResponse{
				ResponseStatus: model.ResponseStatus{
					Code:    be_error.Success.Code(),
					Message: be_error.Success.Message(),
				},
				Item: model.VoteItem{
					ID:          1,
					Name:        "Pita",
					Description: "Kaoklai",
				},
			},
			nil,
		)
	mockRepo.On("GetVoteItem", &model.GetVoteItemRequest{ID: 999}).
		Return(
			&model.GetVoteItemResponse{
				ResponseStatus: model.ResponseStatus{
					Code:    be_error.RepositoryCannotGet.Code(),
					Message: be_error.RepositoryCannotGet.Message(),
				},
				Item: model.VoteItem{},
			},
			be_error.RepositoryCannotGet.Err,
		)
	h, _ := handler.New(
		&handler.Config{
			JWTSecret: jwtSecret,
		},
		&handler.Dependency{
			Repo: mockRepo,
		},
	)
	f := fiber.New()
	h.InitRoutes(f)

	tests := []struct {
		name           string
		args           model.GetVoteItemRequest
		userArg        *model.User
		wantStatusCode int
		wantErr        bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: model.GetVoteItemRequest{ID: 1},
			userArg: &model.User{
				ID:       1,
				Username: "atom",
				IsVoted:  false,
			},
			wantStatusCode: http.StatusOK,
			wantErr:        false,
		},
		{
			name: "fail_get_item",
			args: model.GetVoteItemRequest{ID: 999},
			userArg: &model.User{
				ID:       1,
				Username: "atom",
				IsVoted:  false,
			},
			wantStatusCode: http.StatusInternalServerError,
			wantErr:        false,
		},
		{
			name: "fail_param",
			args: model.GetVoteItemRequest{ID: 0},
			userArg: &model.User{
				ID:       1,
				Username: "atom",
				IsVoted:  false,
			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new HTTP request.
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/v1/voteitem/%d", tt.args.ID), nil)
			// req.Header.Set("Content-Type", "application/json")

			// add authorization header to the req
			jwtToken, err := handler.GenerateJWT(tt.userArg, h.Config.JWTSecret)
			if err != nil {
				t.Error("unable to GenerateJWT")
			}
			req.Header.Add("Authorization", "Bearer "+jwtToken)

			// Use the Fiber app to handle the request.
			resp, err := f.Test(req)
			utils.AssertEqual(t, tt.wantErr, err != nil, "fiber.Test")

			// Check the status code.
			utils.AssertEqual(t, tt.wantStatusCode, resp.StatusCode, "Invalid Status code")

			if tt.wantStatusCode == http.StatusOK {
				// Check the body.
				body, _ := io.ReadAll(resp.Body)
				responseObj := model.GetVoteItemResponse{}
				err = json.Unmarshal(body, &responseObj)
				utils.AssertEqual(t, nil, err, fmt.Sprintf("json.Unmarshal [%s]", string(body)))

				utils.AssertEqual(t, be_error.Success.Message(), responseObj.ResponseStatus.Message, "Invalid response message")
			}
		})
	}
}

func TestUnitHandler_UpdateVoteItem(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	// Set mock repo function
	mockRepo.On("GetVoteItem",
		&model.GetVoteItemRequest{
			ID: 1,
		}).
		Return(
			&model.GetVoteItemResponse{
				ResponseStatus: model.ResponseStatus{
					Code:    be_error.Success.Code(),
					Message: be_error.Success.Message(),
				},
				Item: model.VoteItem{
					ID:          1,
					Name:        "Pita",
					Description: "Kaoklai",
					VoteCount:   0,
				},
			},
			nil,
		)
	mockRepo.On("UpdateVoteItem",
		&model.UpdateVoteItemRequest{
			ID:          1,
			Name:        "Pita1",
			Description: "Kaoklai1",
		}).
		Return(
			&model.UpdateVoteItemResponse{
				ResponseStatus: model.ResponseStatus{
					Code:    be_error.Success.Code(),
					Message: be_error.Success.Message(),
				},
			},
			nil,
		)

	mockRepo.On("GetVoteItem",
		&model.GetVoteItemRequest{
			ID: 100,
		}).
		Return(
			&model.GetVoteItemResponse{
				ResponseStatus: model.ResponseStatus{
					Code:    be_error.Success.Code(),
					Message: be_error.Success.Message(),
				},
				Item: model.VoteItem{
					ID:          100,
					Name:        "Pita",
					Description: "Kaoklai",
					VoteCount:   123,
				},
			},
			nil,
		)

	h, _ := handler.New(
		&handler.Config{
			JWTSecret: jwtSecret,
		},
		&handler.Dependency{
			Repo: mockRepo,
		},
	)
	f := fiber.New()
	h.InitRoutes(f)

	tests := []struct {
		name             string
		args             *model.UpdateVoteItemRequest
		userArg          *model.User
		wantStatusCode   int
		wantResponseCode string
		wantErr          bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: &model.UpdateVoteItemRequest{
				ID:          1,
				Name:        "Pita1",
				Description: "Kaoklai1",
			},
			userArg: &model.User{
				ID:       1,
				Username: "atom",
				IsVoted:  false,
			},
			wantStatusCode:   http.StatusOK,
			wantResponseCode: be_error.Success.Code(),
			wantErr:          false,
		},
		{
			name: "fail_param",
			args: &model.UpdateVoteItemRequest{
				ID:          0,
				Name:        "",
				Description: "",
			},
			userArg: &model.User{
				ID:       1,
				Username: "atom",
				IsVoted:  false,
			},
			wantStatusCode:   http.StatusBadRequest,
			wantResponseCode: "",
			wantErr:          false,
		},
		{
			name: "fail_body",
			args: &model.UpdateVoteItemRequest{
				ID:          1,
				Name:        "",
				Description: "",
			},
			userArg: &model.User{
				ID:       1,
				Username: "atom",
				IsVoted:  false,
			},
			wantStatusCode:   http.StatusBadRequest,
			wantResponseCode: "",
			wantErr:          false,
		},
		{
			name: "fail_already_voted",
			args: &model.UpdateVoteItemRequest{
				ID:          100,
				Name:        "Pita100",
				Description: "Kaoklai100",
			},
			userArg: &model.User{
				ID:       1,
				Username: "atom",
				IsVoted:  false,
			},
			wantStatusCode:   http.StatusOK,
			wantResponseCode: be_error.HandlerUpdateVoteAlreadyVote.Code(),
			wantErr:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqObj := tt.args
			var jsonStr []byte
			if reqObj != nil {
				var err error
				jsonStr, err = json.Marshal(reqObj)
				if err != nil {
					t.Error("unable to marshal json request")
				}
			} else {
				jsonStr = nil
			}

			// Create a new HTTP request.
			req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/v1/voteitem/%d", tt.args.ID), bytes.NewBuffer(jsonStr))
			req.Header.Set("Content-Type", "application/json")

			// add authorization header to the req
			jwtToken, err := handler.GenerateJWT(tt.userArg, h.Config.JWTSecret)
			if err != nil {
				t.Error("unable to GenerateJWT")
			}
			req.Header.Add("Authorization", "Bearer "+jwtToken)

			// Use the Fiber app to handle the request.
			resp, err := f.Test(req)
			utils.AssertEqual(t, tt.wantErr, err != nil, "fiber.Test")

			// Check the status code.
			utils.AssertEqual(t, tt.wantStatusCode, resp.StatusCode, "Invalid Status code")

			if tt.wantStatusCode == http.StatusOK {
				// Check the body.
				body, _ := io.ReadAll(resp.Body)
				responseObj := model.UpdateVoteItemResponse{}
				err = json.Unmarshal(body, &responseObj)
				utils.AssertEqual(t, nil, err, "json.Unmarshal")

				utils.AssertEqual(t, tt.wantResponseCode, responseObj.ResponseStatus.Code, "Invalid response code")

			}
		})
	}
}

func TestUnitHandler_ClearVoteItem(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	// Set mock repo function
	mockRepo.On("ClearVoteItem", &model.ClearVoteItemRequest{ID: 1}).
		Return(
			&model.ClearVoteItemResponse{
				ResponseStatus: model.ResponseStatus{
					Code:    be_error.Success.Code(),
					Message: be_error.Success.Message(),
				},
			},
			nil,
		)
	mockRepo.On("ClearVoteItem", &model.ClearVoteItemRequest{ID: 999}).
		Return(
			&model.ClearVoteItemResponse{
				ResponseStatus: model.ResponseStatus{
					Code:    be_error.RepositoryCannotUpdate.Code(),
					Message: be_error.RepositoryCannotUpdate.Message(),
				},
			},
			be_error.RepositoryCannotUpdate.Err,
		)
	h, _ := handler.New(
		&handler.Config{
			JWTSecret: jwtSecret,
		},
		&handler.Dependency{
			Repo: mockRepo,
		},
	)
	f := fiber.New()
	h.InitRoutes(f)

	tests := []struct {
		name             string
		args             model.ClearVoteItemRequest
		userArg          *model.User
		wantStatusCode   int
		wantResponseCode string
		wantErr          bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: model.ClearVoteItemRequest{
				ID: 1,
			},
			userArg: &model.User{
				ID:       1,
				Username: "atom",
				IsVoted:  false,
			},
			wantStatusCode:   http.StatusOK,
			wantResponseCode: be_error.Success.Code(),
			wantErr:          false,
		},
		{
			name: "fail_clear_vote_param",
			args: model.ClearVoteItemRequest{
				ID: 0,
			},
			userArg: &model.User{
				ID:       1,
				Username: "atom",
				IsVoted:  false,
			},
			wantStatusCode:   http.StatusBadRequest,
			wantResponseCode: "",
			wantErr:          false,
		},
		{
			name: "fail_clear_vote",
			args: model.ClearVoteItemRequest{
				ID: 999,
			},
			userArg: &model.User{
				ID:       1,
				Username: "atom",
				IsVoted:  false,
			},
			wantStatusCode:   http.StatusInternalServerError,
			wantResponseCode: be_error.RepositoryCannotUpdate.Code(),
			wantErr:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqObj := tt.args
			jsonStr, err := json.Marshal(reqObj)
			if err != nil {
				t.Error("unable to marshal json request")
			}

			// Create a new HTTP request.
			req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/v1/voteitem/%d/clear", tt.args.ID), bytes.NewBuffer(jsonStr))
			// req.Header.Set("Content-Type", "application/json")

			// add authorization header to the req
			jwtToken, err := handler.GenerateJWT(tt.userArg, h.Config.JWTSecret)
			if err != nil {
				t.Error("unable to GenerateJWT")
			}
			req.Header.Add("Authorization", "Bearer "+jwtToken)

			// Use the Fiber app to handle the request.
			resp, err := f.Test(req)
			utils.AssertEqual(t, tt.wantErr, err != nil, "fiber.Test")

			// Check the status code.
			utils.AssertEqual(t, tt.wantStatusCode, resp.StatusCode, "Invalid Status code")

			if tt.wantStatusCode == http.StatusOK {
				// Check the body.
				body, _ := io.ReadAll(resp.Body)
				responseObj := model.ClearVoteItemResponse{}
				err = json.Unmarshal(body, &responseObj)
				utils.AssertEqual(t, nil, err, "json.Unmarshal")

				utils.AssertEqual(t, tt.wantResponseCode, responseObj.ResponseStatus.Code, "Invalid response code")
			}
		})
	}
}

func TestUnitHandler_DeleteVoteItem(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	// Set mock repo function
	mockRepo.On("GetVoteItem",
		&model.GetVoteItemRequest{
			ID: 1,
		}).
		Return(
			&model.GetVoteItemResponse{
				ResponseStatus: model.ResponseStatus{
					Code:    be_error.Success.Code(),
					Message: be_error.Success.Message(),
				},
				Item: model.VoteItem{
					ID:          1,
					Name:        "Pita",
					Description: "Kaoklai",
					VoteCount:   0,
				},
			},
			nil,
		)
	mockRepo.On("GetVoteItem",
		&model.GetVoteItemRequest{
			ID: 100,
		}).
		Return(
			&model.GetVoteItemResponse{
				ResponseStatus: model.ResponseStatus{
					Code:    be_error.Success.Code(),
					Message: be_error.Success.Message(),
				},
				Item: model.VoteItem{
					ID:          100,
					Name:        "Pita",
					Description: "Kaoklai",
					VoteCount:   123,
				},
			},
			nil,
		)
	mockRepo.On("GetVoteItem",
		&model.GetVoteItemRequest{
			ID: 999,
		}).
		Return(
			&model.GetVoteItemResponse{
				ResponseStatus: model.ResponseStatus{
					Code:    be_error.Success.Code(),
					Message: be_error.Success.Message(),
				},
				Item: model.VoteItem{
					ID:          999,
					Name:        "Pita",
					Description: "Kaoklai",
					VoteCount:   0,
				},
			},
			nil,
		)
	mockRepo.On("DeleteVoteItem", &model.DeleteVoteItemRequest{ID: 1}).
		Return(
			&model.DeleteVoteItemResponse{
				ResponseStatus: model.ResponseStatus{
					Code:    be_error.Success.Code(),
					Message: be_error.Success.Message(),
				},
			},
			nil,
		)
	mockRepo.On("DeleteVoteItem", &model.DeleteVoteItemRequest{ID: 999}).
		Return(
			&model.DeleteVoteItemResponse{
				ResponseStatus: model.ResponseStatus{
					Code:    be_error.RepositoryCannotDelete.Code(),
					Message: be_error.RepositoryCannotDelete.Message(),
				},
			},
			be_error.RepositoryCannotDelete.Err,
		)
	h, _ := handler.New(
		&handler.Config{
			JWTSecret: jwtSecret,
		},
		&handler.Dependency{
			Repo: mockRepo,
		},
	)
	f := fiber.New()
	h.InitRoutes(f)

	tests := []struct {
		name             string
		args             model.DeleteVoteItemRequest
		userArg          *model.User
		wantStatusCode   int
		wantResponseCode string
		wantErr          bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: model.DeleteVoteItemRequest{
				ID: 1,
			},
			userArg: &model.User{
				ID:       1,
				Username: "atom",
				IsVoted:  false,
			},
			wantStatusCode:   http.StatusOK,
			wantResponseCode: be_error.Success.Code(),
			wantErr:          false,
		},
		{
			name: "fail_delete_vote_param",
			args: model.DeleteVoteItemRequest{
				ID: 0,
			},
			userArg: &model.User{
				ID:       1,
				Username: "atom",
				IsVoted:  false,
			},
			wantStatusCode:   http.StatusBadRequest,
			wantResponseCode: "",
			wantErr:          false,
		},
		{
			name: "fail_delete_vote_already_vote",
			args: model.DeleteVoteItemRequest{
				ID: 100,
			},
			userArg: &model.User{
				ID:       1,
				Username: "atom",
				IsVoted:  false,
			},
			wantStatusCode:   http.StatusOK,
			wantResponseCode: be_error.HandlerDeleteVoteAlreadyVote.Code(),
			wantErr:          false,
		},
		{
			name: "fail_delete_vote",
			args: model.DeleteVoteItemRequest{
				ID: 999,
			},
			userArg: &model.User{
				ID:       1,
				Username: "atom",
				IsVoted:  false,
			},
			wantStatusCode:   http.StatusInternalServerError,
			wantResponseCode: be_error.RepositoryCannotDelete.Code(),
			wantErr:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new HTTP request.
			req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/v1/voteitem/%d", tt.args.ID), nil)
			// req.Header.Set("Content-Type", "application/json")

			// add authorization header to the req
			jwtToken, err := handler.GenerateJWT(tt.userArg, h.Config.JWTSecret)
			if err != nil {
				t.Error("unable to GenerateJWT")
			}
			req.Header.Add("Authorization", "Bearer "+jwtToken)

			// Use the Fiber app to handle the request.
			resp, err := f.Test(req)
			utils.AssertEqual(t, tt.wantErr, err != nil, "fiber.Test")

			// Check the status code.
			utils.AssertEqual(t, tt.wantStatusCode, resp.StatusCode, "Invalid Status code")

			if tt.wantStatusCode == http.StatusOK {
				// Check the body.
				body, _ := io.ReadAll(resp.Body)
				responseObj := model.DeleteVoteItemResponse{}
				err = json.Unmarshal(body, &responseObj)
				utils.AssertEqual(t, nil, err, "json.Unmarshal")

				utils.AssertEqual(t, tt.wantResponseCode, responseObj.ResponseStatus.Code, "Invalid response code")
			}
		})
	}
}
