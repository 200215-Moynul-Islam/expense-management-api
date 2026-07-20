package main

import (
	"expense-management-api/database"
	_ "expense-management-api/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	database.Init()
	
	beego.Run()
}
