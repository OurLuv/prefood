package handler

import (
	"bytes"
	"context"
	"mime/multipart"
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
			inputBody: `{ "name": "Chicken with rice", "price":299}`,
			inputFood: model.Food{Name: "Chicken with rice", RestaurantId: 1, Price: 299, Image: "png"},
			mockBehavior: func(s *mock_service.MockFoodService, data model.Food) {
				s.EXPECT().Create(data).Return(&model.Food{Name: "Chicken with rice", Price: 299, Image: "sdfsd.png"}, nil)
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

			var buf bytes.Buffer
			w := multipart.NewWriter(&buf)
			fw, err := w.CreateFormField("food")
			if err != nil {
				t.Fatal(err)
			}
			fw.Write([]byte(tc.inputBody))
			_, err = w.CreateFormFile("image", "file.png")
			if err != nil {
				t.Fatal(err)
			}
			// if _, err := io.Copy(fw, file); err != nil {
			// 	t.Fatal(err)
			// }
			w.Close()
			// router
			r := mux.NewRouter()
			r.HandleFunc("/restaurants/{restaurant_id}/menu/add", handler.CreateFood).Methods("POST")

			// sending request
			rr := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/restaurants/{restaurant_id}/menu/add", &buf)
			req.Header.Set("Content-Type", w.FormDataContentType())
			newCtx := context.WithValue(req.Context(), "restaurant", &model.Restaurant{Id: 1})

			r.ServeHTTP(rr, req.WithContext(newCtx))

			// assert
			assert.Equal(t, tc.ExpectedStatusCode, rr.Code)
			//assert.Equal(t, tc.ExpectedResult.InStock, w.Code)
		})
	}
}
