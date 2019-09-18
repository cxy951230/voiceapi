package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"voicematch/models"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"time"
)

type MyWebSocketController struct {
	beego.Controller
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},

}
var oldid="0"
func (c *MyWebSocketController) Get() {
	user := beego.AppConfig.String("mysql::user")
	pwd := beego.AppConfig.String("mysql::password")
	host := beego.AppConfig.String("mysql::host")
	port := beego.AppConfig.String("mysql::port")
	database := beego.AppConfig.String("mysql::database")
	db, err := sql.Open("mysql", user+":"+pwd+"@tcp("+host+":"+port+")/"+database+"?charset=utf8")


	ws, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		log.Fatal(err)
	}
	//  defer ws.Close()

	clients[ws] = true

	//不断的广播发送到页面上
	for {
		//目前存在问题 定时效果不好 需要在业务代码替换时改为beego toolbox中的定时器
		//time.Sleep(time.Second * 3)
		id:=""
		callid:=""
		speaker_id:=""
		phone:=""
		name:=""
		customerRow := db.QueryRow("select id,callid,speaker_id,phone,name from call_logs ORDER BY id desc limit 1")
		customerRow.Scan(&id,&callid, &speaker_id, &phone, &name)
		if id!=oldid{
			oldid=id
			msg := models.Message{Message: "callid:"+callid+"，speaker_id:"+speaker_id+"，phone:"+phone+"，name:"+name +", "+time.Now().Format("2006-01-02 15:04:05")}
			broadcast <- msg
		}

	}
}
