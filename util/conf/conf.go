package conf

import (
	"fmt"
	"github.com/go-ini/ini"
)

var (
	HTTP_PORT string

	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string

	JWT_SECRET string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取失败:", err)
	}
	LoadServer(file)
	LoadDB(file)
	LoadJWT(file)
}

func LoadServer(file *ini.File) {
	//RunMODE = file.Section("server").Key("RunMode").MustString("debug")
	HTTP_PORT = file.Section("server").Key("HTTP_PORT").MustString("8080")
}

func LoadDB(file *ini.File) {
	DB_HOST = file.Section("database").Key("DB_HOST").MustString("127.0.0.1")
	DB_PORT = file.Section("database").Key("DB_PORT").MustString("3306")
	DB_USER = file.Section("database").Key("DB_USER").MustString("root")
	DB_PASSWORD = file.Section("database").Key("DB_PASSWORD").MustString("root")
	DB_NAME = file.Section("database").Key("DB_NAME").MustString("")
}

func LoadJWT(file *ini.File) {
	JWT_SECRET = file.Section("jwt").Key("JWT_SECRET").MustString("123456")
}
