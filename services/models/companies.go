package models

import (
	"gopkg.in/mgo.v2/bson"
	"errors"
)

type Company struct {
	UUID				bson.ObjectId `bson:"_id,omitempty"`
	System 				Base
	Name 				string
	Code 				string
	Address 			Address
	Type 				string
	Version 			int
	Order 				float64
}

const CompanyCollectionName = "companies"

func (company *Company) Validate() error {
 	if company.Name == "" {
		return errors.New("invalid field content <name>")
	}
	return nil
}

func VaidateCompany(data interface{}) (out interface{}, err error) {
	var company Company
	err = SetStruct(data, &company)
	if err != nil {
		return company, err
	}
	err = company.Validate()
	return company, err
}
