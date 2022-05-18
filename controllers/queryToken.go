package controllers

import (
	"log"
	"os"
)

var token []byte
var lastModTimeNs int64

func init() {
	file, err := os.Open("/var/lib/rancher/k3s/server/node-token")
	if err != nil {
		log.Printf("os.Open() failed: %v \n", err)
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		log.Printf("file.Stat() failed: %v \n", err)
	}

	token = make([]byte, fileinfo.Size())
	_, err = file.Read(token)
	if err != nil {
		log.Printf("file.Read() failed: %v \n", err)
	}
	lastModTimeNs = fileinfo.ModTime().UnixNano()
}

type QueryTokenController struct {
	MainController
}

func (c *QueryTokenController) Get() {
	file, err := os.Open("/var/lib/rancher/k3s/server/node-token")
	if err != nil {
		log.Printf("os.Open() failed: %v \n", err)
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		log.Printf("file.Stat() failed: %v \n", err)
	}

	if fileinfo.ModTime().UnixNano() > lastModTimeNs {
		_, err = file.Read(token)
		if err != nil {
			log.Printf("file.Read() failed: %v \n", err)
		}
		lastModTimeNs = fileinfo.ModTime().UnixNano()
	}

	c.Ctx.WriteString(string(token))
}
