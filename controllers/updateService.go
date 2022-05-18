package controllers

import (
	"context"
	"encoding/json"

	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type UpdateServiceController struct {
	MainController
}

func (c *UpdateServiceController) Post() {
	service := &v1.Service{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, service)
	if err != nil {
		log.Printf("json.Unmarshal() failed: %v \n", err)
	}

	namespace := service.ObjectMeta.Namespace
	_, err = ClientSet.CoreV1().Services(namespace).Update(context.TODO(), service, metav1.UpdateOptions{})
	if err != nil {
		log.Printf("ClientSet.CoreV1().Pods().Create() failed: %v \n", err)
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
