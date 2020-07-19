package dao

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
)

const DEPENDENCY_ID_DELIMITER = ":"

type Action struct {
	gorm.Model
	Name             *string `gorm:"type:varchar(64);not null"`
	DependentActions []uint  `gorm:"type:varchar(512);not nul"`
	EventID          uint    `gorm:"not null"`
	WeenieDeltaID    uint
	RelationDeltaID  uint
	WeenieDelta      WeenieDelta
	RelationDelta    RelationDelta
	conn             *gorm.DB `gorm:"-"`
}

func NewAction(db *gorm.DB, name *string, dependentActions []uint, weenieDeltaID, relationDeltaID, eventID uint) *Action {
	return &Action{
		Name:             name,
		DependentActions: dependentActions,
		WeenieDeltaID:    weenieDeltaID,
		RelationDeltaID:  relationDeltaID,
		EventID:          eventID,
		conn:             db,
	}
}

func (a *Action) BeforeCreate(scope *gorm.Scope) error {
	a.dependentActionsHook(scope)
	return nil
}

func (a *Action) BeforeSave(scope *gorm.Scope) error {
	a.dependentActionsHook(scope)
	return nil
}

func (a *Action) Migrate() error {
	conn := a.conn.AutoMigrate(a)

	conn = conn.Model(a).AddForeignKey("weenie_delta_id", "weenie_deltas(id)", "RESTRICT", "CASCADE")
	conn = conn.Model(a).AddForeignKey("relation_delta_id", "relation_deltas(id)", "RESTRICT", "CASCADE")
	conn = conn.Model(a).AddForeignKey("event_id", "events(id)", "RESTRICT", "CASCADE")

	errors := conn.GetErrors()
	return mergeErrors(errors)
}

func (a *Action) Load() error {
	return load(a.conn, a, a.ID)
}

func (a *Action) Insert() error {
	return insert(a.conn, a)
}

func (a *Action) Update() error {
	return update(a.conn, a)
}

func (a *Action) Delete() error {
	if a.ID == 0 {
		return fmt.Errorf("Missing primary key 'id' when deleting an action model")
	}

	return delete(a.conn, a)
}

func (a *Action) dependentActionsHook(scope *gorm.Scope) {
	serializedDependentActions := a.serializeDependentActions()
	scope.SetColumn("DependentActions", serializedDependentActions)
}

func (a *Action) serializeDependentActions() (result string) {
	if a.DependentActions == nil || len(a.DependentActions) == 0 {
		return result
	}

	for i := 0; i < len(a.DependentActions)-1; i++ {
		result += fmt.Sprintf("%d%s", a.DependentActions[i], DEPENDENCY_ID_DELIMITER)
	}
	result += strconv.Itoa(int(a.DependentActions[len(a.DependentActions)-1]))

	return result
}
