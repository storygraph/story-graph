package dao

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type RelationType struct {
	gorm.Model
	Name *string  `gorm:"type:varchar(64);unique;not null"`
	conn *gorm.DB `gorm:"-"`
}

func NewRelationType(db *gorm.DB, name *string) *RelationType {
	return &RelationType{
		conn: db,
		Name: name,
	}
}

func (rt *RelationType) Migrate() error {
	errors := rt.conn.AutoMigrate(rt).GetErrors()
	return mergeErrors(errors)
}

func (rt *RelationType) Load() error {
	return load(rt.conn, rt, rt.ID)
}

func (rt *RelationType) Insert() error {
	return insert(rt.conn, rt)
}

func (rt *RelationType) Update() error {
	return update(rt.conn, rt)
}

func (rt *RelationType) Delete() error {
	if rt.ID == 0 {
		return fmt.Errorf("Missing primary key 'id' when deleting a relation type model")
	}

	return delete(rt.conn, rt)
}
