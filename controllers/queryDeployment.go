package controllers

import (
	beego "github.com/beego/beego/v2/server/web"

	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type QueryDeploymentController struct {
	beego.Controller
}

func (c *QueryDeploymentController) Get() {
	values, err := c.Input()
	if err != nil {
		log.Printf("c.Input() failed: %v \n", err)
	}

	namespace := values.Get("namespace")
	deployments, err := ClientSet.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("ClientSet.AppsV1().Deployments().List() failed: %v \n", err)
	}

	c.Data["json"] = deployments
	c.ServeJSON()
}
