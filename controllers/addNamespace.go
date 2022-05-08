package controllers

import (
	beego "github.com/beego/beego/v2/server/web"

	"context"
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type AddNamespaceController struct {
	beego.Controller
}

func (c *AddNamespaceController) Post() {
	values, err := c.Input()
	if err != nil {
		log.Printf("c.Input() failed: %v \n", err)
	}

	namespaceName := values.Get("namespaceName")
	namespace := &v1.Namespace{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Namespace",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: namespaceName,
		},
	}

	namespace, err = ClientSet.CoreV1().Namespaces().Create(context.TODO(), namespace, metav1.CreateOptions{})
	if err != nil {
		log.Printf("ClientSet.CoreV1().Namespaces().Create() failed: %v \n", err)
		c.Data["json"] = &CommonResponse{
			Code: -1,
			Data: err.Error(),
		}
	} else {
		c.Data["json"] = &CommonResponse{
			Code: 0,
			Data: *namespace,
		}
	}

	c.ServeJSON()
}
