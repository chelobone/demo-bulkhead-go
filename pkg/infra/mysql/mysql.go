package mysql

import (
	"context"
	"log"

	"github.com/chelobone/demo_bulkhead_go/pkg/config"
	"github.com/chelobone/demo_bulkhead_go/pkg/domain/model"
	"github.com/chelobone/demo_bulkhead_go/pkg/repository"
	"gorm.io/gorm"
)

type clientRepository struct {
	DB     *gorm.DB
	config *config.ClientConfig
}

func NewClientRepository(db *gorm.DB, conf *config.ClientConfig) repository.IClientRepository {
	return &clientRepository{
		DB:     db,
		config: conf,
	}
}

func (ur *clientRepository) SelectClientById(ctx context.Context, clientID int) (*model.Client, error) {
	var client *model.Client
	log.Println(clientID)
	log.Println(ur.config.GetById)
	result := ur.DB.Raw(ur.config.GetById, clientID).Scan(&client)

	return client, result.Error
}

func (ur *clientRepository) SelectClients(ctx context.Context, config model.QueryConfig) (model.ClientSlice, error) {

	var clients model.ClientSlice
	result := ur.DB.Raw(ur.config.GetAll, config.Page, config.PageSize).Scan(&clients)

	return clients, result.Error
}

func (ur *clientRepository) InsertClient(ctx context.Context, config model.Client) (int, error) {

	var resultId int
	var logging string
	result := ur.DB.Raw(ur.config.PostClient, config.Nombre, config.Apellido, config.Direcion, config.Correo, logging).Scan(&resultId)
	log.Println(logging)
	return resultId, result.Error
}
