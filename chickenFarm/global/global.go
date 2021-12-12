package global

import "gorm.io/gorm"

var DB *gorm.DB

// DB = gorm.Open(sqlite.Open("/www/xj/chickenFarm/db/sqlite.db"), &gorm.Config{})
