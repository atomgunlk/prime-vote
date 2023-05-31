package handler_test

import (
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

func TestUnitHandler_GetVoteResult(t *testing.T) {
	mockRepo := repository.NewRepository(t)
	// Set mock repo function
	mockRepo.On("GetVoteResult", &model.GetVoteResultRequest{}).
		Return(
			&model.GetVoteResultResponse{
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
		args             model.GetVoteResultRequest
		userArg          *model.User
		wantStatusCode   int
		wantResponseCode string
		wantErr          bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: model.GetVoteResultRequest{},
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
			// Create a new HTTP request.
			req := httptest.NewRequest(http.MethodGet, "/v1/voteresult", nil)
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
				responseObj := model.GetVoteResultResponse{}
				err = json.Unmarshal(body, &responseObj)
				utils.AssertEqual(t, nil, err, "json.Unmarshal")

				utils.AssertEqual(t, tt.wantResponseCode, responseObj.ResponseStatus.Code, "Invalid response code")
			}
		})
	}
}
