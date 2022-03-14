package handlers_test

import (
	"github.com/Malushita08/the-credit-line/database"
	"github.com/Malushita08/the-credit-line/handlers"
	"github.com/Malushita08/the-credit-line/models"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

func TestSearchTodos(t *testing.T) {
	client := &database.MockCreditLineClient{}

	tests := map[string]struct {
		payload      string
		expectedCode int
	}{
		"should return 200": {
			payload:      `{"title": "learning golang"}`,
			expectedCode: 200,
		},
		"should return 400": {
			payload:      "invalid json string",
			expectedCode: 400,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			client.On("Search", mock.Anything).Return([]models.CreditLine{}, nil)

			req, _ := http.NewRequest("GET", "/todos?q="+test.payload, nil)
			rec := httptest.NewRecorder()

			r := gin.Default()
			r.GET("/todos", handlers.GetCreditLinesByFoundingName(client))
			r.ServeHTTP(rec, req)

			if test.expectedCode == 200 {
				client.AssertExpectations(t)
			} else {
				client.AssertNotCalled(t, "Search")
			}

		})
	}

}
