package handler_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/atomgunlk/prime-vote/cmd/prime-vote/be_error"
	"github.com/atomgunlk/prime-vote/cmd/prime-vote/handler"
	"github.com/atomgunlk/prime-vote/cmd/prime-vote/model"
	"github.com/atomgunlk/prime-vote/mocks/cmd/prime-vote/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestUnitHandler_Vote(t *testing.T) {
	mockRepo := repository.NewRepository(t)
	// Set mock repo function
	mockRepo.On("GetUserByID", uint64(1)).
		Return(
			&model.User{
				ID:       1,
				Username: "atom",
				IsVoted:  false,
			},
			nil,
		)

	mockRepo.On("Vote", &model.VoteRequest{ID: 1, UserId: 1}).
		Return(
			&model.VoteResponse{
				ResponseStatus: model.ResponseStatus{
					Code:    be_error.Success.Code(),
					Message: be_error.Success.Message(),
				},
			},
			nil,
		)
	// mockRepo.On("Vote", &model.VoteRequest{ID: 2}).
	// 	Return(
	// 		&model.VoteResponse{
	// 			ResponseStatus: model.ResponseStatus{
	// 				Code:    be_error.RepositoryCannotCreate.Code(),
	// 				Message: be_error.RepositoryCannotCreate.Message(),
	// 			},
	// 		},
	// 		be_error.RepositoryCannotCreate.Err,
	// 	)

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
		voteArg          *model.VoteRequest
		userArg          *model.User
		wantStatusCode   int
		wantResponseCode string
		wantErr          bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			voteArg: &model.VoteRequest{
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqObj := tt.voteArg
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
			req := httptest.NewRequest(http.MethodPost, "/v1/vote", bytes.NewBuffer(jsonStr))
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
