package database

import "database/sql"

type Warehouse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
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
