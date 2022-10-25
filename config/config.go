package config

import (
	"github.com/shiena/ansicolor"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"sync"
)

var (
	Api   *ApiConfig
	Mysql *MysqlConfig
	Redis *RedisConfig
	once  sync.Once
)

type ApiConfig struct {
	Address string
	Port    int
}

type MysqlConfig struct {
	Host     string
	Port     int
	Account  string
	Password string
	Db       string
	Charset  string
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	Db       int
}

func Init() {
	once.Do(func() {

		initLogrus()

		viper.SetConfigName("config.toml")
		viper.SetConfigType("toml")
		viper.AddConfigPath(".")
		viper.AddConfigPath("./config")
		viper.AddConfigPath("/etc/")
		viper.AddConfigPath("$HOME/.config/")

		err := viper.ReadInConfig()
		if err != nil {
			logrus.Error(err.Error())
			return
		}

		c := struct {
			Api   *ApiConfig
			Mysql *MysqlConfig
			Redis *RedisConfig
		}{}

		err = viper.Unmarshal(&c)
		if err != nil {
			logrus.Error(err.Error())
			return
		}

		logrus.WithFields(logrus.Fields{
			"api":   *c.Api,
			"mysql": *c.Mysql,
			"redis": *c.Redis,
		}).Info("config : ")

		Api = c.Api
		Mysql = c.Mysql
		Redis = c.Redis

	})
}

func initLogrus() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	logrus.SetOutput(ansicolor.NewAnsiColorWriter(os.Stdout))
	logrus.SetLevel(logrus.DebugLevel)
}
