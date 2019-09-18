package main

import (
	"flag"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"voicematch/models"
	_ "voicematch/routers"
)
var g_logfile string
var g_config string

func read_config() {
	beego.LoadAppConfig("ini", g_config)
	logfile := fmt.Sprintf("{\"filename\":\"%s\"}", g_logfile) //get new now time, remove the http request time!!!
	beego.SetLogger("file", logfile)
	beego.SetLogFuncCall(false)

	models.RegisterDB()
	// 开启 ORM 调试模式
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)
}

func getArgs() error {
	//flag.StringVar(&g_config, "c", "/data/etc/conf/microapi.conf", "配置文件")
	//flag.StringVar(&g_logfile, "l", "/data/logs/microapi.log", "日志文件")
	flag.StringVar(&g_config, "c", "../voicematch/conf/app.conf", "配置文件")
	flag.Parse()
	return nil
}

func main() {
	getArgs()
	read_config()
	beego.Run()
}

