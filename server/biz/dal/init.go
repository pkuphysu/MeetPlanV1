package dal

import (
	"meetplan/biz/dal/mysql"
	"meetplan/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

func Close() {
	redis.Close()
	mysql.Close()
}