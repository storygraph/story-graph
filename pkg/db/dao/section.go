package dao

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Section struct {
	gorm.Model
	Name    *string  `gorm:"type:varchar(256);not null"`
	SeqID   uint     `gorm:"not null"`
	StoryID uint     `gorm:"not null"`
	Events  []Event  `gorm:"many2many:section_events"`
	conn    *gorm.DB `gorm:"-"`
}

func NewSection(db *gorm.DB, name *string, seqID, storyID uint) *Section {
	return &Section{
		Name:    name,
		SeqID:   seqID,
		StoryID: storyID,
		conn:    db,
	}
}

func (s *Section) Migrate() error {
	conn := s.conn.AutoMigrate(s)
	conn = conn.Model(s).AddForeignKey("story_id", "stories(id)", "RESTRICT", "CASCADE")

	errors := s.conn.GetErrors()
	return mergeErrors(errors)
}

func (s *Section) Load() error {
	return load(s.conn, s, s.ID)
}

func (s *Section) Insert() error {
	return insert(s.conn, s)
}

func (s *Section) Update() error {
	return update(s.conn, s)
}

func (s *Section) Delete() error {
	if s.ID == 0 {
		return fmt.Errorf("Missing primary key 'id' when deleting section model")
	}

	return delete(s.conn, s)
}
