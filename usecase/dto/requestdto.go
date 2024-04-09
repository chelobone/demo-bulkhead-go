package dto

// Request dto
type Request struct {
	RequestPartyReferenceDataDirectory RequestPartyReferenceDataDirectory
}

type RequestPartyReferenceDataDirectory struct {
	Customer          RequestCustomer
	Address           RequestAddress
	ElectronicAddress RequestElectronicAddress
}

type RequestCustomer struct {
	CustomerName     string `json:"customerName"`
	CustomerLastName string `json:"customerLastName"`
}

type RequestAddress struct {
	ResidentialAddress string `json:"residentialAddress"`
}

type RequestElectronicAddress struct {
	EmailAddress string `json:"emailAddress"`
}
