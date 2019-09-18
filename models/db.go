package models

import (
	"fmt"
	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type CallLogs struct {
	Id        int64  `orm:"index"`
	ValidCode string `orm:"index"`
	SpeakerId string `orm:"null"`
	Name      string `orm:"null"`
	Company   string `orm:"null"`
	Address   string `orm:"null"`
}

func RegisterDB() {
	//注册 model
	orm.RegisterModel(new(CallLogs))
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册默认数据库
	dbuser := beego.AppConfig.String("mysql::user")
	dbpassword := beego.AppConfig.String("mysql::password")
	dbhost := beego.AppConfig.String("mysql::host")
	dbport := beego.AppConfig.String("mysql::port")
	db_name := beego.AppConfig.String("mysql::database")
	conn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + db_name + "?charset=utf8"
	fmt.Println("RegisterDB： conn[", conn, "]")
	orm.RegisterDataBase("default", "mysql", conn) //密码为空格式
}
