package controllers

import (
	"context"
	"log"

	beego "github.com/beego/beego/v2/server/web"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DeleteNodeController struct {
	beego.Controller
}

func (c *DeleteNodeController) Get() {
	values, err := c.Input()
	if err != nil {
		log.Printf("c.Input() failed: %v \n", err)
	}

	nodeName := values.Get("nodeName")
	err = ClientSet.CoreV1().Nodes().Delete(context.TODO(), nodeName, metav1.DeleteOptions{})
	if err != nil {
		log.Printf("ClientSet.CoreV1().Nodes().Delete() failed: %v \n", err)
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
