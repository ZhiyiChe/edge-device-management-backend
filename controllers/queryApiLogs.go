package controllers

import (
	"log"
	"strconv"

	"github.com/beego/beego/v2/core/logs"

	"edge-device-management-backend/models"
)

type QueryApiLogsController struct {
	MainController
}

func (c *QueryApiLogsController) Get() {
	values, err := c.Input()
	if err != nil {
		log.Printf("c.Input() failed: %v \n", err)
	}

	num, err := strconv.Atoi(values.Get("num"))
	if err != nil {
		log.Printf("strconv.Atoi() failed: %v \n", err)
	}

	logs.Debug("my book is bought in the year of ", num)

	apilogs := []models.Apilog{}
	result := models.Database.Limit(num).Order("id desc").Find(&apilogs)
	if result.Error != nil {
		log.Printf("models.Database.Limit().Order().Find() failed: %v \n", result.Error)
	}

	c.Data["json"] = &CommonResponse{
		Code: 0,
		Data: apilogs,
	}

	c.ServeJSON()
}
