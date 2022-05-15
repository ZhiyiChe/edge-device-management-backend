package controllers

import (
	beego "github.com/beego/beego/v2/server/web"

	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type QueryServiceController struct {
	beego.Controller
}

func (c *QueryServiceController) Get() {
	values, err := c.Input()
	if err != nil {
		log.Printf("c.Input() failed: %v \n", err)
	}

	namespace := values.Get("namespace")
	services, err := ClientSet.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("ClientSet.CoreV1().Services().List() failed: %v \n", err)
	}

	c.Data["json"] = services
	c.ServeJSON()
}
