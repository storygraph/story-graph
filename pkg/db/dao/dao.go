package dao

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

const ERRORS_DELIMITER = ", "

type DataModel interface {
	// Migrate applies DDL changes to the db schema.
	Migrate() error
}

type DAO interface {
	DataModel
	// Load fetches the data by id.
	Load() error
	// Insert creates a new row into the db.
	Insert() error
	// Update modifies the corresponding row into the db.
	Update() error
	// Delete deletes the corresponding row from the db. The primary key is mandatory.
	Delete() error
}

func load(conn *gorm.DB, model DataModel, id uint) error {
	if id == 0 {
		return fmt.Errorf("Missing primary key 'id' when loading model")
	}

	if errors := conn.Where("id = ?", id).First(model).GetErrors(); len(errors) > 0 {
		return mergeErrors(errors)
	}

	return nil
}

func insert(conn *gorm.DB, model DataModel) error {
	if errors := conn.Create(model).GetErrors(); len(errors) > 0 {
		return mergeErrors(errors)
	}

	return nil
}

func update(conn *gorm.DB, model DataModel) error {
	if errors := conn.Save(model).GetErrors(); len(errors) > 0 {
		return mergeErrors(errors)
	}

	return nil
}

func delete(conn *gorm.DB, model DataModel) error {
	if errors := conn.Delete(model).GetErrors(); len(errors) > 0 {
		return mergeErrors(errors)
	}

	return nil
}

func mergeErrors(errors []error) error {
	if len(errors) == 0 {
		return nil
	}

	mergedErrors := ""

	for i := 0; i < len(errors)-1; i++ {
		mergedErrors += errors[i].Error() + ERRORS_DELIMITER
	}
	mergedErrors += errors[len(errors)-1].Error()

	return fmt.Errorf(mergedErrors)
}
