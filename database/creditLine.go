package database

import (
	"errors"
	"github.com/Malushita08/the-credit-line/models"
	"github.com/jinzhu/gorm"
	"math"
	"strings"
	"time"
)

type CreditLineInterface interface {
	CreateCreditLine(creditLineRequestBody models.CreditLineRequestBody) (creditLine models.ResponseBody, err error)
	GetCreditLinesByFoundingName(foundingName string) (CreditLines []models.CreditLine, err error)
}

type CreditLineClient struct {
	DbSession *gorm.DB
}

func (db *CreditLineClient) DefineCreditLine(creditLineRequestBody *models.CreditLineRequestBody, creditLine *models.CreditLine) (err error) {
	timeStp, _ := time.Parse(time.RFC3339, creditLineRequestBody.RequestedDate)
	if err != nil {
		return err
	}
	creditLine.FoundingType = creditLineRequestBody.FoundingType
	creditLine.FoundingName = creditLineRequestBody.FoundingName
	creditLine.CashBalance = creditLineRequestBody.CashBalance
	creditLine.MonthlyRevenue = creditLineRequestBody.MonthlyRevenue
	creditLine.RequestedCreditLine = creditLineRequestBody.RequestedCreditLine
	creditLine.RequestedDate = timeStp
	creditLine.RequestedServerDate = time.Now()
	_ = db.CalculateNotRequestedData(creditLine)
	return nil
}

func (db *CreditLineClient) CalculateNotRequestedData(CreditLine *models.CreditLine) (err error) {
	//Calculate recommendedCreditLine
	if strings.ToUpper(CreditLine.FoundingType) == "SME" {
		CreditLine.RecommendedCreditLine = CreditLine.MonthlyRevenue / 5
	}
	if strings.ToUpper(CreditLine.FoundingType) == "STARTUP" {
		CreditLine.RecommendedCreditLine = math.Max(CreditLine.CashBalance/3, CreditLine.MonthlyRevenue/5)
	}
	//Calculate state
	if CreditLine.RecommendedCreditLine > CreditLine.RequestedCreditLine {
		CreditLine.State = "ACCEPTED"
		CreditLine.LastAcceptedRequestDate = time.Now()
	} else {
		CreditLine.State = "REJECTED"
	}
	//Calculate attemptNumber
	attemptNumber := int64(0)
	_ = db.DbSession.Model(&CreditLine).Where("founding_name = ?", CreditLine.FoundingName).Count(&attemptNumber).Error
	CreditLine.AttemptNumber = attemptNumber + 1
	return nil
}

func (db *CreditLineClient) GetCreditLinesByFoundingName(foundingName string) (CreditLines []models.CreditLine, err error) {
	//creditLines := []models.CreditLine{}
	err = db.DbSession.Where("founding_name = ?", foundingName).Find(&CreditLines).Error
	if err != nil {
		return CreditLines, err
	}
	return CreditLines, nil
}

func (db *CreditLineClient) CreateCreditLine(creditLineRequestBody models.CreditLineRequestBody) (responseBody models.ResponseBody, err error) {
	creditLine := models.CreditLine{}
	creditLineResponseBody := models.CreditLineResponseBody{}
	_ = db.DefineCreditLine(&creditLineRequestBody, &creditLine)
	_ = db.DefineCreditLineResponseBody(&creditLine, &creditLineResponseBody)
	responseBody.Message = "ACCEPTED"
	responseBody.Data = &creditLineResponseBody
	if creditLine.State == "REJECTED" {
		responseBody.Message = "REJECTED"
	}
	err = db.ValidateTimes(&creditLine)
	if err != nil {
		if err.Error() == "A sales agent will contact you" || err.Error() == "Please, wait 30 seconds" || err.Error() == "You've done more than 3 request within the last 2 minutes" {
			responseBody.Data = nil
		}
		responseBody.Message = err.Error()
		return responseBody, err
	}
	err = db.DbSession.Create(&creditLine).Error
	return responseBody, nil
}

func (db *CreditLineClient) ValidateTimes(CreditLine *models.CreditLine) error {
	//Validate the attemptNumber
	if CreditLine.AttemptNumber > 1 {
		//Get the last request
		lastCreditLine := models.CreditLine{}
		_ = db.DbSession.Model(&CreditLine).Where("founding_name = ?", CreditLine.FoundingName).Last(&lastCreditLine).Error

		//Validate the last creditLine state
		if lastCreditLine.State == "ACCEPTED" {
			lastCreditLine.AttemptAcceptedNumber++
			db.DbSession.Save(&lastCreditLine)
			//Validate not more than 3 request within 2 minutes
			afterTwoMinutes := lastCreditLine.LastAcceptedRequestDate.Add(time.Second * 3)
			if CreditLine.RequestedServerDate.Before(afterTwoMinutes) {
				if lastCreditLine.AttemptAcceptedNumber < 3 {
					lastCreditLine.LastAcceptedRequestDate = time.Now()
					db.DbSession.Save(lastCreditLine)
					return errors.New("CONGRATULATIONS!! you already have an approved credit line")
				}
				return errors.New("You've done more than 3 request within the last 2 minutes")
			}
			lastCreditLine.LastAcceptedRequestDate = time.Now()
			lastCreditLine.AttemptAcceptedNumber = 0
			db.DbSession.Save(lastCreditLine)
			return errors.New("CONGRATULATIONS!! you already have an approved credit line")
		} else {
			//Validate 30 seconds later the last request
			afterThirtySeconds := lastCreditLine.RequestedServerDate.Add(time.Second * 3)
			if CreditLine.RequestedServerDate.Before(afterThirtySeconds) {
				return errors.New("Please, wait 30 seconds")
			} else {
				if CreditLine.AttemptNumber <= 3 {
					return nil
				}
				return errors.New("A sales agent will contact you")
			}
		}
	}
	return nil
}

func (db *CreditLineClient) DefineCreditLineResponseBody(creditLine *models.CreditLine, creditLineResponseBody *models.CreditLineResponseBody) (err error) {
	creditLineResponseBody.FoundingType = creditLine.FoundingType
	creditLineResponseBody.FoundingName = creditLine.FoundingName
	creditLineResponseBody.CashBalance = creditLine.CashBalance
	creditLineResponseBody.MonthlyRevenue = creditLine.MonthlyRevenue
	creditLineResponseBody.RequestedCreditLine = creditLine.RequestedCreditLine
	creditLineResponseBody.RequestedDate = creditLine.RequestedDate
	creditLineResponseBody.RequestedServerDate = creditLine.RequestedServerDate
	creditLineResponseBody.RecommendedCreditLine = creditLine.RecommendedCreditLine
	return nil
}
