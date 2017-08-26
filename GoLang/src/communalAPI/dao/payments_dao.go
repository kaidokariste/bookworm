package dao

import (
	"log"

	. "communalAPI/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PaymentDAO struct {
	DBUrl string
	AuthDatabase string

}

var db *mgo.Database

const (
	COLLECTION = "payment"
)

func (m *PaymentDAO) Connect() {
	session, err := mgo.Dial(m.DBUrl)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.AuthDatabase)
}

// Find list of payments
func (m *PaymentDAO) FindAll() ([]Payment, error) {
	var payments []Payment
	err := db.C(COLLECTION).Find(bson.M{}).All(&payments)
	return payments, err
}

// Find a payment by its id
func (m *PaymentDAO) FindById(id string) (Payment, error) {
	var payment Payment
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&payment)
	return payment, err
}

// Insert a paymrent into database
func (m *PaymentDAO) Insert(payment Payment) error {
	err := db.C(COLLECTION).Insert(&payment)
	return err
}

// Delete an existing payment
func (m *PaymentDAO) Delete(payment Payment) error {
	err := db.C(COLLECTION).Remove(&payment)
	return err
}

// Update an existing payment
func (m *PaymentDAO) Update(payment Payment) error {
	err := db.C(COLLECTION).UpdateId(payment.ID, &payment)
	return err
}
