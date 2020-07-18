package dto

import "github.com/jinzhu/gorm"

type RelationType struct {
	Name *string `gorm:"primary_key"`
}

type Relation struct {
	gorm.Model
	Description  string
	Name         *string `gorm:"column:type"`
	RelationType RelationType
}

func (rt *RelationType) Migrate(db *gorm.DB) []error {
	return db.AutoMigrate(&RelationType{}).GetErrors()
}

func (rt *RelationType) GetPrimaryKeyName() string {
	return "name"
}

func (rt *RelationType) GetPrimaryKeyVal() interface{} {
	return *rt.Name
}

func (r *Relation) Migrate(db *gorm.DB) []error {
	errors := db.AutoMigrate(r).Model(r).AddForeignKey("type", "relation_types(name)", "RESTRICT", "CASCADE").GetErrors()
	if len(errors) > 0 {
		return errors
	}

	return []error{}
}

func (r *Relation) GetPrimaryKeyName() string {
	return "id"
}

func (r *Relation) GetPrimaryKeyVal() interface{} {
	return r.ID
}
