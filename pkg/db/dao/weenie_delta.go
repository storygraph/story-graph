package dao

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type WeenieDelta struct {
	gorm.Model
	Status          *string `gorm:"type:varchar(64);not null"`
	Name            *string `gorm:"type:varchar(256);not null"`
	CoordX          uint    `gorm:"not null"`
	CoordY          uint    `gorm:"not null"`
	IsTravelling    bool    `gorm:"not null;default:false"`
	WeenieID        uint    `gorm:"not null"`
	TagDeltas       []TagDelta
	PossessionDelta []PossessionDelta
	conn            *gorm.DB `gorm:"-"`
}

func NewWeenieDelta(
	db *gorm.DB,
	status, name *string,
	coordX, coordY uint,
	isTravelling bool,
	weenieID uint,
) *WeenieDelta {
	return &WeenieDelta{
		Status:       status,
		Name:         name,
		CoordX:       coordX,
		CoordY:       coordY,
		IsTravelling: isTravelling,
		WeenieID:     weenieID,
		conn:         db,
	}
}

func (wd *WeenieDelta) TableName() string {
	return "weenie_deltas"
}

func (wd *WeenieDelta) Migrate() error {
	conn := wd.conn.AutoMigrate(wd)
	conn = conn.Model(wd).AddForeignKey("weenie_id", "weenies(id)", "RESTRICT", "CASCADE")

	errors := conn.GetErrors()
	return mergeErrors(errors)
}

func (wd *WeenieDelta) Load() error {
	return load(wd.conn, wd, wd.ID)
}

func (wd *WeenieDelta) Insert() error {
	return insert(wd.conn, wd)
}

func (wd *WeenieDelta) Update() error {
	return update(wd.conn, wd)
}

func (wd *WeenieDelta) Delete() error {
	if wd.ID == 0 {
		return fmt.Errorf("Missing primary key 'id' when deleting a weenie delta model")
	}

	return delete(wd.conn, wd)
}
