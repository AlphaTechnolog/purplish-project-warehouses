package main

import (
	"net/http"

	"github.com/alphatechnolog/purplish-warehouses/internal/di"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := di.MustOpenDB("sqlite3", "./database.db")
	defer db.Close()

	router := gin.Default()
	defer router.Run(":8001")

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	warehouseGroup := router.Group("/warehouses")

	warehouseInjector := di.NewWarehouseInjector(db)
	warehouseInjector.Inject(warehouseGroup)
}
