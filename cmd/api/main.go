package main

import (
	"log"
	"net/http"

	"github.com/alphatechnolog/purplish-warehouses/internal/config"
	"github.com/alphatechnolog/purplish-warehouses/internal/di"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

const ENV_FILE = ".env"

func main() {
	cfg, err := config.LoadConfig(ENV_FILE)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
		panic(err)
	}

	db := di.MustOpenDB("sqlite3", cfg.DatabaseURL)
	defer db.Close()

	router := gin.Default()
	defer router.Run(":" + cfg.ServerPort)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	warehouseGroup := router.Group("/warehouses")

	warehouseInjector := di.NewWarehouseInjector(db)
	warehouseInjector.Inject(warehouseGroup)
}
