package usecase

import (
	"context"

	"github.com/chelobone/demo_bulkhead_go/pkg/service"
	"github.com/chelobone/demo_bulkhead_go/pkg/usecase/model"
)

type IClientUsecase interface {
	FindClientById(ctx context.Context, clientID int) (*model.Client, error)
	FindClients(ctx context.Context, config *model.QueryConfig) (model.ClientSlice, error)
	CreateClient(ctx context.Context, config *model.Client) (int, error)
}

func NewClientUsecase(ss service.IClientService) IClientUsecase {
	return &clientUsecase{
		cs: ss,
	}
}

type clientUsecase struct {
	cs service.IClientService
}

func (cu *clientUsecase) FindClientById(ctx context.Context, clientID int) (*model.Client, error) {
	result, err := cu.cs.FindClientById(ctx, clientID)

	return model.ClientFromDomainModel(result), err
}

func (cu *clientUsecase) FindClients(ctx context.Context, config *model.QueryConfig) (model.ClientSlice, error) {
	result, err := cu.cs.FindClients(ctx, model.ConfigFromUseCaseModel(config))

	sSlice := make(model.ClientSlice, 0, len(result))
	for _, ms := range result {
		sSlice = append(sSlice, model.ClientFromDomainModel(ms))
	}

	return sSlice, err
}

func (cu *clientUsecase) CreateClient(ctx context.Context, client *model.Client) (int, error) {
	result, err := cu.cs.CreateClient(ctx, model.ClientFromUseCaseModel(client))

	return result, err
}
