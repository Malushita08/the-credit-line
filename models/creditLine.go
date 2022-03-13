package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"math"
	"strings"
	"time"
)

type CreditLine struct {
	ID                    uint      `bson:"_id,omitempty" json:"id,omitempty"`
	FoundingType          string    `bson:"foundingType" json:"foundingType"`
	FoundingName          string    `bson:"foundingName" json:"foundingName"`
	CashBalance           float64   `bson:"cashBalance" json:"cashBalance"`
	MonthlyRevenue        float64   `bson:"monthlyRevenue" json:"monthlyRevenue"`
	RequestedCreditLine   float64   `bson:"requestedCreditLine" json:"requestedCreditLine"`
	RequestedDate         time.Time `bson:"requestedDate" json:"requestedDate"`
	RequestedServerDate   time.Time `bson:"requestedServerDate" json:"requestedServerDate"`
	RecommendedCreditLine float64   `bson:"recommendedCreditLine" json:"recommendedCreditLine"`
	State                 string    `bson:"state" json:"state"`
	AllowedRequest        bool      `bson:"allowedRequest" json:"allowedRequest"`
	AttemptNumber         int64     `bson:"attemptNumber" json:"attemptNumber"`
	AttemptAcceptedNumber int64     `bson:"attemptAcceptedNumber" json:"attemptAcceptedNumber"`
}

func GetCreditLines(db *gorm.DB, CreditLines *[]CreditLine) (err error) {
	err = db.Find(CreditLines).Error
	if err != nil {
		return err
	}
	return nil
}

func GetCreditLine(db *gorm.DB, CreditLine *CreditLine, id string) (err error) {
	err = db.Where("id = ?", id).First(CreditLine).Error
	if err != nil {
		return err
	}
	return nil
}

func GetCreditLinesByFoundingName(db *gorm.DB, CreditLines *[]CreditLine, foundingName string) (err error) {
	err = db.Where("founding_name = ?", foundingName).Find(CreditLines).Error
	if err != nil {
		return err
	}
	return nil
}

func CreateCreditLine(db *gorm.DB, CreditLine *CreditLine, LastCreditLine *CreditLine) (err error) {
	CreditLine.AllowedRequest, err = ValidateTimes(CreditLine, db, LastCreditLine)
	if err != nil {
		return err
	}
	err = db.Create(CreditLine).Error
	if err != nil {
		return err
	}
	return nil
}

func ValidateTimes(CreditLine *CreditLine, db *gorm.DB, lastCreditLine *CreditLine) (bool, error) {
	//Validate the attempt number
	if CreditLine.AttemptNumber > 1 {
		//Get the last request
		_ = db.Model(&CreditLine).Where("founding_name = ?", CreditLine.FoundingName).Last(&lastCreditLine).Error

		//Validate the last creditLine state
		if lastCreditLine.State == "ACCEPTED" {
			//Validate not more than 3 request within 2 minutes
			afterTwoMinutes := lastCreditLine.RequestedServerDate.Add(time.Second * 2)
			if CreditLine.AttemptNumber > 3 && afterTwoMinutes.After(CreditLine.RequestedServerDate) {
				return true, errors.New("Wait 2 minutes please!!")
			}
			return true, errors.New("CONGRATULATIONS!! you already have an approved credit line")
		} else {
			//Validate 30 seconds before the last request
			afterThirtySeconds := lastCreditLine.RequestedServerDate.Add(time.Second * 3)
			if CreditLine.RequestedServerDate.Before(afterThirtySeconds) {
				return false, errors.New("Wait 30 seconds please")
			} else {
				if CreditLine.AttemptNumber <= 3 {
					return true, nil
				}
				return true, errors.New("A sales agent will contact you")
			}
		}
	}
	return true, nil
}

func UpdateCreditLine(db *gorm.DB, CreditLine *CreditLine) (err error) {
	db.Save(CreditLine)
	return nil
}

func DeleteCreditLine(db *gorm.DB, CreditLine *CreditLine, id string) (err error) {
	db.Where("id = ?", id).Delete(CreditLine)
	return nil
}

func CalculateNotRequestedData(CreditLine *CreditLine, db *gorm.DB) (err error) {
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
	} else {
		CreditLine.State = "REJECTED"
	}
	//Calculate attemptNumber
	attemptNumber := int64(0)
	_ = db.Model(&CreditLine).Where("founding_name = ?", CreditLine.FoundingName).Count(&attemptNumber).Error
	CreditLine.AttemptNumber = attemptNumber + 1
	return nil
}

func DefineCreditLine(db *gorm.DB, creditLineRequestBody *CreditLineRequestBody, creditLine *CreditLine) (err error) {
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
	_ = CalculateNotRequestedData(creditLine, db)
	return nil
}

func DefineCreditLineResponseBody(creditLine *CreditLine, creditLineResponseBody *CreditLineResponseBody) (err error) {
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

type CreditLineRequestBody struct {
	FoundingType        string  `bson:"foundingType" json:"foundingType"`
	FoundingName        string  `bson:"foundingName" json:"foundingName"`
	CashBalance         float64 `bson:"cashBalance" json:"cashBalance"`
	MonthlyRevenue      float64 `bson:"monthlyRevenue" json:"monthlyRevenue"`
	RequestedCreditLine float64 `bson:"requestedCreditLine" json:"requestedCreditLine"`
	RequestedDate       string  `bson:"requestedDate" json:"requestedDate"`
}

type CreditLineResponseBody struct {
	FoundingType          string    `bson:"foundingType" json:"foundingType"`
	FoundingName          string    `bson:"foundingName" json:"foundingName"`
	CashBalance           float64   `bson:"cashBalance" json:"cashBalance"`
	MonthlyRevenue        float64   `bson:"monthlyRevenue" json:"monthlyRevenue"`
	RequestedCreditLine   float64   `bson:"requestedCreditLine" json:"requestedCreditLine"`
	RequestedDate         time.Time `bson:"requestedDate" json:"requestedDate"`
	RequestedServerDate   time.Time `bson:"requestedServerDate" json:"requestedServerDate"`
	RecommendedCreditLine float64   `bson:"recommendedCreditLine" json:"recommendedCreditLine"`
}

type ResponseBody struct {
	Message string                  `bson:"message" json:"message"`
	Data    *CreditLineResponseBody `bson:"data" json:"data"`
	Error   *string                 `bson:"error" json:"error"`
}
