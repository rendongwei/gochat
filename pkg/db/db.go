package db

import "sync"

var once sync.Once

func Init() {
	once.Do(func() {
		initMysql()
		initRedis()
	})
}
