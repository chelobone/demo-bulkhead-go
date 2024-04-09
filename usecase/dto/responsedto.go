package dto

import (
	"github.com/chelobone/demo_bulkhead_go/domain/model"
)

// Response dto

// Estructura de response para creación de customer
type Response struct {
	ResponsePartyReferenceDataDirectory *ResponsePartyReferenceDataDirectory
}

// Estructura de response para obtención de cliente por Id
type ResponseModel struct {
	ResponsePartyReferenceDataDirectory ResponsePartyReferenceDataDirectoryModel
}

// Estructura de response para obtención de todos los clientes
type ResponseArray struct {
	ResponsePartyReferenceDataDirectory []ResponsePartyReferenceDataDirectoryModel
}

type ResponsePartyReferenceDataDirectory struct {
	Customer ResponseCustomer
}

// Para colección de objeto Customer
type ResponsePartyReferenceDataDirectoryModel struct {
	Customer     ResponseCustomerModel
	Address      ResponseAddressModel
	EmailAddress ResponseElectronicAddressModel
}

// Objeto Customer solo con referencia al identificador único
type ResponseCustomer struct {
	CustomerReference int `json:"customerReference"`
}

// Objeto Customer solo con referencia a
type ResponseCustomerModel struct {
	CustomerName     string `json:"customerName"`
	CustomerLastName string `json:"customerLastName"`
}

type ResponseAddressModel struct {
	ResidentialAddress string `json:"residentialAddress"`
}

type ResponseElectronicAddressModel struct {
	EmailAddress string `json:"emailAddress"`
}

func GenerateResponseModel(customerReference int) Response {
	return Response{
		ResponsePartyReferenceDataDirectory: &ResponsePartyReferenceDataDirectory{
			Customer: ResponseCustomer{CustomerReference: customerReference},
		},
	}
}

func GenerateResponseCustomerModel(customer model.Customer) ResponsePartyReferenceDataDirectoryModel {
	return ResponsePartyReferenceDataDirectoryModel{
		Customer:     ResponseCustomerModel{CustomerName: customer.Nombre, CustomerLastName: customer.Apellido},
		Address:      ResponseAddressModel{ResidentialAddress: customer.Direccion},
		EmailAddress: ResponseElectronicAddressModel{EmailAddress: customer.Correo},
	}
}

func GenerateResponseCustomerModelArray(customers model.CustomerSlice) ResponseArray {

	sSlice := make([]ResponsePartyReferenceDataDirectoryModel, 0, len(customers))
	for _, x := range customers {
		sSlice = append(sSlice, GenerateResponseCustomerModel(x))
	}

	return ResponseArray{ResponsePartyReferenceDataDirectory: sSlice}
}

func GenerateRequestModel(request RequestPartyReferenceDataDirectory) model.Customer {
	return model.Customer{
		Nombre:    request.Customer.CustomerName,
		Apellido:  request.Customer.CustomerLastName,
		Direccion: request.Address.ResidentialAddress,
		Correo:    request.ElectronicAddress.EmailAddress,
	}
}

// Customer configurations

type QueryConfig struct {
	PageSize int
	Page     int
}
