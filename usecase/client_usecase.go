package usecase

import (
	"context"

	"github.com/chelobone/demo_bulkhead_go/domain/model"
	"github.com/chelobone/demo_bulkhead_go/service"
	"github.com/chelobone/demo_bulkhead_go/usecase/dto"
)

type IClientUsecase interface {
	FindClientById(ctx context.Context, clientID int) (*dto.Result[dto.ResponseModel, dto.ResponseError], error)
	FindClients(ctx context.Context, config model.QueryConfig) (*dto.Result[dto.ResponseArray, dto.ResponseError], error)
	CreateClient(ctx context.Context, config dto.RequestPartyReferenceDataDirectory) (*dto.Result[dto.Response, dto.ResponseError], error)
}

func NewClientUsecase(ss service.IClientService) IClientUsecase {
	return &clientUsecase{
		cs: ss,
	}
}

type clientUsecase struct {
	cs service.IClientService
}

func (cu *clientUsecase) FindClientById(ctx context.Context, customerReference int) (*dto.Result[dto.ResponseModel, dto.ResponseError], error) {
	result, err := cu.cs.FindClientById(ctx, customerReference)

	resultPattern := new(dto.Result[dto.ResponseModel, dto.ResponseError])
	resultPattern.SetValue(dto.ResponseModel{
		ResponsePartyReferenceDataDirectory: dto.GenerateResponseCustomerModel(result),
	})
	return resultPattern, err
}

func (cu *clientUsecase) FindClients(ctx context.Context, config model.QueryConfig) (*dto.Result[dto.ResponseArray, dto.ResponseError], error) {
	result, err := cu.cs.FindClients(ctx, config)
	resultPattern := new(dto.Result[dto.ResponseArray, dto.ResponseError])
	resultPattern.SetValue(dto.GenerateResponseCustomerModelArray(result))
	return resultPattern, err
}

func (cu *clientUsecase) CreateClient(ctx context.Context, client dto.RequestPartyReferenceDataDirectory) (*dto.Result[dto.Response, dto.ResponseError], error) {

	// Transformar dto a modelo de dominio de datos de la tabla
	clientModel := dto.GenerateRequestModel(client)

	result, err := cu.cs.CreateClient(ctx, clientModel)

	//Definir objeto para patrón de resultado
	resultPattern := new(dto.Result[dto.Response, dto.ResponseError])

	if err != nil {
		var responseError *dto.DomainError = &dto.DomainError{
			StatusCode: dto.Storage,
			Message:    "Error al ejecutar procedimiento almacenado",
		}

		resultPattern.SetFault(dto.ResponseError{Response: responseError})
	} else {
		// Transformar dominio de datos de la tabla a dto
		modelRequest := dto.GenerateResponseModel(result)

		resultPattern.SetValue(modelRequest)
	}
	return resultPattern, err
}
