package core

import (
	"database/sql"
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

func CreateWarehousesRoutes(d *sql.DB, r *gin.RouterGroup) {
	r.GET("/", WrapError(WithDB(d, getWarehouses)))
}
