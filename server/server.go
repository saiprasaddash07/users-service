package server

import (
	"github.com/saiprasaddash07/users-service/config"
	"github.com/saiprasaddash07/users-service/services/db"
	"github.com/saiprasaddash07/users-service/services/logger"
	"github.com/saiprasaddash07/users-service/services/redis"
)

func Init() {
	config := config.Get()
	logger.InitLogger()
	db.Init()
	redis.Init()
	r := NewRouter()
	r.Run(":" + config.ServerPort)
}
