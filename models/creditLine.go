package models

import (
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
	RecommendedCreditLine   float64   `bson:"recommendedCreditLine" json:"recommendedCreditLine"`
	State                   string    `bson:"state" json:"state"`
	AllowedRequest          bool      `bson:"allowedRequest" json:"allowedRequest"`
	AttemptNumber           int64     `bson:"attemptNumber" json:"attemptNumber"`
	AttemptAcceptedNumber   int64     `bson:"attemptAcceptedNumber" json:"attemptAcceptedNumber"`
	LastAcceptedRequestDate time.Time `bson:"lastAcceptedRequestDate" json:"lastAcceptedRequestDate"`
}

// CreditLineRequestBody
// swagger:model
type CreditLineRequestBody struct {
	FoundingType        string  `bson:"foundingType" json:"foundingType"`
	FoundingName        string  `bson:"foundingName" json:"foundingName"`
	CashBalance         float64 `bson:"cashBalance" json:"cashBalance"`
	MonthlyRevenue      float64 `bson:"monthlyRevenue" json:"monthlyRevenue"`
	RequestedCreditLine float64 `bson:"requestedCreditLine" json:"requestedCreditLine"`
}

type CreditLineResponseBody struct {
	FoundingType          string    `bson:"foundingType" json:"foundingType"`
	FoundingName          string    `bson:"foundingName" json:"foundingName"`
	CashBalance           float64   `bson:"cashBalance" json:"cashBalance"`
	MonthlyRevenue        float64   `bson:"monthlyRevenue" json:"monthlyRevenue"`
	RequestedCreditLine   float64   `bson:"requestedCreditLine" json:"requestedCreditLine"`
	RequestedDate         time.Time `bson:"requestedDate" json:"requestedDate"`
	RecommendedCreditLine float64   `bson:"recommendedCreditLine" json:"recommendedCreditLine"`
}

type ResponseBody struct {
	Message string                  `bson:"message" json:"message"`
	Data    *CreditLineResponseBody `bson:"data" json:"data"`
	Error   *string                 `bson:"error" json:"error"`
}
