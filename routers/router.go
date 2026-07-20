package routers

import (
	"expense-management-api/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/api/v1",

		beego.NSRouter(
			"/health",
			&controllers.HealthController{},
		),

		beego.NSRouter(
			"/auth/register",
			&controllers.AuthController{},
			"post:Register",
		),

		beego.NSRouter(
			"/auth/login",
			&controllers.AuthController{},
			"post:Login",
		),
	)

	beego.AddNamespace(ns)
}