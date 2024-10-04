package apiHelper

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/speedianet/os/src/presentation/service"
)

type newFormattedResponse struct {
	Status int         `json:"status"`
	Body   interface{} `json:"body"`
}

func ServiceResponseWrapper(
	c echo.Context,
	serviceOutput service.ServiceOutput,
) error {
	responseStatus := http.StatusOK
	switch serviceOutput.Status {
	case service.Created:
		responseStatus = http.StatusCreated
	case service.MultiStatus:
		responseStatus = http.StatusMultiStatus
	case service.UserError:
		responseStatus = http.StatusBadRequest
	case service.InfraError:
		responseStatus = http.StatusInternalServerError
	}

	formattedResponse := newFormattedResponse{
		Status: responseStatus,
		Body:   serviceOutput.Body,
	}
	return c.JSON(responseStatus, formattedResponse)
}

func ServiceTokenResponseWrapper(
	c echo.Context,
	serviceOutput service.ServiceOutput,
) error {
	c.Response().Header().Set("X-Is-Token-Response", "true")

	return ServiceResponseWrapper(c, serviceOutput)
}
