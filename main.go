package main

import (
	_ "myblog/routers"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/astaxie/beego/session/mysql"
	"github.com/astaxie/beego"
	_ "github.com/gomodule/redigo/redis"
	"myblog/utils"
)

func main() {

	utils.InitMysql()
	beego.Run()
}

