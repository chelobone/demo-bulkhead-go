package mysql

import (
	"context"
	"log"
	"time"

	"github.com/chelobone/demo_bulkhead_go/config"
	"github.com/chelobone/demo_bulkhead_go/domain/model"
	"github.com/chelobone/demo_bulkhead_go/repository"
	"github.com/sethvargo/go-retry"
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

func (ur *clientRepository) SelectClientById(ctx context.Context, customerReference int) (model.Customer, error) {
	var client model.Customer
	log.Println(customerReference)
	log.Println(ur.config.GetById)
	result := ur.DB.Raw(ur.config.GetById, customerReference).Scan(&client)

	return client, result.Error
}

func (ur *clientRepository) SelectClients(ctx context.Context, config model.QueryConfig) (model.CustomerSlice, error) {

	var clients model.CustomerSlice
	result := ur.DB.Raw(ur.config.GetAll, config.Page, config.PageSize).Scan(&clients)

	return clients, result.Error
}

func (ur *clientRepository) InsertClient(ctx context.Context, info model.Customer) (int, error) {

	var resultId int
	var logging string

	b := retry.NewFibonacci(1 * time.Second)
	b = retry.WithMaxRetries(uint64(ur.config.MaxRetries), b)

	if err := retry.Do(ctx, b, func(ctx context.Context) error {
		result := ur.DB.Raw(ur.config.PostClient, info.Nombre, info.Apellido, info.Direccion, info.Correo, logging).Scan(&resultId)
		log.Println(logging)
		if result.Error != nil {
			log.Println("Ocurrió un error, reintentando")
			return retry.RetryableError(result.Error)
		}
		return nil
	}); err != nil {
		return -1, err
	}

	return resultId, nil
}
