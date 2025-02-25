package database

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Warehouse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

type CreateWarehousePayload struct {
	Name string `json:"name"`
}

func GetWarehouses(d *sql.DB) ([]Warehouse, error) {
	var warehouses []Warehouse

	rows, err := d.Query("SELECT * FROM warehouses WHERE status = 1;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var warehouse Warehouse
		err = rows.Scan(&warehouse.ID, &warehouse.Name, &warehouse.Status)
		if err != nil {
			return nil, err
		}

		warehouses = append(warehouses, warehouse)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return warehouses, nil
}

func CreateWarehouse(d *sql.DB, createPayload CreateWarehousePayload) error {
	sql := `
		INSERT INTO warehouses (id, name, status)
		VALUES
			(?, ?, 1);
	`

	_, err := d.Exec(sql, uuid.New().String(), createPayload.Name)
	if err != nil {
		return fmt.Errorf("Unable to create warehouse: %w", err)
	}

	return nil
}

func RemoveWarehouse(d *sql.DB, warehouseID string) error {
	sql := `
		DELETE FROM warehouses WHERE id = ?;
	`

	if _, err := d.Exec(sql, warehouseID); err != nil {
		return fmt.Errorf("Unable to remove warehouse by id '%s': %w", warehouseID, err)
	}

	return nil
}
