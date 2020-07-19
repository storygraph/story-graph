package dao

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Weenie struct {
	gorm.Model
	Status  *string  `gorm:"type:varchar(64);not null"`
	Image   *string  `gorm:"type:varchar(256);not null"`
	Name    *string  `gorm:"type:varchar(256);not null"`
	StoryID uint     `gorm:"not null"`
	conn    *gorm.DB `gorm:"-"`
}

func NewWeenie(db *gorm.DB, status, image, name *string, storyID uint) *Weenie {
	return &Weenie{
		Status:  status,
		Image:   image,
		Name:    name,
		StoryID: storyID,
		conn:    db,
	}
}

func (w *Weenie) Migrate() error {
	conn := w.conn.AutoMigrate(w)
	conn = conn.Model(w).AddForeignKey("story_id", "stories(id)", "RESTRICT", "CASCADE")

	errors := conn.GetErrors()
	return mergeErrors(errors)
}

func (w *Weenie) Load() error {
	return load(w.conn, w, w.ID)
}

func (w *Weenie) Insert() error {
	return insert(w.conn, w)
}

func (w *Weenie) Update() error {
	return update(w.conn, w)
}

func (w *Weenie) Delete() error {
	if w.ID == 0 {
		return fmt.Errorf("Missing primary key 'id' when deleting a weenie model")
	}

	return delete(w.conn, w)
}
