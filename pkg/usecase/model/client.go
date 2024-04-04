package model

import (
	"github.com/chelobone/demo_bulkhead_go/pkg/domain/model"
)

type Client struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Direcion string `json:"direccion"`
	Correo   string `json:"correo"`
}

type ClientSlice []*Client

type QueryConfig struct {
	PageSize int
	Page     int
}

func ConfigFromUseCaseModel(m *QueryConfig) model.QueryConfig {
	s := model.QueryConfig{
		Page:     m.Page,
		PageSize: m.PageSize,
	}

	return s
}

func ClientFromUseCaseModel(m *Client) model.Client {
	s := model.Client{
		ID:       m.ID,
		Nombre:   m.Nombre,
		Apellido: m.Apellido,
	}

	return s
}

func ClientFromDomainModel(m *model.Client) *Client {
	s := &Client{
		ID:       m.ID,
		Nombre:   m.Nombre,
		Apellido: m.Apellido,
		Direcion: m.Direcion,
		Correo:   m.Correo,
	}

	return s
}
