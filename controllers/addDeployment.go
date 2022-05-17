package controllers

import (
	"context"
	"encoding/json"

	beego "github.com/beego/beego/v2/server/web"

	"log"

	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type AddDeploymentController struct {
	beego.Controller
}

func (c *AddDeploymentController) Post() {
	deployment := &v1.Deployment{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, deployment)
	if err != nil {
		log.Printf("json.Unmarshal() failed: %v \n", err)
	}

	deployment.Spec.Template.Spec.NodeSelector["isApproved"] = "yes"

	namespace := deployment.ObjectMeta.Namespace
	_, err = ClientSet.AppsV1().Deployments(namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		log.Printf("ClientSet.AppsV1().Deployments().Create() failed: %v \n", err)
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
