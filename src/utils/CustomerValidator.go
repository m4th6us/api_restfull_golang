package utils

import (
	"api_restfull/src/models"
	"fmt"
)

func VerifyExistCustomerForCreate(arrayCustomers []models.Customer, customer models.Customer) bool {

	for _, v := range arrayCustomers{

		if v.CustomerId == customer.CustomerId{

			return true

		}

	}

	return false

}

func VerifyExistCustomerForFindOne(idCustomer int, arrayCustomers []models.Customer) (*models.Customer, error){

	for _, v := range arrayCustomers{

		if v.CustomerId == idCustomer{

			return &v, nil

		}

	}

	return nil, fmt.Errorf("customer with ID %d not found", idCustomer)

}