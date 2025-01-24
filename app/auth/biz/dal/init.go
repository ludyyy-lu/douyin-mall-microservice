package dal

import (
	"douyinmall/biz/dal/mysql"
	"douyinmall/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
