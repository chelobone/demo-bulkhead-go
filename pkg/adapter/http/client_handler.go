package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/chelobone/demo_bulkhead_go/pkg/usecase"
	"github.com/chelobone/demo_bulkhead_go/pkg/usecase/model"
	"github.com/labstack/echo/v4"
)

type clientHandler struct {
	usecase usecase.IClientUsecase
}

func NewClientHandler(su usecase.IClientUsecase) *clientHandler {
	return &clientHandler{
		usecase: su,
	}
}

func (sh *clientHandler) FindAllClients() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		pageSize, err := strconv.Atoi(c.QueryParam("pageSize"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		page, err := strconv.Atoi(c.QueryParam("page"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		config := &model.QueryConfig{Page: page, PageSize: pageSize}
		clients, err := sh.usecase.FindClients(ctx, config)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, clients)
	}
}

func (ch *clientHandler) FindClientById() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		clientID, err := strconv.Atoi(c.Param("clientID"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		client, err := ch.usecase.FindClientById(ctx, clientID)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, client)
	}
}

func (ch *clientHandler) CreateClient() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		dec := json.NewDecoder(c.Request().Body)
		var clientInfo *model.Client = &model.Client{}

		err := dec.Decode(clientInfo)

		if err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusBadRequest, "La estructura del body no es correcta.")
		}

		client, err := ch.usecase.CreateClient(ctx, clientInfo)

		if err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, "Ocurrió un error al ejecutar el procedimiento, favor revise los logs")
		}

		return c.JSON(http.StatusOK, client)
	}
}
