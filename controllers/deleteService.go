package controllers

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DeleteServiceController struct {
	MainController
}

func (c *DeleteServiceController) Get() {
	values, err := c.Input()
	if err != nil {
		log.Printf("c.Input() failed: %v \n", err)
	}

	serviceName := values.Get("name")
	namespace := values.Get("namespace")
	err = ClientSet.CoreV1().Services(namespace).Delete(context.TODO(), serviceName, metav1.DeleteOptions{})
	if err != nil {
		log.Printf("ClientSet.CoreV1().Services().Delete() failed: %v \n", err)
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
