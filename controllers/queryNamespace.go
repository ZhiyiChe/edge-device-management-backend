package controllers

import (
	beego "github.com/beego/beego/v2/server/web"

	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type QueryNamespaceController struct {
	beego.Controller
}

func (c *QueryNamespaceController) Get() {
	namespaces, err := ClientSet.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("ClientSet.CoreV1().Namespaces().List() failed: %v \n", err)
	}

	c.Data["json"] = namespaces
	c.ServeJSON()
}
