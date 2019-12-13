package repositories

import "github.com/jinzhu/gorm"

type Database interface {
	Find(out interface{}, where ...interface{}) *gorm.DB
}