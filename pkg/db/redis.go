package db

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gochat/config"
	"runtime"
)

var Redis *redis.Client

func initRedis() {
	conf := config.Redis
	Redis = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password:     conf.Password,
		DB:           conf.Db,
		PoolSize:     runtime.NumCPU() * 30,
		MinIdleConns: 10,
	})
	logrus.Infof("redis success %s", Redis.Ping().Val())
}
