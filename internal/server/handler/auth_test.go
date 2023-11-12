package handler

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/OurLuv/prefood/internal/model"
	"github.com/OurLuv/prefood/internal/service"
	mock_service "github.com/OurLuv/prefood/internal/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
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
				Email:    "f1lewis@gmail.com",
				Password: "ferrari44s",
			},
			mockBehavior: func(s *mock_service.MockUserService, d data) {
				s.EXPECT().Login(d.Email, d.Password).Return(&user, nil)
			},
			expectedStatusCode: 200,
		},
		{
			name: "empty json",
			inputBody: `{"email" :"", "password": ""}`,
		}
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Init deps
			c := gomock.NewController(t)
			defer c.Finish()

			login := mock_service.NewMockUserService(c)
			tc.mockBehavior(login, tc.inputUser)

			services := service.Service{UserService: login}
			handler := NewHandler(services)

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
			//assert.Equal(t, tc.expectedStatusCode, w.Body.String())
		})
	}
}
