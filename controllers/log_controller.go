package controllers

import (
	"github.com/anggacipta/order-management-api/dto"
	"github.com/anggacipta/order-management-api/helpers"
	"github.com/anggacipta/order-management-api/services"
	"github.com/gin-gonic/gin"
)

type LogController struct {
	logService services.LogService
}

func NewLogController(logService services.LogService) *LogController {
	return &LogController{logService: logService}
}

func (ctrl *LogController) GetAllLogs(c *gin.Context) {
	logs, err := ctrl.logService.GetAllLogs()
	if err != nil {
		helpers.RespondInternalError(c, err)
		return
	}
	c.JSON(200, logs)
}

func (ctrl *LogController) CreateLog(c *gin.Context) {
	var input dto.LogRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.RespondValidationError(c, err)
		return
	}

	if err := ctrl.logService.CreateLog(input); err != nil {
		helpers.RespondInternalError(c, err)
		return
	}

	c.JSON(200, gin.H{"message": "Log created successfully"})
}
