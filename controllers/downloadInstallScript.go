package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type DownloadInstallScriptController struct {
	beego.Controller
}

func (c *DownloadInstallScriptController) Get() {
	c.Ctx.Output.Download("install/install.sh", "install.sh")
}
