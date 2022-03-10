package models

import "github.com/jinzhu/gorm"

type CreditLine struct {
	ID                  uint    `bson:"_id,omitempty" json:"id,omitempty"`
	FoundingType        string  `bson:"foundingType" json:"foundingType"`
	CashBalance         float64 `bson:"cashBalance" json:"cashBalance"`
	MonthlyRevenue      float64 `bson:"monthlyRevenue" json:"monthlyRevenue"`
	RequestedCreditLine float64 `bson:"requestedCreditLine" json:"requestedCreditLine"`
	RequestedDate       string  `bson:"requestedDate" json:"requestedDate"`
}

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
