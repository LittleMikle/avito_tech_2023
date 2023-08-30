package handler

import (
	"bytes"
	"errors"
	tech "github.com/LittleMikle/avito_tech_2023"
	"github.com/LittleMikle/avito_tech_2023/pkg/service"
	mock_service "github.com/LittleMikle/avito_tech_2023/pkg/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_createSegment(t *testing.T) {
	type mockBehavior func(r *mock_service.MockSegmentation, segment tech.Segment)

	testTable := []struct {
		name                 string
		inputBody            string
		inputSegment         tech.Segment
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"title": "AVITO_VOICE_MESSAGES"}`,
			inputSegment: tech.Segment{
				Title: "AVITO_VOICE_MESSAGES",
			},
			mockBehavior: func(r *mock_service.MockSegmentation, segment tech.Segment) {
				r.EXPECT().CreateSegment(segment).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:                 "Bad Input",
			inputSegment:         tech.Segment{},
			mockBehavior:         func(r *mock_service.MockSegmentation, segment tech.Segment) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"EOF"}`,
		},
		{
			name:      "Service Error",
			inputBody: `{"title": "Melushev_AVITO_DISCOUNT_30"}`,
			inputSegment: tech.Segment{
				Title: "Melushev_AVITO_DISCOUNT_30",
			},
			mockBehavior: func(r *mock_service.MockSegmentation, segment tech.Segment) {
				r.EXPECT().CreateSegment(segment).Return(0, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockSegmentation(c)
			testCase.mockBehavior(repo, testCase.inputSegment)

			services := &service.Service{Segmentation: repo}
			handler := Handler{
				services,
			}

			r := gin.New()
			r.POST("/create", handler.createSegment)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/create",
				bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedResponseBody)
		})
	}
}

func TestHandler_deleteSegment(t *testing.T) {
	type mockBehavior func(r *mock_service.MockSegmentation, userId int)

	testTable := []struct {
		name                 string
		segmentId            int
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			segmentId: 1,
			mockBehavior: func(r *mock_service.MockSegmentation, userId int) {
				r.EXPECT().DeleteSegment(userId).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockSegmentation(c)

			services := &service.Service{Segmentation: repo}
			handler := Handler{
				services,
			}

			r := gin.New()
			r.DELETE("/:id", handler.deleteSegment)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("DELETE", "/:id",
				bytes.NewBufferString(""))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedResponseBody)
		})
	}
}
