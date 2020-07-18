package dao

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type RelationDelta struct {
	gorm.Model
	Description  *string `gorm:"type:varchar(256);not null"`
	IsAddition   bool    `gorm:"default:true"`
	Type         uint    `gorm:"not null"`
	WeenieX      uint    `gorm:"not null"`
	WeenieY      uint    `gorm:"not null"`
	RelationType RelationType
	Weenie       Weenie
	conn         *gorm.DB `gorm:"-"`
}

func NewRelationDelta(db *gorm.DB, description *string, isAddition bool, relationType, weenieX, weenieY uint) *RelationDelta {
	return &RelationDelta{
		Description: description,
		IsAddition:  isAddition,
		Type:        relationType,
		WeenieX:     weenieX,
		WeenieY:     weenieY,
		conn:        db,
	}
}

func (rd *RelationDelta) TableName() string {
	return "relation_deltas"
}

func (r *RelationDelta) Migrate() error {
	conn := r.conn.AutoMigrate(r)

	conn = conn.Model(r).AddForeignKey("type", "relation_types(id)", "RESTRICT", "CASCADE")

	conn = conn.Model(r).AddForeignKey("weenie_x", "weenies(id)", "RESTRICT", "CASCADE")
	conn = conn.Model(r).AddForeignKey("weenie_y", "weenies(id)", "RESTRICT", "CASCADE")

	errors := conn.GetErrors()
	return mergeErrors(errors)
}

func (r *RelationDelta) Load() error {
	return load(r.conn, r, r.ID)
}

func (r *RelationDelta) Insert() error {
	if err := r.validateWeeniesAreDifferent(); err != nil {
		return fmt.Errorf("Error inserting relation: %s", err.Error())
	}

	return insert(r.conn, r)
}

func (r *RelationDelta) Update() error {
	if err := r.validateWeeniesAreDifferent(); err != nil {
		return fmt.Errorf("Error updating relation: %s", err.Error())
	}

	return update(r.conn, r)
}

func (r *RelationDelta) Delete() error {
	if r.ID == 0 {
		return fmt.Errorf("Missing primary key 'id' when deleting a relation model")
	}

	return delete(r.conn, r)
}

func (r *RelationDelta) validateWeeniesAreDifferent() error {
	if r.WeenieX == r.WeenieY {
		return fmt.Errorf("A relation between one weenie and itself isn't allowed")
	}

	return nil
}
