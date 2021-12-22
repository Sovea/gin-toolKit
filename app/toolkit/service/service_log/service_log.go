package service_log

import (
	"toolKit/app/toolkit/service/service_path"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	logrus "github.com/sirupsen/logrus"
	"path"
	"time"
)

var (
	Users_Logger   = logrus.New()
	Service_Logger = logrus.New()
)

func Get_UsersLogger() *logrus.Logger {
	return Users_Logger
}

func Get_ServiceLogger() *logrus.Logger {
	return Service_Logger
}

func init() {
	current_path, err := service_path.GetCurrentPath()
	if err != nil {
		fmt.Println(err)
		return
	}
	users_log_path := path.Join(current_path, "/log", "/users_log") + "/Users_Log.log"
	service_log_path := path.Join(current_path, "/log", "/service_log") + "/Service_Log.log"
	users_log_writer, _ := rotatelogs.New(
		users_log_path+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(users_log_path),
		rotatelogs.WithMaxAge(time.Duration(72)*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(60)*time.Second),
	)
	Users_Logger.SetOutput(users_log_writer)
	service_log_writer, _ := rotatelogs.New(
		service_log_path+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(service_log_path),
		rotatelogs.WithMaxAge(time.Duration(72)*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(60)*time.Second),
	)
	Service_Logger.SetOutput(service_log_writer)
}
