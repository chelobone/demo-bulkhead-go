package dto

type DomainErrorType string

const (
	NetworkRequest DomainErrorType = "001"
	JsonParsing    DomainErrorType = "002"
	Storage        DomainErrorType = "003"
)

type DomainError struct {
	StatusCode DomainErrorType
	Message    string
}

type ResponseError struct {
	Response *DomainError
}
