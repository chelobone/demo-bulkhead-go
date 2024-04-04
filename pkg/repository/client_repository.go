package repository

import (
	"context"

	"github.com/chelobone/demo_bulkhead_go/pkg/domain/model"
)

type IClientRepository interface {
	SelectClientById(ctx context.Context, clientID int) (*model.Client, error)
	SelectClients(ctx context.Context, config model.QueryConfig) (model.ClientSlice, error)
	InsertClient(ctx context.Context, config model.Client) (int, error)
}
