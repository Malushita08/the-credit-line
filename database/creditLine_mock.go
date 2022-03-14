package database

import (
	"github.com/Malushita08/the-credit-line/models"

	"github.com/stretchr/testify/mock"
)

type MockCreditLineClient struct {
	mock.Mock
}

func (m *MockCreditLineClient) CreateCreditLine(creditLineRequestBody models.CreditLineRequestBody) (creditLine models.ResponseBody, err error) {
	args := m.Called(creditLineRequestBody)
	return args.Get(0).(models.ResponseBody), args.Error(1)
}

func (m *MockCreditLineClient) GetCreditLinesByFoundingName(foundingName string) (CreditLines []models.CreditLine, err error) {
	args := m.Called(foundingName)
	return args.Get(0).([]models.CreditLine), args.Error(1)
}
