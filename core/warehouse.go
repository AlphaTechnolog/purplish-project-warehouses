package core

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"

	"github.com/alphatechnolog/purplish-warehouses/database"
	"github.com/gin-gonic/gin"
)

func getWarehouses(d *sql.DB, c *gin.Context) error {
	warehouses, err := database.GetWarehouses(d)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, gin.H{
		"warehouses": warehouses,
	})

	return nil
}

func createWarehouse(d *sql.DB, c *gin.Context) error {
	bodyContents, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}

	var createPayload database.CreateWarehousePayload
	if err = json.Unmarshal(bodyContents, &createPayload); err != nil {
		return err
	}

	if err = database.CreateWarehouse(d, createPayload); err != nil {
		return err
	}

	c.JSON(http.StatusCreated, gin.H{"ok": true})

	return nil
}

func removeWarehouse(d *sql.DB, c *gin.Context) error {
	warehouseID := c.Param("ID")
	if warehouseID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return nil
	}

	if err := database.RemoveWarehouse(d, warehouseID); err != nil {
		return err
	}

	c.JSON(http.StatusOK, gin.H{"ok": true})

	return nil
}

func CreateWarehousesRoutes(d *sql.DB, r *gin.RouterGroup) {
	r.GET("/", WrapError(WithDB(d, getWarehouses)))
	r.POST("/", WrapError(WithDB(d, createWarehouse)))
	r.DELETE("/:ID", WrapError(WithDB(d, removeWarehouse)))
}
