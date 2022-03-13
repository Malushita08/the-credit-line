package services_test

type TestCreditLine struct {
	Name         string
	Model        interface{}
	ID           int64
	ThereIsError bool
}

/* set up global variables*/

//var IdLegalPerson = primitive.NewObjectID()

//Cases for add a legalPerson

var CasesAddCreditLine = []TestCreditLine{
	{
		Name: "Add a new accepted creditLine",
		Model: map[string]interface{}{
			"foundingType":        "startup",
			"foundingName":        "J",
			"cashBalance":         300,
			"monthlyRevenue":      3000,
			"requestedCreditLine": 45,
			"requestedDate":       "2022-03-10T16:59:19.29889741-05:00"},
		ThereIsError: false,
	},
}

/*Cases for  Update a legalPerson*/
//
//var CasesUpdateLegalPerson = []TestCreditLine{
//	{
//		Name: "Update an legalPerson  already added",
//		ID:   IdLegalPerson,
//		Model: bson.M{
//			"_id,omitempty": IdLegalPerson,
//			"businessName":  "updateTest",
//			"ruc":           "updateTest",
//		},
//		ThereIsError: false,
//	},
//	{
//		Name: "Update an not added legalPerson",
//		ID:   primitive.NewObjectID(),
//		Model: bson.M{
//			"businessName": "updateTest",
//			"ruc":          "updateTest",
//		},
//		ThereIsError: true,
//	},
//}
//
///*Cases for get a legalPerson*/
//var CasesGetLegalPerson = []TestCreditLine{
//	{
//		Name:         "Get legalPerson already added",
//		ID:           IdLegalPerson,
//		ThereIsError: false,
//	},
//	{
//		Name:         "Get not added legalPerson",
//		ID:           primitive.NewObjectID(),
//		ThereIsError: true,
//	},
//}
//
///*Cases for removing legalPerson*/
//var CasesRemovingLegalPerson = []TestCreditLine{
//	{
//		Name:         "Removing an already added legalPerson",
//		ID:           IdLegalPerson,
//		ThereIsError: false,
//	},
//	{
//		Name:         "Removing a legalPerson not added",
//		ID:           primitive.NewObjectID(),
//		ThereIsError: true,
//	},
//}
