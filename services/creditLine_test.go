package services_test

//
//import (
//	"github.com/gin-gonic/gin"
//	"github.com/stretchr/testify/mock"
//	"net/http"
//	"net/http/httptest"
//	"strings"
//	"testing"
//)
//
//func TestInsertTodo(t *testing.T) {
//	client := &database.MockTodoClient{}
//	tests := map[string]struct {
//		payload      string
//		expectedCode int
//	}{
//		"should return 200": {
//			payload:      `{"userId":1,"title":"learning golang","completed":false}`,
//			expectedCode: 200,
//		},
//		"should return 400": {
//			payload:      "invalid json string",
//			expectedCode: 400,
//		},
//	}
//
//	for name, test := range tests {
//		t.Run(name, func(t *testing.T) {
//			client.On("Insert", mock.Anything).Return(models.Todo{}, nil)
//			req, _ := http.NewRequest("POST", "/todos", strings.NewReader(test.payload))
//			rec := httptest.NewRecorder()
//
//			r := gin.Default()
//			r.POST("/todos", handlers.InsertTodo(client))
//			r.ServeHTTP(rec, req)
//
//			if test.expectedCode == 200 {
//				client.AssertExpectations(t)
//			} else {
//				client.AssertNotCalled(t, "Insert")
//			}
//		})
//	}
//}
//
/////*test for adding a naturalPerson*/
////func TestAddNaturalPerson(t *testing.T) {
////	creditLine := services.NewCreditLine()
////	//Ejecutamos el test para cada uno de los casos planteados
////	for i, k := range CasesAddCreditLine {
////		//fmt.Printf("key[%s] value[%s]\n", k, v)
////		fmt.Println("Case N°:", i+1, "for Add an NaturalPerson")
////		//Corremos el test
////		t.Run(k.Name, func(t *testing.T) {
////			result := creditLine.CreateCreditLine(k.Model)
////			//Comparamos el resultado que recibimos del servicio con el que esperamos en nuestro test
////			if (result != nil) != k.ThereIsError {
////				t.Fatalf("%v", result)
////			}
////		},
////		)
////		fmt.Println("----------------")
////	}
////}
