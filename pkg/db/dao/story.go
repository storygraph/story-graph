package dao

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Story struct {
	gorm.Model
	Name     *string `gorm:"type:varchar(256);not null"`
	Weenies  []Weenie
	Events   []Event
	Sections []Section
	conn     *gorm.DB `gorm:"-"`
}

func NewStory(db *gorm.DB, name *string) *Story {
	return &Story{
		Name: name,
		conn: db,
	}
}

func (s *Story) TableName() string {
	return "stories"
}

func (s *Story) Migrate() error {
	errors := s.conn.AutoMigrate(s).GetErrors()
	return mergeErrors(errors)
}

func (s *Story) Load() error {
	return load(s.conn, s, s.ID)
}

func (s *Story) Insert() error {
	return insert(s.conn, s)
}

func (s *Story) Update() error {
	return update(s.conn, s)
}

func (s *Story) Delete() error {
	if s.ID == 0 {
		return fmt.Errorf("Missing primary key 'id' when deleting a story model")
	}

	return delete(s.conn, s)
}
