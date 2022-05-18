package controllers

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type QueryPodController struct {
	MainController
}

func (c *QueryPodController) Get() {
	values, err := c.Input()
	if err != nil {
		log.Printf("c.Input() failed: %v \n", err)
	}

	namespace := values.Get("namespace")
	pods, err := ClientSet.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("ClientSet.CoreV1().Pods().List() failed: %v \n", err)
	}

	c.Data["json"] = pods
	c.ServeJSON()
}
