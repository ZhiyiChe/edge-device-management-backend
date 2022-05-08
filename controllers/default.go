package controllers

import (
	"log"

	beego "github.com/beego/beego/v2/server/web"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var ClientSet *kubernetes.Clientset

func init() {
	kubeConfig := "/etc/rancher/k3s/k3s.yaml"
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		log.Printf("clientcmd.BuildConfigFromFlags() failed: %v \n", err)
	}
	// create the clientset
	ClientSet, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Printf("kubernetes.NewForConfig() failed: %v \n", err)
	}
}

type CommonResponse struct {
	Code int
	Data interface{}
}

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
