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
	//Complete and calculate not requested data
	CreditLine.RequestedServerDate = time.Now()
	CreditLine.RecommendedCreditLine, CreditLine.State = CalculateRecommendedCreditLineAndState(CreditLine)
	CreditLine.AttemptNumber, _ = CalculateAttemptNumber(CreditLine, db)

	//lastCreditLine :=
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

func CalculateRecommendedCreditLineAndState(CreditLine *CreditLine) (recommendedCreditLine float64, state string) {
	if strings.ToUpper(CreditLine.FoundingType) == "SME" {
		recommendedCreditLine = CreditLine.MonthlyRevenue / 5
	}
	if strings.ToUpper(CreditLine.FoundingType) == "STARTUP" {
		recommendedCreditLine = math.Max(CreditLine.CashBalance/3, CreditLine.MonthlyRevenue/5)
	}
	if recommendedCreditLine > CreditLine.RequestedCreditLine {
		return recommendedCreditLine, "ACCEPTED"
	} else {
		return recommendedCreditLine, "REJECTED"
	}
}

func CalculateAttemptNumber(CreditLine *CreditLine, db *gorm.DB) (attemptNumber int64, err error) {
	_ = db.Model(&CreditLine).Where("founding_name = ?", CreditLine.FoundingName).Count(&attemptNumber).Error
	return attemptNumber + 1, nil
}

func ValidateTimes(CreditLine *CreditLine, db *gorm.DB, lastCreditLine *CreditLine) (bool, error) {
	//Validate the attempt number
	if CreditLine.AttemptNumber > 1 {
		//Get the last request
		_ = db.Model(&CreditLine).Where("founding_name = ?", CreditLine.FoundingName).Last(lastCreditLine).Error

		//Validate the last creditLine state
		if lastCreditLine.State == "ACCEPTED" {
			//Validate not more than 3 request within 2 minutes
			afterTwoMinutes := lastCreditLine.RequestedServerDate.Add(time.Second * 2)
			if CreditLine.AttemptNumber > 3 && afterTwoMinutes.After(CreditLine.RequestedServerDate) {
				return false, errors.New("CONGRATULATIONS!!")
			}
			return true, errors.New("CONGRATULATIONS!! YOU ALREADY HAVE AN APPROVED CREDIT LINE")
		} else {
			//Validate 30 seconds before the last request
			afterThirtySeconds := lastCreditLine.RequestedServerDate.Add(time.Second * 3)
			if CreditLine.RequestedServerDate.Before(afterThirtySeconds) {
				return false, errors.New("WAIT 30 SEC")
			} else {
				if CreditLine.AttemptNumber <= 3 {
					return true, nil
				}
				return false, errors.New("A SALES AGENT WILL CONTACT YOU")
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

type CreditLineRequestBody struct {
	FoundingType        string  `bson:"foundingType" json:"foundingType"`
	FoundingName        string  `bson:"foundingName" json:"foundingName"`
	CashBalance         float64 `bson:"cashBalance" json:"cashBalance"`
	MonthlyRevenue      float64 `bson:"monthlyRevenue" json:"monthlyRevenue"`
	RequestedCreditLine float64 `bson:"requestedCreditLine" json:"requestedCreditLine"`
	RequestedDate       string  `bson:"requestedDate" json:"requestedDate"`
}

type CreditLineResponseBody struct {
	Data    *CreditLine `bson:"data" json:"data"`
	Message string      `bson:"message" json:"message"`
	Error   *string     `bson:"error" json:"error"`
}
