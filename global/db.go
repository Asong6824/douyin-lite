package global

import (
	"github.com/jinzhu/gorm"
	"github.com/go-redis/redis/v8"
)

var (
	DBEngine *gorm.DB
	Redis    *redis.Client
)