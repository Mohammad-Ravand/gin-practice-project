package models
import (
	"gorm.io/gorm"
)
// Class model
type ClassRoom struct {
	gorm.Model
	Id   int
	Name string
}
