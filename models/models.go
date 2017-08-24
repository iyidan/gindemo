package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // gorm mysql implements

	"github.com/iyidan/gindemo/conf"
	"github.com/iyidan/goutils/mise"
)

var (
	db *gorm.DB
)

// Startup db
func Startup() {
	if db != nil {
		return
	}
	var err error
	db, err = gorm.Open("mysql", conf.String("mysql"))
	if err != nil {
		mise.PanicOnError(err, "models.Startup")
	}
}

// Close close db
func Close() {
	if db != nil {
		db.Close()
	}
}
