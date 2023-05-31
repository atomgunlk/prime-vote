package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/atomgunlk/prime-vote/cmd/prime-vote/handler"
	"github.com/atomgunlk/prime-vote/cmd/prime-vote/model"
	"github.com/atomgunlk/prime-vote/internal/encrypt"
	"github.com/atomgunlk/prime-vote/mocks/cmd/prime-vote/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestUnitHandler_Login(t *testing.T) {
	mockRepo := repository.NewRepository(t)
	// Set mock repo function
	mockRepo.On("GetUserByUsername", "john").
		Return(
			&model.User{
				Username: "john",
				Password: getHashPassword("doe"),
			},
			nil,
		)
	mockRepo.On("GetUserByUsername", "john1").
		Return(
			&model.User{
				Username: "john1",
				Password: getHashPassword("1234"),
			},
			nil,
		)
	mockRepo.On("GetUserByUsername", "john999").
		Return(
			nil,
			errors.New("user not found"),
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
		args           model.LoginRequest
		wantStatusCode int
		wantErr        bool
	}{
		// TODO: Add test cases.
		{
			name: "success_login",
			args: model.LoginRequest{
				Username: "john",
				Password: "doe",
			},
			wantStatusCode: http.StatusOK,
			wantErr:        false,
		},
		{
			name: "fail_login_invalid_password",
			args: model.LoginRequest{
				Username: "john1",
				Password: "doe1",
			},
			wantStatusCode: http.StatusUnauthorized,
			wantErr:        false,
		},
		{
			name: "fail_login_user_notfound",
			args: model.LoginRequest{
				Username: "john999",
				Password: "doe999",
			},
			wantStatusCode: http.StatusUnauthorized,
			wantErr:        false,
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
			req := httptest.NewRequest(http.MethodPost, "/v1/login", bytes.NewBuffer(jsonStr))
			req.Header.Set("Content-Type", "application/json")

			// Use the Fiber app to handle the request.
			resp, err := f.Test(req)
			utils.AssertEqual(t, tt.wantErr, err != nil, "fiber.Test")

			// Check the status code.
			utils.AssertEqual(t, tt.wantStatusCode, resp.StatusCode, "Invalid Status code")

			if tt.wantStatusCode == http.StatusOK {
				// Check the body.
				body, _ := io.ReadAll(resp.Body)
				responseObj := model.LoginReqponse{}
				err = json.Unmarshal(body, &responseObj)
				utils.AssertEqual(t, nil, err, "json.Unmarshal")

				utils.AssertEqual(t, true, len(responseObj.Token) > 0, "Don't Have Token")

				// // extract Claims
				// // Parse the token string without validating the signature
				// token, _, err := new(jwt.Parser).ParseUnverified(responseObj.Token, jwt.MapClaims{})
				// if err != nil {
				// 	fmt.Printf("Error while parsing token: %v\n", err)
				// 	return
				// }
				// if claims, ok := token.Claims.(jwt.MapClaims); ok {
				// 	fmt.Printf("Claims: %+v\n", claims)
				// } else {
				// 	fmt.Println("Invalid token claims")
				// }

				// if !reflect.DeepEqual(got, tt.want) {
				// 	t.Errorf("Repository.GetUserByUsername() = %v, want %v", got, tt.want)
				// }
			}
		})
	}
}

// func BenchmarkUnitLogin(b *testing.B) {
// 	// username := ""
// 	// for i := 0; i < b.N; i++ {
// 	// 	h.Login()
// 	// }
// }

func getHashPassword(password string) string {
	hashed, err := encrypt.HashPassword(password)
	if err != nil {
		return ""
	}
	return hashed
}
