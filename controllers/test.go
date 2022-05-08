package controllers

import (
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
)

type TestResult struct {
	ParaA string
	ParaB string
}

type TestController struct {
	beego.Controller
}

func (c *TestController) Get() {
	var result TestResult
	val, err := c.Input()
	if err == nil {
		result.ParaA = val.Get("ParaA")
		result.ParaB = val.Get("ParaB")
		fmt.Println(result)
	} else {
		fmt.Println(err.Error())
	}
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *TestController) Post() {
	var result TestResult
	val, err := c.Input()
	if err == nil {
		result.ParaA = val.Get("ParaA")
		result.ParaB = val.Get("ParaB")
		fmt.Println(result)
	} else {
		fmt.Println(err.Error())
	}
	c.Data["json"] = result
	c.ServeJSON()
}
