package handlers

import (
	"RESTful/dto"
	"RESTful/service"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	Service service.AccountService
}

// NewAccount parse the account request from user side
func (h AccountHandler) NewAccount(ctx *gin.Context) {
	var request dto.NewAccountRequest
	err := json.NewDecoder(ctx.Request.Body).Decode(&request)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	request.CustomerId = ctx.Param("id")
	account, appErr := h.Service.NewAccount(request)
	if appErr != nil {
		ctx.JSON(appErr.Code, appErr.Message)
		return
	}

	ctx.JSON(http.StatusCreated, account)
}