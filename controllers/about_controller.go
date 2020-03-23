package controllers

type AboutMeController struct {
	BaseController
}

func (c *AboutMeController) Get() {

	c.Data["qq"] = "QQ：1686125590"
	c.TplName = "aboutme.html"
}
