package controllers

type DownloadInstallScriptController struct {
	MainController
}

func (c *DownloadInstallScriptController) Get() {
	c.Ctx.Output.Download("install/install.sh", "install.sh")
}
