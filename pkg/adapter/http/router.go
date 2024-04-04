package http

import (
	"fmt"

	"github.com/chelobone/demo_bulkhead_go/pkg/infra"
	"github.com/chelobone/demo_bulkhead_go/pkg/infra/mysql"
	"github.com/chelobone/demo_bulkhead_go/pkg/service"
	"github.com/chelobone/demo_bulkhead_go/pkg/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	apiVersion      = "/v1"
	healthCheckRoot = "/health_chek"

	//client
	clientsAPIRoot = apiVersion + "/clients"
	clientIDParam  = "clientID"
)

func InitRouter() *echo.Echo {
	e := echo.New()
	e.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	//health chek
	healthCheckGroup := e.Group(healthCheckRoot)
	{
		relativePath := ""
		healthCheckGroup.GET(relativePath, healthCheck)
	}

	mySQLConn := infra.NewMySQLConnector()
	config := infra.NewConfig()
	clientRepository := mysql.NewClientRepository(mySQLConn.Conn, config)
	clientService := service.NewClientService(clientRepository)
	clientUsecase := usecase.NewClientUsecase(clientService)

	clientGroup := e.Group(clientsAPIRoot)
	{
		handler := NewClientHandler(clientUsecase)
		// v1/clients?page=1&pageSize=10
		relativePath := "" //fmt.Sprintf("/:%s/:%s", pageParam, pageSizeParam)
		clientGroup.GET(relativePath, handler.FindAllClients())
		// v1/clients/{client_id}
		relativePath = fmt.Sprintf("/:%s", clientIDParam)
		clientGroup.GET(relativePath, handler.FindClientById())

		// v1/students/
		relativePath = ""
		clientGroup.POST(relativePath, handler.CreateClient())
	}

	return e
}
