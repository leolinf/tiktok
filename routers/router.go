package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"tiktok/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/api/v1",
		beego.NSCtrlGet("/user", (*controllers.MainController).Get),
	)
	beego.AddNamespace(ns)
}
