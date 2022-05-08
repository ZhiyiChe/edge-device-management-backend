package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type DownloadK3sController struct {
	beego.Controller
}

func (c *DownloadK3sController) Get() {
	c.Ctx.Output.Download("install/k3s", "k3s")
}
