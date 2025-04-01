package usecase

import (
	"fmt"
	"time"

	"github.com/alphatechnolog/purplish-warehouses/internal/domain"
	"github.com/alphatechnolog/purplish-warehouses/internal/repository"
	"github.com/google/uuid"
)

type WarehouseUsecase struct {
	sqldbRepo repository.SQLDBRepository
}

func NewWarehouseUsecase(sqldbRepo repository.SQLDBRepository) *WarehouseUsecase {
	return &WarehouseUsecase{
		sqldbRepo,
	}
}

func (uc *WarehouseUsecase) GetWarehouses() ([]domain.Warehouse, error) {
	query := "SELECT id, name, status, created_at, updated_at FROM warehouses"
	warehouses := []domain.Warehouse{}

	rows, err := uc.sqldbRepo.Query(query)
	if err != nil {
		return warehouses, fmt.Errorf("unable to query warehouses: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var warehouse domain.Warehouse
		err = rows.Scan(&warehouse.ID, &warehouse.Name, &warehouse.Status, &warehouse.CreatedAt, &warehouse.UpdatedAt)
		if err != nil {
			return warehouses, fmt.Errorf("cannot scan queryset: %w", err)
		}
		warehouses = append(warehouses, warehouse)
	}

	return warehouses, nil
}

func (uc *WarehouseUsecase) GetWarehouse(id string) (*domain.Warehouse, error) {
	query := "SELECT id, name, status, created_at, updated_at FROM warehouses WHERE id = ?"
	row := uc.sqldbRepo.QueryRow(query, id)

	warehouse := &domain.Warehouse{}
	err := row.Scan(&warehouse.ID, &warehouse.Name, &warehouse.Status, &warehouse.CreatedAt, &warehouse.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to scan warehouse: %w", err)
	}

	return warehouse, nil
}

func (uc *WarehouseUsecase) CreateWarehouse(warehouse *domain.Warehouse) error {
	query := "INSERT INTO warehouses (id, name, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	now := time.Now()
	warehouse.ID = uuid.New().String()
	warehouse.CreatedAt = now
	warehouse.UpdatedAt = now

	_, err := uc.sqldbRepo.Execute(query, warehouse.ID, warehouse.Name, warehouse.Status, warehouse.CreatedAt, warehouse.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to create warehouse: %w", err)
	}

	return nil
}

func (uc *WarehouseUsecase) UpdateWarehouse(warehouse *domain.Warehouse) error {
	query := "UPDATE warehouses SET name = ?, status = ?, updated_at = ? WHERE id = ?"
	warehouse.UpdatedAt = time.Now()

	_, err := uc.sqldbRepo.Execute(query, warehouse.Name, warehouse.Status, warehouse.UpdatedAt, warehouse.ID)
	if err != nil {
		return fmt.Errorf("failed to update warehouse: %w", err)
	}

	return nil
}

func (uc *WarehouseUsecase) DeleteWarehouse(id string) error {
	query := "DELETE FROM warehouses WHERE id = ?"
	_, err := uc.sqldbRepo.Execute(query, id)
	if err != nil {
		return fmt.Errorf("failed to remove warehouse by id '%s': %w", id, err)
	}

	return nil
}
