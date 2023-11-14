package handler

import (
	"bytes"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/OurLuv/prefood/internal/model"
	"github.com/OurLuv/prefood/internal/service"
	mock_service "github.com/OurLuv/prefood/internal/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slog"
)

type data struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func TestHandler__login(t *testing.T) {
	type mockBehavior func(s *mock_service.MockUserService, d data)
	user := model.User{
		Id:       3,
		Email:    "f1lewis@gmail.com",
		Password: "asdsadasd23e2dx3",
	}
	_ = user
	// test cases
	testCases := []struct {
		name                string
		inputBody           string
		inputUser           data
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"email" :"f1lewis@gmail.com", "password": "ferrari44s"}`,
			inputUser: data{
				"f1lewis@gmail.com",
				"ferrari44s",
			},
			mockBehavior: func(s *mock_service.MockUserService, d data) {
				s.EXPECT().Login(d.Email, d.Password).Return(&user, nil)
			},
			expectedStatusCode: 200,
		},
		{
			name:               "empty json",
			inputBody:          `{"email":"", "password": ""}`,
			mockBehavior:       func(s *mock_service.MockUserService, d data) {},
			expectedStatusCode: 400,
		},
		{
			name:               "email",
			inputBody:          `{"email":"123456", "password": "123456"}`,
			mockBehavior:       func(s *mock_service.MockUserService, d data) {},
			expectedStatusCode: 400,
		},
	}

	// go through all test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Init deps
			c := gomock.NewController(t)
			defer c.Finish()

			login := mock_service.NewMockUserService(c)
			tc.mockBehavior(login, tc.inputUser)

			services := service.Service{UserService: login}
			handler := NewHandler(services, setupLogger())

			//Test server
			r := mux.NewRouter()
			r.HandleFunc("/login", handler.login).Methods("POST")

			// Test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/login", bytes.NewBufferString(tc.inputBody))

			// Perform request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, tc.expectedStatusCode, w.Code)
			//assert.Equal(t, tc.expectedRequestBody, w.Body.String())
		})
	}
}

func TestHandler__signup(t *testing.T) {
	type mockBehavior func(mock *mock_service.MockUserService, u model.User)
	user := model.User{
		Firstname: "John",
		Lastname:  "Doe",
		Password:  "123456",
		Email:     "john.doe@example.com",
	}
	testCases := []struct {
		name         string
		inputBody    string
		inputUser    model.User
		mockBehavior mockBehavior
		expectedCode int
		expectedBody string
	}{
		// testcase 1
		{
			name:      "Valid input",
			inputBody: `{ "firstname": "John", "lastname": "Doe", "password": "123456", "email": "john.doe@example.com" }`,
			inputUser: model.User{
				Firstname: "John",
				Lastname:  "Doe",
				Password:  "123456",
				Email:     "john.doe@example.com",
			},
			mockBehavior: func(mock *mock_service.MockUserService, u model.User) {
				mock.EXPECT().Create(u).Return(nil)
			},
			expectedCode: 200,
			expectedBody: "",
		},
		// testcase 2
		{
			name:         "Invalid input",
			inputBody:    `{ "firstname": "", "password": "123456", "email": "john.doe@example.com" }`,
			inputUser:    model.User{Firstname: "John", Password: "123456", Email: "john.doe@example.com"},
			mockBehavior: func(mock *mock_service.MockUserService, u model.User) {},
			expectedCode: 400,
			expectedBody: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			//creating controller
			c := gomock.NewController(t)
			defer c.Finish()

			// init deps
			s := mock_service.NewMockUserService(c)
			tc.mockBehavior(s, user)

			services := service.Service{
				UserService: s,
			}
			handler := NewHandler(services, setupLogger())

			// start server
			r := mux.NewRouter()
			r.HandleFunc("/create-test", handler.signup).Methods("POST")

			// test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/create-test", bytes.NewBufferString(tc.inputBody))

			//perform
			r.ServeHTTP(w, req)

			//assert
			assert.Equal(t, tc.expectedCode, w.Code)
		})
	}
}

func setupLogger() *slog.Logger {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	return log
}
