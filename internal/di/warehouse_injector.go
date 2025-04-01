package di

import (
	"database/sql"

	"github.com/alphatechnolog/purplish-warehouses/delivery/http"
	"github.com/alphatechnolog/purplish-warehouses/internal/repository"
	"github.com/alphatechnolog/purplish-warehouses/internal/usecase"
	"github.com/gin-gonic/gin"
)

type WarehouseInjector struct {
	db *sql.DB
}

func NewWarehouseInjector(db *sql.DB) ModuleInjector {
	return &WarehouseInjector{db: db}
}

func (wi *WarehouseInjector) Inject(routerGroup *gin.RouterGroup) {
	sqliteRepo := repository.NewSQLiteRepository(wi.db)
	warehouseUseCase := usecase.NewWarehouseUsecase(sqliteRepo)
	warehouseHandler := http.NewWarehouseHandler(warehouseUseCase)

	routerGroup.GET("", warehouseHandler.GetWarehouses)
	routerGroup.GET("/:id", warehouseHandler.GetWarehouse)
	routerGroup.POST("", warehouseHandler.CreateWarehouse)
	routerGroup.PUT("/:id", warehouseHandler.UpdateWarehouse)
	routerGroup.DELETE("/:id", warehouseHandler.DeleteWarehouse)
}
