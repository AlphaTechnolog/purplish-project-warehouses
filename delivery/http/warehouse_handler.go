package http

import (
	"net/http"

	"github.com/alphatechnolog/purplish-warehouses/internal/domain"
	"github.com/alphatechnolog/purplish-warehouses/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type WarehouseHandler struct {
	warehouseUsecase *usecase.WarehouseUsecase
}

func NewWarehouseHandler(warehouseUsecase *usecase.WarehouseUsecase) *WarehouseHandler {
	return &WarehouseHandler{warehouseUsecase}
}

func (h *WarehouseHandler) GetWarehouses(c *gin.Context) {
	warehouses, err := h.warehouseUsecase.GetWarehouses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"warehouses": warehouses})
}

func (h *WarehouseHandler) GetWarehouse(c *gin.Context) {
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	warehouse, err := h.warehouseUsecase.GetWarehouse(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"warehouse": warehouse})
}

func (h *WarehouseHandler) CreateWarehouse(c *gin.Context) {
	var warehouse domain.Warehouse
	if err := c.ShouldBindJSON(&warehouse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.warehouseUsecase.CreateWarehouse(&warehouse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create warehouse"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"warehouse": warehouse})
}

func (h *WarehouseHandler) UpdateWarehouse(c *gin.Context) {
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var warehouse domain.Warehouse
	if err := c.ShouldBindJSON(&warehouse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	warehouse.ID = id

	if err := h.warehouseUsecase.UpdateWarehouse(&warehouse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update warehouse"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"warehouse": warehouse})
}

func (h *WarehouseHandler) DeleteWarehouse(c *gin.Context) {
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.warehouseUsecase.DeleteWarehouse(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete warehouse"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": true})
}
