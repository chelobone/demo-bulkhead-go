package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/chelobone/demo_bulkhead_go/domain/model"
	"github.com/chelobone/demo_bulkhead_go/usecase"
	"github.com/chelobone/demo_bulkhead_go/usecase/dto"
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
		var resultInfo *dto.Result[dto.ResponseArray, dto.ResponseError] = new(dto.Result[dto.ResponseArray, dto.ResponseError])

		pageSize, err := strconv.Atoi(c.QueryParam("pageSize"))

		if err != nil {
			log.Println(err.Error())

			var responseError *dto.DomainError = &dto.DomainError{
				StatusCode: dto.JsonParsing,
				Message:    "El parámetro pageSize no se encuentra en la Query",
			}

			resultInfo.SetFault(dto.ResponseError{Response: responseError})
		}

		page, err := strconv.Atoi(c.QueryParam("page"))
		if err != nil {
			log.Println(err.Error())

			var responseError *dto.DomainError = &dto.DomainError{
				StatusCode: dto.JsonParsing,
				Message:    "El parámetro page no se encuentra en la Query",
			}

			resultInfo.SetFault(dto.ResponseError{Response: responseError})
		}

		config := model.QueryConfig{Page: page, PageSize: pageSize}

		if resultInfo.State() == dto.Success {
			client, err := sh.usecase.FindClients(ctx, config)
			resultInfo.SetValue(client.Value())
			log.Println(err)
		}

		if resultInfo.State() == dto.Success {
			log.Println(resultInfo.Value())
			return c.JSON(http.StatusOK, resultInfo.Value())
		} else {
			return c.JSON(http.StatusBadRequest, resultInfo.Fault())
		}
	}
}

func (ch *clientHandler) FindClientById() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		clientID, err := strconv.Atoi(c.Param("customerReference"))

		var resultInfo *dto.Result[dto.ResponseModel, dto.ResponseError] = new(dto.Result[dto.ResponseModel, dto.ResponseError])
		if err != nil {
			log.Println(err.Error())

			var responseError *dto.DomainError = &dto.DomainError{
				StatusCode: dto.JsonParsing,
				Message:    "La estructura del json es incorrecta",
			}

			resultInfo.SetFault(dto.ResponseError{Response: responseError})
		}

		if resultInfo.State() == dto.Success {
			resultInfo, err = ch.usecase.FindClientById(ctx, clientID)

			log.Println(resultInfo)
			log.Println(err)
		}

		if resultInfo.State() == dto.Success {
			return c.JSON(http.StatusOK, resultInfo.Value())
		} else {
			return c.JSON(http.StatusBadRequest, resultInfo.Fault())
		}
	}
}

func (ch *clientHandler) CreateClient() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		dec := json.NewDecoder(c.Request().Body)
		var clientInfo *dto.Request = &dto.Request{}

		var resultInfo *dto.Result[dto.Response, dto.ResponseError] = new(dto.Result[dto.Response, dto.ResponseError])
		err := dec.Decode(clientInfo)

		if err != nil {
			log.Println(err.Error())

			var responseError *dto.DomainError = &dto.DomainError{
				StatusCode: dto.JsonParsing,
				Message:    "La estructura del json es incorrecta",
			}

			resultInfo.SetFault(dto.ResponseError{Response: responseError})
		}

		if resultInfo.State() == dto.Success {
			client, err := ch.usecase.CreateClient(ctx, clientInfo.RequestPartyReferenceDataDirectory)
			resultInfo = client
			log.Println(client)
			log.Println(err)
		}

		if resultInfo.State() == dto.Success {
			return c.JSON(http.StatusOK, resultInfo.Value())
		} else {
			return c.JSON(http.StatusBadRequest, resultInfo.Fault())
		}
	}
}
