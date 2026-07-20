package controllers

type HealthController struct {
	BaseController
}

func (c *HealthController) Get() {
	c.Data["json"] = map[string]any{
		"success": true,
		"message": "Server is running.",
	}
	c.ServeJSON()
}
