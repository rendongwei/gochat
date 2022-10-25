package db

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gochat/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var DB *gorm.DB

func initMysql() {

	conf := config.Mysql

	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true",
		conf.Account, conf.Password, conf.Host, conf.Port, conf.Db, conf.Charset)

	var err error

	DB, err = gorm.Open(mysql.Open(url), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "im_",
			SingularTable: true,
			//NameReplacer:  nil,
			//NoLowerCase:   false,
		},
	})

	if err != nil {
		logrus.Error(err.Error())
	}

	db, err := DB.DB()
	if err != nil {
		logrus.Error(err.Error())
	}

	db.SetMaxOpenConns(10000)
	db.SetMaxIdleConns(1000)
	db.SetConnMaxLifetime(time.Minute * 10)
	db.SetConnMaxIdleTime(time.Minute * 6)

	logrus.Info("mysql success")
}
