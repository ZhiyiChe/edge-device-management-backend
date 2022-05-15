package main

import (
	_ "edge-device-management-backend/routers"

	"github.com/beego/beego/v2/server/web"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
)

func main() {
	web.BConfig.WebConfig.Session.SessionOn = true // 使用session

	// 解决跨域问题
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		// AllowAllOrigins: true,
		AllowOrigins: []string{"http://10.*.*.*:*", "http://localhost:*", "http://127.0.0.1:*"},
		//AllowOrigins:      []string{"https://192.168.0.102"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "Access-Control-Allow-Credentials", "withCredentials", "token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "Access-Control-Allow-Credentials"},
		AllowCredentials: true,
	}))

	beego.Run()
}
