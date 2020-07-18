package dao

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type PossessionDelta struct {
	gorm.Model
	Name          *string  `gorm:"type:varchar(128);not null"`
	Description   *string  `gorm:"type:varchar(256);not null"`
	IsAddition    bool     `gorm:"default:true"`
	WeenieDeltaID uint     `gorm:"not null"`
	conn          *gorm.DB `gorm:"-"`
}

func NewPossessionDelta(db *gorm.DB, name, description *string, isAddition bool, weenieDeltaID uint) *PossessionDelta {
	return &PossessionDelta{
		Name:          name,
		Description:   description,
		IsAddition:    isAddition,
		WeenieDeltaID: weenieDeltaID,
		conn:          db,
	}
}

func (pd *PossessionDelta) TableName() string {
	return "possession_deltas"
}

func (pd *PossessionDelta) Migrate() error {
	conn := pd.conn.AutoMigrate(pd)
	conn = conn.Model(pd).AddForeignKey("weenie_delta_id", "weenie_deltas(id)", "RESTRICT", "CASCADE")

	errors := conn.GetErrors()
	return mergeErrors(errors)
}

func (pd *PossessionDelta) Load() error {
	return load(pd.conn, pd, pd.ID)
}

func (pd *PossessionDelta) Insert() error {
	return insert(pd.conn, pd)
}

func (pd *PossessionDelta) Update() error {
	return update(pd.conn, pd)
}

func (pd *PossessionDelta) Delete() error {
	if pd.ID == 0 {
		return fmt.Errorf("Missing primary key 'id' when deleting a possession model")
	}

	return delete(pd.conn, pd)
}
