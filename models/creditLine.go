package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type CreditLine struct {
	ID                      uint      `bson:"_id,omitempty" json:"id,omitempty"`
	FoundingType            string    `bson:"foundingType" json:"foundingType"`
	FoundingName            string    `bson:"foundingName" json:"foundingName"`
	CashBalance             float64   `bson:"cashBalance" json:"cashBalance"`
	MonthlyRevenue          float64   `bson:"monthlyRevenue" json:"monthlyRevenue"`
	RequestedCreditLine     float64   `bson:"requestedCreditLine" json:"requestedCreditLine"`
	RequestedDate           time.Time `bson:"requestedDate" json:"requestedDate"`
	RequestedServerDate     time.Time `bson:"requestedServerDate" json:"requestedServerDate"`
	RecommendedCreditLine   float64   `bson:"recommendedCreditLine" json:"recommendedCreditLine"`
	State                   string    `bson:"state" json:"state"`
	AllowedRequest          bool      `bson:"allowedRequest" json:"allowedRequest"`
	AttemptNumber           int64     `bson:"attemptNumber" json:"attemptNumber"`
	AttemptAcceptedNumber   int64     `bson:"attemptAcceptedNumber" json:"attemptAcceptedNumber"`
	LastAcceptedRequestDate time.Time `bson:"lastAcceptedRequestDate" json:"lastAcceptedRequestDate"`
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
