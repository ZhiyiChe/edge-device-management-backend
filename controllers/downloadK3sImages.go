package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type DownloadK3sImagesController struct {
	beego.Controller
}

func (c *DownloadK3sImagesController) Get() {
	c.Ctx.Output.Download("install/k3s-airgap-images-amd64.tar", "k3s-airgap-images-amd64.tar")
}
