package controllers

import (
	"bytes"
	"log"
	"os"
	"os/exec"

	"github.com/beego/beego/v2/core/logs"
)

type QueryBeegoLogsController struct {
	MainController
}

func (c *QueryBeegoLogsController) Get() {
	values, err := c.Input()
	if err != nil {
		log.Printf("c.Input() failed: %v \n", err)
	}

	// args := strings.Split(kubectlCmd, " ")
	// args = args[1:] // 删除开头的"kubectl"
	// cmd := exec.Command("kubectl", args...)
	cmd := exec.Command("tail", "/home/ubuntu/edge-device-management-backend/beego.log", "-n", values.Get("num"))
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		logs.Error("cmd.Start() failed: %v \n", err)
		c.Ctx.WriteString(err.Error())
	}
	if err := cmd.Wait(); err != nil {
		logs.Error("cmd.Wait() failed: %v \n", err)
		c.Ctx.WriteString(err.Error())
	}

	if err == nil {
		c.Data["json"] = &CommonResponse{
			Code: 0,
			Data: out.String(),
		}
	}

	c.ServeJSON()
}
