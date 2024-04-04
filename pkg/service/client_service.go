package service

import (
	"context"

	"github.com/chelobone/demo_bulkhead_go/pkg/domain/model"
	"github.com/chelobone/demo_bulkhead_go/pkg/repository"
)

type IClientService interface {
	FindClientById(ctx context.Context, clientID int) (*model.Client, error)
	FindClients(ctx context.Context, config model.QueryConfig) (model.ClientSlice, error)
	CreateClient(ctx context.Context, config model.Client) (int, error)
}

func NewClientService(sr repository.IClientRepository) IClientService {
	return &clientService{
		repo: sr,
	}
}

type clientService struct {
	repo repository.IClientRepository
}

func (cs *clientService) FindClientById(ctx context.Context, clientID int) (*model.Client, error) {
	return cs.repo.SelectClientById(ctx, clientID)
}

func (cs *clientService) FindClients(ctx context.Context, config model.QueryConfig) (model.ClientSlice, error) {
	return cs.repo.SelectClients(ctx, config)
}

func (cs *clientService) CreateClient(ctx context.Context, config model.Client) (int, error) {
	return cs.repo.InsertClient(ctx, config)
}
