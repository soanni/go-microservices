package user

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Printf("uuid.NewV4() something went wrong: %s", err)
		return err
	}
	return scope.SetColumn("Id", uuid.String())
}