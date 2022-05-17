package controllers

import (
	"context"
	"encoding/json"

	beego "github.com/beego/beego/v2/server/web"

	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type AddServiceController struct {
	beego.Controller
}

func (c *AddServiceController) Post() {
	service := &v1.Service{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, service)
	if err != nil {
		log.Printf("json.Unmarshal() failed: %v \n", err)
	}

	namespace := service.ObjectMeta.Namespace
	_, err = ClientSet.CoreV1().Services(namespace).Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		log.Printf("ClientSet.CoreV1().Services().Create() failed: %v \n", err)
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
