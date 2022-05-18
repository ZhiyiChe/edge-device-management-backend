package controllers

type DownloadK3sImagesController struct {
	MainController
}

func (c *DownloadK3sImagesController) Get() {
	c.Ctx.Output.Download("install/k3s-airgap-images-amd64.tar", "k3s-airgap-images-amd64.tar")
}
