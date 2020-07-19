package dao

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type TagDelta struct {
	gorm.Model
	Key           *string  `gorm:"type:varchar(64);not null"`
	Description   *string  `gorm:"type:varchar(256);not null"`
	IsAddition    bool     `gorm:"default:true"`
	WeenieDeltaID uint     `gorm:"not null"`
	conn          *gorm.DB `gorm:"-"`
}

func NewTagDelta(db *gorm.DB, key, description *string, isAddition bool, weenieDeltaID uint) *TagDelta {
	return &TagDelta{
		Key:           key,
		Description:   description,
		IsAddition:    isAddition,
		WeenieDeltaID: weenieDeltaID,
		conn:          db,
	}
}

func (td *TagDelta) TableName() string {
	return "tag_deltas"
}

func (td *TagDelta) Migrate() error {
	conn := td.conn.AutoMigrate(td)
	conn = conn.Model(td).AddForeignKey("weenie_delta_id", "weenie_deltas(id)", "RESTRICT", "CASCADE")

	errors := conn.GetErrors()
	return mergeErrors(errors)
}

func (td *TagDelta) Load() error {
	return load(td.conn, td, td.ID)
}

func (td *TagDelta) Insert() error {
	return insert(td.conn, td)
}

func (td *TagDelta) Update() error {
	return update(td.conn, td)
}

func (td *TagDelta) Delete() error {
	if td.ID == 0 {
		return fmt.Errorf("Missing primary key 'id' when deleting a tag model")
	}

	return delete(td.conn, td)
}
