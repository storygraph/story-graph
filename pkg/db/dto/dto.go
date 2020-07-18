package dto

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type DataModel interface {
	Migrate(db *gorm.DB) []error
	GetPrimaryKeyName() string
	GetPrimaryKeyVal() interface{}
}

type DTO struct {
	conn  *gorm.DB
	model DataModel
}

func NewDTO(dbConn *gorm.DB, dataModel DataModel) *DTO {
	dto := &DTO{
		model: dataModel,
		conn:  dbConn,
	}

	return dto
}

func (dto *DTO) Load() error {
	query := fmt.Sprintf("%s = ?", dto.model.GetPrimaryKeyName())
	if errors := dto.conn.Where(query, dto.model.GetPrimaryKeyVal()).First(dto.model).GetErrors(); len(errors) > 0 {
		return mergeErrors(errors, ", ")
	}

	return nil
}

func (dto *DTO) Insert() error {
	if errors := dto.conn.Create(dto.model).GetErrors(); len(errors) > 0 {
		return mergeErrors(errors, ", ")
	}

	return nil
}

func (dto *DTO) Update() error {
	if errors := dto.conn.Save(dto.model).GetErrors(); len(errors) > 0 {
		return mergeErrors(errors, ", ")
	}

	return nil
}

func (dto *DTO) Delete() (err error) {
	// Prevents a deletion of all the rows.
	if dto.Count() == 0 {
		return fmt.Errorf(
			"Primary key %s is required for deletion",
			dto.model.GetPrimaryKeyName(),
		)
	}

	dto.conn.Delete(dto.model)

	return nil
}

func (dto *DTO) Count() (count int64) {
	query := fmt.Sprintf("%s = ?", dto.model.GetPrimaryKeyName())
	dto.conn.Model(dto.model).Where(query).Count(&count)

	return count
}

func mergeErrors(errors []error, delimiter string) error {
	mergedErrors := ""

	for i := 0; i < len(errors)-1; i++ {
		mergedErrors += errors[i].Error() + delimiter
	}
	mergedErrors += errors[len(errors)-1].Error()

	return fmt.Errorf(mergedErrors)
}
