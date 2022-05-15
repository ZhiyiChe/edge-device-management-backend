package controllers

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	beego "github.com/beego/beego/v2/server/web"

	"log"
)

type ExecKubectlController struct {
	beego.Controller
}

func (c *ExecKubectlController) Get() {
	values, err := c.Input()
	if err != nil {
		log.Printf("c.Input() failed: %v \n", err)
	}

	kubectlCmd := values.Get("kubectlCmd")
	fmt.Println(kubectlCmd)
	if len(kubectlCmd) < 7 || kubectlCmd[0:7] != "kubectl" || strings.Contains(kubectlCmd, "&") || strings.Contains(kubectlCmd, "|") {
		c.Ctx.WriteString("command must begin with 'kubectl', and cannot include '&' or '|'")
	} else {
		args := strings.Split(kubectlCmd, " ")
		args = args[1:] // 删除开头的"kubectl"
		fmt.Println(args)
		cmd := exec.Command("kubectl", args...)

		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = os.Stderr
		if err := cmd.Start(); err != nil {
			log.Printf("cmd.Start() failed: %v \n", err)
		}
		if err := cmd.Wait(); err != nil {
			log.Printf("cmd.Wait() failed: %v \n", err)
		}

		c.Ctx.WriteString(out.String())
	}
}
