package routers

import (
	"expense-management-api/controllers"
	"expense-management-api/middlewares"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.InsertFilter(
		"/api/v1/categories",
		beego.BeforeRouter,
		middlewares.AuthFilter,
	)

	beego.InsertFilter(
		"/api/v1/categories/*",
		beego.BeforeRouter,
		middlewares.AuthFilter,
	)

	beego.InsertFilter(
		"/api/v1/expenses",
		beego.BeforeRouter,
		middlewares.AuthFilter,
	)

	beego.InsertFilter(
		"/api/v1/expenses/*",
		beego.BeforeRouter,
		middlewares.AuthFilter,
	)
	
	ns := beego.NewNamespace("/api/v1",

		beego.NSRouter(
			"/health",
			&controllers.HealthController{},
		),

		beego.NSNamespace(
			"/auth",

			beego.NSRouter(
				"/register",
				&controllers.AuthController{},
				"post:Register",
			),

			beego.NSRouter(
				"/login",
				&controllers.AuthController{},
				"post:Login",
			),
		),

		beego.NSNamespace(
			"/categories",

			beego.NSRouter(
				"",
				&controllers.CategoryController{},
				"post:Create;get:GetByUserID",
			),

			beego.NSRouter(
				"/:id",
				&controllers.CategoryController{},
				"get:GetByID;put:Update;delete:Delete",
			),
		),

		beego.NSNamespace(
			"/expenses",

			beego.NSRouter(
				"",
				&controllers.ExpenseController{},
				"post:Create;get:GetAll",
			),

			beego.NSRouter(
				"/:id",
				&controllers.ExpenseController{},
				"get:GetByID;put:Update;delete:Delete",
			),
		),
	)

	beego.AddNamespace(ns)
}
