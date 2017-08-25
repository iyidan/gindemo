package models

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // gorm mysql implements

	"github.com/iyidan/gindemo/conf"
	"github.com/iyidan/goutils/mise"
)

// BaseModel base model
type BaseModel struct {
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

// BeforeCreate callback
func (base *BaseModel) BeforeCreate() error {
	//log.Info("BeforeCreate called")
	nowtm := time.Now().Unix()
	base.CreatedAt = nowtm
	base.UpdatedAt = nowtm
	return nil
}

// BeforeUpdate callback
func (base *BaseModel) BeforeUpdate() error {
	//log.Info("BeforeUpdate called")
	base.UpdatedAt = time.Now().Unix()
	return nil
}

var (
	// DB exported to other
	DB *gorm.DB
)

// Startup db
func Startup() {
	if DB != nil {
		return
	}
	var err error
	DB, err = gorm.Open("mysql", conf.String("mysql"))
	if err != nil {
		mise.PanicOnError(err, "models.Startup")
	}
	DB.SingularTable(true)
	// Enable Logger, show detailed log
	DB.LogMode(true)
	// logger
	//DB.SetLogger(log.DefaultLogger)
}

// Close close db
func Close() {
	if DB != nil {
		DB.Close()
	}
}

// IsNotExists whether err is gorm.ErrRecordNotFound
func IsNotExists(err error) bool {
	if err == gorm.ErrRecordNotFound {
		return true
	}
	return false
}
