package models

import "github.com/jinzhu/gorm"

//MODEL
type CreditLine struct {
	ID                  uint    `bson:"_id,omitempty" json:"id,omitempty"`
	FoundingType        string  `bson:"foundingType" json:"foundingType"`
	CashBalance         float64 `bson:"cashBalance" json:"cashBalance"`
	MonthlyRevenue      float64 `bson:"monthlyRevenue" json:"monthlyRevenue"`
	RequestedCreditLine float64 `bson:"requestedCreditLine" json:"requestedCreditLine"`
	RequestedDate       string  `bson:"requestedDate" json:"requestedDate"`
}

//get people
func GetCreditLine(db *gorm.DB, CreditLine *[]CreditLine) (err error) {
	err = db.Find(CreditLine).Error
	if err != nil {
		return err
	}
	return nil
}

//create a book
func CreateCreditLine(db *gorm.DB, CreditLine *CreditLine) (err error) {
	err = db.Create(CreditLine).Error
	if err != nil {
		return err
	}
	return nil
}
