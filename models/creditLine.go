package models

import (
	"github.com/jinzhu/gorm"
	"math"
	"strings"
	"time"
)

type CreditLine struct {
	ID                    uint      `bson:"_id,omitempty" json:"id,omitempty"`
	FoundingType          string    `bson:"foundingType" json:"foundingType"`
	CashBalance           float64   `bson:"cashBalance" json:"cashBalance"`
	MonthlyRevenue        float64   `bson:"monthlyRevenue" json:"monthlyRevenue"`
	RequestedCreditLine   float64   `bson:"requestedCreditLine" json:"requestedCreditLine"`
	RequestedDate         time.Time `bson:"requestedDate" json:"requestedDate"`
	RequestedServerDate   time.Time `bson:"requestedServerDate" json:"requestedServerDate"`
	RecommendedCreditLine float64   `bson:"recommendedCreditLine" json:"recommendedCreditLine"`
	State                 string    `bson:"state" json:"state"`
	AllowedRequest        bool      `bson:"allowedRequest" json:"allowedRequest"`
}

type CreditLineRequestBody struct {
	FoundingType        string  `bson:"foundingType" json:"foundingType"`
	CashBalance         float64 `bson:"cashBalance" json:"cashBalance"`
	MonthlyRevenue      float64 `bson:"monthlyRevenue" json:"monthlyRevenue"`
	RequestedCreditLine float64 `bson:"requestedCreditLine" json:"requestedCreditLine"`
	RequestedDate       string  `bson:"requestedDate" json:"requestedDate"`
}

type ResponseBody struct {
	Data    *CreditLine `bson:"data" json:"data"`
	Error   *string     `bson:"error" json:"error"`
	Message string      `bson:"message" json:"message"`
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

func CreateCreditLine(db *gorm.DB, CreditLine *CreditLine, LastCreditLine *CreditLine) (err error) {
	CreditLine.RequestedServerDate = time.Now()
	foundingType := CreditLine.FoundingType
	cashBalance := CreditLine.CashBalance
	monthlyRevenue := CreditLine.MonthlyRevenue

	CreditLine.RecommendedCreditLine, CreditLine.State = CalculateState(
		foundingType,
		cashBalance,
		monthlyRevenue,
		CreditLine.RequestedCreditLine)

	//lastCreditLine :=
	CreditLine.AllowedRequest, _ = ValidateTimes(CreditLine.State, db, LastCreditLine)
	err = db.Create(CreditLine).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateCreditLine(db *gorm.DB, CreditLine *CreditLine) (err error) {
	db.Save(CreditLine)
	return nil
}

func DeleteCreditLine(db *gorm.DB, CreditLine *CreditLine, id string) (err error) {
	db.Where("id = ?", id).Delete(CreditLine)
	return nil
}

func CalculateState(foundingType string, cashBalance float64, monthlyRevenue float64, requestedCreditLine float64) (recommendedCreditLine float64, state string) {
	if strings.ToUpper(foundingType) == "SME" {
		recommendedCreditLine = monthlyRevenue / 5
	}
	if strings.ToUpper(foundingType) == "STARTUP" {
		recommendedCreditLine = math.Max(cashBalance/3, monthlyRevenue/5)
	}
	if recommendedCreditLine > requestedCreditLine {
		return recommendedCreditLine, "ACCEPTED"
	} else {
		return recommendedCreditLine, "REJECTED"
	}
}

//func GetCreditLine(db *gorm.DB, CreditLine *CreditLine, id string) (err error) {
//	err = db.Where("id = ?", id).First(CreditLine).Error
//	if err != nil {
//		return err
//	}
//	return nil
//}
func ValidateTimes(state string, db *gorm.DB, LastCreditLine *CreditLine) (allowedRequest bool, err error) {
	//Get the last request
	//var creditLines []CreditLine
	err = db.Last(LastCreditLine).Error
	if err != nil {
		return false, err
	}
	//fmt.Printf("jiijiiiii?: ", LastCreditLine)

	//err = db.Where("id = ?", id).First(CreditLine).Error
	//if state == "ACCEPTED" {
	//
	//} else {
	//	//Validate 30 seconds before the last
	//	if lastCreditLine.R
	//}
	return true, nil
}
