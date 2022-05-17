package controllers

import (
	"context"

	beego "github.com/beego/beego/v2/server/web"

	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DeleteDeploymentController struct {
	beego.Controller
}

func (c *DeleteDeploymentController) Get() {
	values, err := c.Input()
	if err != nil {
		log.Printf("c.Input() failed: %v \n", err)
	}

	deploymentName := values.Get("name")
	namespace := values.Get("namespace")
	err = ClientSet.AppsV1().Deployments(namespace).Delete(context.TODO(), deploymentName, metav1.DeleteOptions{})
	if err != nil {
		log.Printf("ClientSet.AppsV1().Deployments().Delete() failed: %v \n", err)
		c.Data["json"] = &CommonResponse{
			Code: -1,
			Data: err.Error(),
		}
	} else {
		c.Data["json"] = &CommonResponse{
			Code: 0,
			Data: "success",
		}
	}

	c.ServeJSON()
}
