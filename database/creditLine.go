package database

import (
	"github.com/Malushita08/the-credit-line/models"
	"github.com/jinzhu/gorm"
	"math"
	"strings"
	"time"
)

type CreditLineInterface interface {
	CreateCreditLine(creditLineRequestBody models.CreditLineRequestBody) (creditLine models.CreditLine, err error)
}

type CreditLineClient struct {
	DbSession *gorm.DB
}

func (db *CreditLineClient) CreateCreditLine(creditLineRequestBody models.CreditLineRequestBody) (creditLine models.CreditLine, err error) {
	err = db.DefineCreditLine(&creditLineRequestBody, &creditLine)
	if err != nil {
		return models.CreditLine{}, err
	}
	//CreditLine.AllowedRequest, err = ValidateTimes(CreditLine, db, LastCreditLine)
	//if err != nil {
	//	return err
	//}
	err = db.DbSession.Create(&creditLine).Error
	//db.Model(&CreditLine)
	//err = db.DbSession.Model(&creditLine).Find().First(&creditLine).Error
	//if err != nil {
	//	return err
	//}
	return creditLine, nil
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

//func (db *CreditLineClient )CreateCreditLine(CreditLine *models.CreditLine, LastCreditLine *models.CreditLine) (err error) {
//	//CreditLine.AllowedRequest, err = ValidateTimes(CreditLine, db, LastCreditLine)
//	if err != nil {
//		return err
//	}
//	err = db.DbSession.Create(CreditLine).Error
//	if err != nil {
//		return err
//	}
//	return nil
//}

//func ValidateTimes(CreditLine *CreditLine, db *gorm.DB, lastCreditLine *CreditLine) (bool, error) {
//	//Validate the attemptNumber
//	if CreditLine.AttemptNumber > 1 {
//		//Get the last request
//		_ = db.Model(&CreditLine).Where("founding_name = ?", CreditLine.FoundingName).Last(&lastCreditLine).Error
//
//		//Validate the last creditLine state
//		if lastCreditLine.State == "ACCEPTED" {
//			//lastCreditLine.AttemptAcceptedNumber++
//			//db.Save(lastCreditLine)
//
//			//Validate not more than 3 request within 2 minutes
//			afterTwoMinutes := lastCreditLine.LastAcceptedRequestDate.Add(time.Second * 3)
//			if CreditLine.RequestedServerDate.Before(afterTwoMinutes) {
//				return false, errors.New("You've done more than 3 request within 2 minutes")
//
//				//return false, errors.New("Please, wait two minutes")
//			}
//			//if lastCreditLine.AttemptAcceptedNumber < 3 {
//			//	lastCreditLine.AttemptAcceptedNumber++
//			//	if CreditLine.RequestedServerDate.Before(afterTwoMinutes) {
//			//		return false, errors.New("Please, wait two minutes")
//			//	}
//			//	lastCreditLine.LastAcceptedRequestDate = time.Now()
//			//	//lastCreditLine.AttemptAcceptedNumber = 0
//			//	db.Save(lastCreditLine)
//			//	return false, errors.New("aaaaa")
//			//}
//			//lastCreditLine.AttemptAcceptedNumber = 0
//			lastCreditLine.LastAcceptedRequestDate = time.Now()
//			db.Save(lastCreditLine)
//			return true, errors.New("CONGRATULATIONS!! you already have an approved credit line")
//
//			//return true, errors.New("CONGRATULATIONS!! you already have an approved credit line")
//			//fmt.Printf("aaaaaa", lastCreditLine)
//			//fmt.Printf("bbbbbb", CreditLine)
//			//if CreditLine.RequestedServerDate.Before(afterTwoMinutes) {
//			//	fmt.Printf("entrororo aki")
//			//	lastCreditLine.AttemptAcceptedNumber++
//			//	//lastCreditLine.LastAcceptedRequestDate = time.Now()
//			//	db.Save(lastCreditLine)
//			//	if lastCreditLine.AttemptAcceptedNumber >= 3 {
//			//		lastCreditLine.AttemptAcceptedNumber = 0
//			//		db.Save(lastCreditLine)
//			//		//return false, errors.New("aaaa")
//			//		return false, errors.New("Please, wait two minutes")
//			//	}
//			//	//return false, errors.New("Please, wait two minutes")
//			//}
//			//if lastCreditLine.AttemptAcceptedNumber >= 3 && afterTwoMinutes.After(CreditLine.RequestedServerDate) {
//			//	lastCreditLine.AttemptAcceptedNumber = 0
//			//	db.Save(lastCreditLine)
//			//	return false, errors.New("Please, wait two minutes")
//			//}
//			//lastCreditLine.LastAcceptedRequestDate = time.Now()
//
//		} else {
//			//Validate 30 seconds before the last request
//			afterThirtySeconds := lastCreditLine.RequestedServerDate.Add(time.Second * 3)
//
//			if CreditLine.RequestedServerDate.Before(afterThirtySeconds) {
//				return false, errors.New("Please, wait 30 seconds")
//			} else {
//				if CreditLine.AttemptNumber <= 3 {
//					return true, nil
//				}
//				return true, errors.New("A sales agent will contact you")
//			}
//		}
//	}
//	return true, nil
//}
//
