package controllers

type DownloadK3sController struct {
	MainController
}

func (c *DownloadK3sController) Get() {
	c.Ctx.Output.Download("install/k3s", "k3s")
}
