// product.go
package models
import "gorm.io/gorm"
type Product struct {
	gorm.Model
	Code  string
	Price uint
}
func CreateProduct(db *gorm.DB, product *Product) error {
	return db.Create(product).Error
}
// Define other database operations (Read, Update, Delete) similarly...
