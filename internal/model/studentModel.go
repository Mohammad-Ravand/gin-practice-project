package models
import (
	"gorm.io/gorm"
)
// Student model
type Student struct {
	gorm.Model
	Name        string
	Age         int
	ClassRoomId uint // Foreign key to Class model
}
