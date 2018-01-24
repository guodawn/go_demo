package controllers



type HomeController struct {
	BaseController
}

func (c *HomeController) Index() {

	c.Layout = "common/layout.tpl"
	c.TplName = "home/index.tpl"
	
}
