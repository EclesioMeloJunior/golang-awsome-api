package handlers

import (
	"go-challenge/internals/models"
	"go-challenge/internals/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Healthcheck handlers define the endpoint controllers
// to access the API status
type Healthcheck struct {
	hcService services.Healthcheck
}

// NewHealthcheckHandler injects the healthcheck service
// into handler
func NewHealthcheckHandler(hcService services.Healthcheck) *Healthcheck {
	return &Healthcheck{hcService}
}

// GetAPIStatus returns the status of mongodb connection
// when the last sync occours and the system info
func (h *Healthcheck) GetAPIStatus(e echo.Context) error {
	var err error
	status := new(models.Status)

	status.Database.Description = "All database checks are done"
	status.Database.Status = "OK"
	status.MemUsage = h.hcService.GetMemUsage()
	status.OnlineT = h.hcService.OnlineSince().String()

	if status.LastSync, err = h.hcService.LastSyncExecution(); err != nil {
		return e.String(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, status)
}
