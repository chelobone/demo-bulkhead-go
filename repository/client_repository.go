package repository

import (
	"context"

	"github.com/chelobone/demo_bulkhead_go/domain/model"
)

type IClientRepository interface {
	SelectClientById(ctx context.Context, customerReference int) (model.Customer, error)
	SelectClients(ctx context.Context, config model.QueryConfig) (model.CustomerSlice, error)
	InsertClient(ctx context.Context, config model.Customer) (int, error)
}
