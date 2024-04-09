package route

import (
	"fmt"

	"github.com/chelobone/demo_bulkhead_go/api/controller"
	"github.com/chelobone/demo_bulkhead_go/infra"
	"github.com/chelobone/demo_bulkhead_go/infra/mysql"
	"github.com/chelobone/demo_bulkhead_go/service"
	"github.com/chelobone/demo_bulkhead_go/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	apiVersion      = "/v1"
	healthCheckRoot = "/health_chek"

	//client
	clientsAPIRoot = apiVersion + "/customer"
	clientIDParam  = "customerReference"
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
		healthCheckGroup.GET(relativePath, controller.NewHealthCheck)
	}

	mySQLConn := infra.NewMySQLConnector()
	config := infra.NewConfig()
	clientRepository := mysql.NewClientRepository(mySQLConn.Conn, config)
	clientService := service.NewClientService(clientRepository)
	clientUsecase := usecase.NewClientUsecase(clientService)

	clientGroup := e.Group(clientsAPIRoot)
	{
		handler := controller.NewClientHandler(clientUsecase)
		// v1/customer?page=1&pageSize=10
		relativePath := ""
		clientGroup.GET(relativePath, handler.FindAllClients())
		// v1/customer/{customerReference}
		relativePath = fmt.Sprintf("/:%s", clientIDParam)
		clientGroup.GET(relativePath, handler.FindClientById())

		// v1/customer/
		relativePath = ""
		clientGroup.POST(relativePath, handler.CreateClient())
	}

	return e
}
