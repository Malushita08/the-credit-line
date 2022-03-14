package handlers_test

import (
	"github.com/Malushita08/the-credit-line/database"
	"github.com/Malushita08/the-credit-line/handlers"
	"github.com/Malushita08/the-credit-line/models"

	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

func TestInsertTodo(t *testing.T) {
	client := &database.MockCreditLineClient{}
	tests := map[string]struct {
		payload      string
		expectedCode int
	}{
		"should return 200": {
			payload: `{
				"foundingType": "startup",
				"foundingName": "J",
				"cashBalance": 300,
				"monthlyRevenue": 3000,
				"requestedCreditLine": 45,
				"requestedDate": "2022-03-10T16:59:19.29889741-05:00"
			}`,
			expectedCode: 200,
		},
		"should return rejected": {
			payload: `{
				"foundingType": "startup",
				"foundingName": "J",
				"cashBalance": 300,
				"monthlyRevenue": 3000,
				"requestedCreditLine": 450000000,
				"requestedDate": "2022-03-10T16:59:19.29889741-05:00"
			}`,
			expectedCode: 400,
		},
		"should return 426": {
			payload: `{
				"foundingType": "startup",
				"foundingName": "J",
				"cashBalance": 300,
				"monthlyRevenue": 3000,
				"requestedCreditLine": 45,
				"requestedDate": "2022-03-10T16:59:19.29889741-05:00"
			}`,
			expectedCode: 426,
		},
		"should return 400": {
			payload:      "invalid json string",
			expectedCode: 400,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			client.On("CreateCreditLine", mock.Anything).Return(models.CreditLineResponseBody{}, nil)
			req, _ := http.NewRequest("POST", "/creditLines", strings.NewReader(test.payload))
			rec := httptest.NewRecorder()

			r := gin.Default()
			r.POST("/creditLines", handlers.InsertCreditLine(client))
			r.ServeHTTP(rec, req)

			if test.expectedCode == 200 {
				client.AssertExpectations(t)
			} else {
				client.AssertNotCalled(t, "Insert")
			}
		})
	}
}
