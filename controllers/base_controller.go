package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
}

func (c *CategoryController) getUserID() (int, bool) {
	userID, ok := c.Ctx.Input.GetData("userID").(int)
	return userID, ok
}
