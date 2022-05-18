package controllers

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type QueryDeploymentController struct {
	MainController
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
