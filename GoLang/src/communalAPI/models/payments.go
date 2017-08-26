package models

import (
	"gopkg.in/mgo.v2/bson"
)

// Represents a payment, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type(

	sAddress struct {
		District string `bson:"district" json:"district"`
		Street string `bson:"street" json:"street"`
		House int `bson:"house" json:"house"`
		PostalCode int `bson:"postalcode" json:"postalcode"`
		HouseType string `bson:"housetype" json:"housetype"`
	}

	sFees struct {
		Year int `bson:"year" json:"year"`
		Month string `bson:"month" json:"month"`
		CentralHeating float64 `bson:"centralheating" json:"centralheating"`
		WaterHeating float64 `bson:"waterheating" json:"waterheating"`
		WaterAndSewer float64 `bson:"waterAndSewer" json:"waterAndSewer"`
		Garbagerecycling float64 `bson:"garbagerecycling" json:"garbagerecycling"`
		Housecleaning float64 `bson:"housecleaning" json:"housecleaning"`
		Gas float64 `bson:"gas" json:"gas"`
		Electricity float64 `bson:"electricity" json:"electricity"`
		Renovationfund float64 `bson:"renovationfund" json:"renovationfund"`
		Administrativefee float64 `bson:"administrativefee" json:"administrativefee"`
		HouseLock float64 `bson:"houseLock" json:"houseLock"`
		HouseInsurance float64 `bson:"houseInsurance" json:"houseInsurance"`
		BankLoan float64 `bson:"bankLoan" json:"bankLoan"`
		Total float64 `bson:"total" json:"total"`
	}

	Payment struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	City        string        `bson:"city" json:"city"`
	Address  []sAddress         `bson:"address" json:"address"`
	Description []sFees         `bson:"fees" json:"fees"`
	}

)