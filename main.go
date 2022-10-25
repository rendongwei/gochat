package main

import (
	"github.com/sirupsen/logrus"
	"gochat/config"
	"gochat/pkg/db"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	config.Init()

	db.Init()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit

	logrus.Info("main quit")

}
