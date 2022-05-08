package controllers

import (
	"context"
	"log"

	beego "github.com/beego/beego/v2/server/web"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ApproveNodeController struct {
	beego.Controller
}

func (c *ApproveNodeController) Post() {
	values, err := c.Input()
	if err != nil {
		log.Printf("c.Input() failed: %v \n", err)
	}

	nodeName := values.Get("nodeName")
	node, err := ClientSet.CoreV1().Nodes().Get(context.TODO(), nodeName, metav1.GetOptions{})
	if err != nil {
		log.Printf("ClientSet.CoreV1().Nodes().Get() failed: %v \n", err)
	}

	(*node).ObjectMeta.Labels["isApproved"] = "yes" // add label
	node, err = ClientSet.CoreV1().Nodes().Update(context.TODO(), node, metav1.UpdateOptions{})
	if err != nil {
		log.Printf("ClientSet.CoreV1().Nodes().Update() failed: %v \n", err)
	}

	c.Data["json"] = node
	c.ServeJSON()
}
