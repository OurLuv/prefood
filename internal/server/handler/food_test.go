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

func TestHandler__CreateFood(t *testing.T) {
	type mockBehavior func(s *mock_service.MockFoodService, data model.Food)

	testCases := []struct {
		name               string
		inputBody          string
		inputFood          model.Food
		mockBehavior       mockBehavior
		ExpectedStatusCode int
	}{
		{
			name:      "OK",
			inputBody: `{ "name": "Chicken with rice", "price":299 }`,
			inputFood: model.Food{Name: "Chicken with rice", Price: 299},
			mockBehavior: func(s *mock_service.MockFoodService, data model.Food) {
				s.EXPECT().Create(data).Return(nil)
			},
			ExpectedStatusCode: 200,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Init deps
			c := gomock.NewController(t)
			defer c.Finish()

			s := mock_service.NewMockFoodService(c)
			tc.mockBehavior(s, tc.inputFood)
			services := service.Service{
				FoodService: s,
			}
			handler := NewHandler(services, setupLogger())

			// router
			r := mux.NewRouter()
			r.HandleFunc("/restaurants/{restaurant_id}/menu/add", handler.CreateFood).Methods("POST")

			// sending request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/restaurants/{restaurant_id}/menu/add", bytes.NewBufferString(tc.inputBody))
			r.ServeHTTP(w, req)

			// assert
			assert.Equal(t, tc.ExpectedStatusCode, w.Code)
			//assert.Equal(t, tc.ExpectedResult.InStock, w.Code)
		})
	}
}
