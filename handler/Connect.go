package handler

import (
	"app/constanta"
	"app/db"
	"app/model"
	"app/pkg/util"
	"app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ConnectHandler struct {
	CRUDHandler
	ConnectService service.ConnectService
}

func NewConnectHandler(db *db.DataStore) ConnectHandler {
	svc := service.NewConnectService(db)
	h := ConnectHandler{
		ConnectService: svc,
	}
	return h
}



func (h *ConnectHandler) GetList(c *gin.Context) {
	// Call Service
	result, err := h.ConnectService.GetList()

	// Construct Response
	if err != nil {
		c.Errors = append(c.Errors, err.GenerateReponse(util.GetTransactionID(c)))
		c.JSON(err.Status, err.Response)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func (h *ConnectHandler) Insert(c *gin.Context) {
	// Cast data from request
	var data model.ConnectInput
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Construct Connect Model with the request data
	newConnect := &model.Connect{}
	newConnect.PopulateFromDTOInput(data)
	// Call create service
	err := h.ConnectService.Insert(newConnect)

	// Construct Response
	if err != nil {
		err.GenerateReponse(util.GetTransactionID(c))
		c.JSON(err.Status, err.Response)
	} else {
		c.JSON(http.StatusCreated, constanta.SuccessMessage)
	}
}