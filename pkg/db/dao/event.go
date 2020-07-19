package dao

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Event struct {
	gorm.Model
	Name    *string `gorm:"type:varchar(128);not null"`
	SeqID   uint    `gorm:"not null"`
	StoryID uint    `gorm:"not null"`
	Actions []Action
	conn    *gorm.DB `gorm:"-"`
}

func NewEvent(db *gorm.DB, name *string, seqID, storyID uint) *Event {
	return &Event{
		Name:    name,
		SeqID:   seqID,
		StoryID: storyID,
		conn:    db,
	}
}

func (e *Event) Migrate() error {
	conn := e.conn.AutoMigrate(e)
	conn = conn.Model(e).AddForeignKey("story_id", "stories(id)", "RESTRICT", "CASCADE")

	errors := e.conn.GetErrors()
	return mergeErrors(errors)
}

func (e *Event) Load() error {
	return load(e.conn, e, e.ID)
}

func (e *Event) Insert() error {
	return insert(e.conn, e)
}

func (e *Event) Update() error {
	return update(e.conn, e)
}

func (e *Event) Delete() error {
	if e.ID == 0 {
		return fmt.Errorf("Missing primary key 'id' when deleting an event model")
	}

	return delete(e.conn, e)
}
