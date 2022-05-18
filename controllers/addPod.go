package controllers

import (
	"context"
	"encoding/json"
	"fmt"

	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type AddPodController struct {
	MainController
}

func (c *AddPodController) Post() {
	pod := &v1.Pod{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, pod)
	if err != nil {
		log.Printf("json.Unmarshal() failed: %v \n", err)
	}

	fmt.Println(c.Ctx.Input.RequestBody)
	fmt.Println(pod)

	pod.Spec.NodeSelector["isApproved"] = "yes"

	namespace := pod.ObjectMeta.Namespace
	pod, err = ClientSet.CoreV1().Pods(namespace).Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		log.Printf("ClientSet.CoreV1().Pods().Create() failed: %v \n", err)
		c.Data["json"] = &CommonResponse{
			Code: -1,
			Data: err.Error(),
		}
	} else {
		c.Data["json"] = &CommonResponse{
			Code: 0,
			Data: *pod,
		}
	}

	c.ServeJSON()
}
