package models

import (
	"github.com/jinzhu/gorm"
	"math"
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
}

//type CreditLineRequestBody struct {
//	ID                  uint      `bson:"_id,omitempty" json:"id,omitempty"`
//	FoundingType        string    `bson:"foundingType" json:"foundingType"`
//	CashBalance         float64   `bson:"cashBalance" json:"cashBalance"`
//	MonthlyRevenue      float64   `bson:"monthlyRevenue" json:"monthlyRevenue"`
//	RequestedCreditLine float64   `bson:"requestedCreditLine" json:"requestedCreditLine"`
//	RequestedDate       time.Time `bson:"requestedDate" json:"requestedDate"`
//	RequestedServerDate time.Time `bson:"requestedServerDate" json:"requestedServerDate"`
//}

func GetCreditLines(db *gorm.DB, CreditLine *[]CreditLine) (err error) {
	err = db.Find(CreditLine).Error
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

func CreateCreditLine(db *gorm.DB, CreditLine *CreditLine) (err error) {
	foundingType := CreditLine.FoundingType
	cashBalance := float64(CreditLine.CashBalance)
	monthlyRevenue := float64(CreditLine.MonthlyRevenue)
	CreditLine.RecommendedCreditLine = RecommendedCreditLine(foundingType, cashBalance, monthlyRevenue)
	CreditLine.RequestedServerDate = time.Now()
	CreditLine.State = CalculateState(CreditLine.RecommendedCreditLine, CreditLine.RequestedCreditLine)

	//fmt.Printf("creditLine?: ", CreditLine, "\n")
	err = db.Create(CreditLine).Error
	if err != nil {
		return err
	}
	return nil
}

//func CreateCreditLine2(db *gorm.DB, CreditLineRequestBody *CreditLineRequestBody) (err error) {
//	err = db.Create(CreditLineRequestBody).Error
//	if err != nil {
//		return err
//	}
//	return nil
//}

func UpdateCreditLine(db *gorm.DB, CreditLine *CreditLine) (err error) {
	db.Save(CreditLine)
	return nil
}

func DeleteCreditLine(db *gorm.DB, CreditLine *CreditLine, id string) (err error) {
	db.Where("id = ?", id).Delete(CreditLine)
	return nil
}

func RecommendedCreditLine(foundingType string, cashBalance float64, monthlyRevenue float64) (recommendedCreditLine float64) {
	if foundingType == "SME" {
		recommendedCreditLine = monthlyRevenue / 3
	}
	if foundingType == "Startup" {
		recommendedCreditLine = math.Max(cashBalance/3, monthlyRevenue/5)
	}
	return recommendedCreditLine
}

func CalculateState(recommendedCreditLine float64, requestedCreditLine float64) (state string) {
	if recommendedCreditLine > requestedCreditLine {
		return "ACCEPTED"
	} else {
		return "REJECTED"
	}
}
