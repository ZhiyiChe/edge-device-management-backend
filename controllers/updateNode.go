package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	beego "github.com/beego/beego/v2/server/web"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type UpdateNodeController struct {
	beego.Controller
}

func (c *UpdateNodeController) Post() {
	node := &v1.Node{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, node)
	if err != nil {
		log.Printf("json.Unmarshal() failed: %v \n", err)
	}
	fmt.Println(node.ObjectMeta.Labels)

	// 通过nodeName查找Node
	nodeTmp, err := ClientSet.CoreV1().Nodes().Get(context.TODO(), node.ObjectMeta.Name, metav1.GetOptions{})
	if err != nil {
		log.Printf("ClientSet.CoreV1().Nodes().Get() failed: %v \n", err)
		c.Data["json"] = &CommonResponse{
			Code: -1,
			Data: "failure",
		}
		goto END
	}

	// 更新Node
	nodeTmp.ObjectMeta.Labels = node.ObjectMeta.Labels
	_, err = ClientSet.CoreV1().Nodes().Update(context.TODO(), nodeTmp, metav1.UpdateOptions{})
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
