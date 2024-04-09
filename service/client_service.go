package service

import (
	"context"

	"github.com/chelobone/demo_bulkhead_go/domain/model"
	"github.com/chelobone/demo_bulkhead_go/repository"
)

type IClientService interface {
	FindClientById(ctx context.Context, customerReference int) (model.Customer, error)
	FindClients(ctx context.Context, config model.QueryConfig) (model.CustomerSlice, error)
	CreateClient(ctx context.Context, config model.Customer) (int, error)
}

func NewClientService(sr repository.IClientRepository) IClientService {
	return &clientService{
		repo: sr,
	}
}

type clientService struct {
	repo repository.IClientRepository
}

func (cs *clientService) FindClientById(ctx context.Context, customerReference int) (model.Customer, error) {
	return cs.repo.SelectClientById(ctx, customerReference)
}

func (cs *clientService) FindClients(ctx context.Context, config model.QueryConfig) (model.CustomerSlice, error) {
	return cs.repo.SelectClients(ctx, config)
}

func (cs *clientService) CreateClient(ctx context.Context, config model.Customer) (int, error) {
	return cs.repo.InsertClient(ctx, config)
}
