package handler

import (
	"app/db"
	"app/pkg/log"
	"app/pkg/util"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type CRUDHandler interface {
	GetList(c *gin.Context)
	Get(c *gin.Context)
	Insert(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
}

type SystemHandler struct {
	ds *db.DataStore
}

func NewSystemHandler(ds *db.DataStore) SystemHandler {
	return SystemHandler{
		ds: ds,
	}
}

func (h *SystemHandler) GetSystemHealth(c *gin.Context) {
	// Variable
	var redis bool = true
	var database_primary bool = true
	var database_secondary bool = true
	var hostname string = "Unknown"

	redis = h.ds.Redis != nil && func() bool {
		_, err := h.ds.Redis.Ping().Result()
		return err == nil
	}()

	database_secondary = h.ds.Db != nil && func() bool {
		sqlDB, err := h.ds.Db.DB()
		return err == nil && sqlDB.Ping() == nil
	}()

	database_primary = h.ds.Db != nil && func() bool {
		sqlDB, err := h.ds.Db.DB()
		return err == nil && sqlDB.Ping() == nil
	}()

	hostname = func() string {
		hostname, err := os.Hostname()
		if err != nil {
			return "Unknown"
		}
		return hostname
	}()

	log.Debug(util.GetTransactionID(c), "Debug System", nil)
	result := map[string]interface{}{
		"redis":              redis,
		"database_primary":   database_primary,
		"database_secondary": database_secondary,
		"app":                true,
		"hostname":           hostname,
	}

	c.JSON(http.StatusOK, result)
}