// database/database.go
package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	// Thông tin kết nối cơ sở dữ liệu
	dsn := "hung:dn2V3NSU13zOFq8@tcp(localhost:3306)/users_nfruit?charset=utf8mb4&parseTime=True&loc=Local"

	// Kết nối đến cơ sở dữ liệu
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
