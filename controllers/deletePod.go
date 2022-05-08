package controllers

import (
	beego "github.com/beego/beego/v2/server/web"

	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DeletePodController struct {
	beego.Controller
}

func (c *DeletePodController) Post() {
	values, err := c.Input()
	if err != nil {
		log.Printf("c.Input() failed: %v \n", err)
	}

	podName := values.Get("podName")
	namespace := values.Get("namespace")
	err = ClientSet.CoreV1().Pods(namespace).Delete(context.TODO(), podName, metav1.DeleteOptions{})
	if err != nil {
		log.Printf("ClientSet.CoreV1().Pods().Delete() failed: %v \n", err)
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
