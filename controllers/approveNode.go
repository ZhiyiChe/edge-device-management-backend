package controllers

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ApproveNodeController struct {
	MainController
}

func (c *ApproveNodeController) Get() {
	values, err := c.Input()
	if err != nil {
		log.Printf("c.Input() failed: %v \n", err)
	}

	nodeName := values.Get("nodeName")
	node, err := ClientSet.CoreV1().Nodes().Get(context.TODO(), nodeName, metav1.GetOptions{})
	if err != nil {
		log.Printf("ClientSet.CoreV1().Nodes().Get() failed: %v \n", err)
		c.Data["json"] = &CommonResponse{
			Code: -1,
			Data: "failure",
		}
		goto END
	}

	(*node).ObjectMeta.Labels["isApproved"] = "yes" // add label
	_, err = ClientSet.CoreV1().Nodes().Update(context.TODO(), node, metav1.UpdateOptions{})
	if err != nil {
		log.Printf("ClientSet.CoreV1().Nodes().Update() failed: %v \n", err)
		c.Data["json"] = &CommonResponse{
			Code: -1,
			Data: "failure",
		}
		goto END
	}

	c.Data["json"] = &CommonResponse{
		Code: 0,
		Data: "success",
	}

END:
	c.ServeJSON()
}
