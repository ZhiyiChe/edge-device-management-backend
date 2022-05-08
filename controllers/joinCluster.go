package controllers

import (
	"log"
	"os"

	beego "github.com/beego/beego/v2/server/web"
)

type JoinClusterController struct {
	beego.Controller
}

func (c *JoinClusterController) Get() {
	file, err := os.Open("/home/ubuntu/edge-device-management-backend/install/joinCluster.sh")
	if err != nil {
		log.Printf("os.Open() failed: %v \n", err)
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		log.Printf("file.Stat() failed: %v \n", err)
	}

	data := make([]byte, fileinfo.Size())
	_, err = file.Read(data)
	if err != nil {
		log.Printf("file.Read() failed: %v \n", err)
	}

	c.Ctx.WriteString(string(data))
}
