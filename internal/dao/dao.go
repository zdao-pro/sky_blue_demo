package dao

import (
	"github.com/zdao-pro/sky_blue/pkg/cache/redis"
	"github.com/zdao-pro/sky_blue/pkg/database/sql"
	"github.com/zdao-pro/sky_blue/pkg/peach"
)

//mysql
var testDB *sql.DB

//redis
var testRedis *redis.Redis

//Init ..
func Init() {
	//初始化mysql客户端
	var c sql.Config
	peach.Get("mysql_test.yaml").UnmarshalYAML(&c)
	testDB = sql.NewMySQL(&c)

	//初始化redis客户端
	var redisConfig redis.NewConfig
	peach.Get("redis_test.yaml").UnmarshalYAML(&redisConfig)
	testRedis = redis.NewRedisClient(&redisConfig)
}
