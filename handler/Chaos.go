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

type ChaosHandler struct {
	CRUDHandler
	ChaosService service.ChaosService
}

func NewChaosHandler(db *db.DataStore) ChaosHandler {
	svc := service.NewChaosService(db)
	h := ChaosHandler{
		ChaosService: svc,
	}
	return h
}



func (h *ChaosHandler) Attack(c *gin.Context) {
	// Cast data from request
	var data model.ChaosInput
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Construct Chaos Model with the request data
	newChaos := &model.Chaos{}
	newChaos.PopulateFromDTOInput(data)
	// Call create service
	err := h.ChaosService.Attack(newChaos)

	// Construct Response
	if err != nil {
		err.GenerateReponse(util.GetTransactionID(c))
		c.JSON(err.Status, err.Response)
	} else {
		c.JSON(http.StatusCreated, constanta.SuccessMessage)
	}
}