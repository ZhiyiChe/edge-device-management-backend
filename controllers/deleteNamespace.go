package controllers

import (
	beego "github.com/beego/beego/v2/server/web"

	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DeleteNamespaceController struct {
	beego.Controller
}

func (c *DeleteNamespaceController) Post() {
	values, err := c.Input()
	if err != nil {
		log.Printf("c.Input() failed: %v \n", err)
	}

	namespaceName := values.Get("namespaceName")
	err = ClientSet.CoreV1().Namespaces().Delete(context.TODO(), namespaceName, metav1.DeleteOptions{})
	if err != nil {
		log.Printf("ClientSet.CoreV1().Namespaces().Delete() failed: %v \n", err)
		c.Data["json"] = &CommonResponse{
			Code: -1,
			Data: err.Error(),
		}
	} else {
		c.Data["json"] = &CommonResponse{
			Code: 0,
		}
	}

	c.ServeJSON()
}
