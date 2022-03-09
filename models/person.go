package models

import "github.com/jinzhu/gorm"

//MODEL
type Person struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	City      string `json:"city"`
}

//get people
func GetPeople(db *gorm.DB, Person *[]Person) (err error) {
	err = db.Find(Person).Error
	if err != nil {
		return err
	}
	return nil
}
