package routers


import (
	"github.com/astaxie/beego"
	"voicematch/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/ws", &controllers.MyWebSocketController{})
}