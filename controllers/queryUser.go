package controllers

import (
	"log"

	beego "github.com/beego/beego/v2/server/web"

	"edge-device-management-backend/models"
)

type QueryUserController struct {
	beego.Controller
}

func (c *QueryUserController) Get() {
	users := []models.User{}
	result := models.Database.Find(&users)
	if result.Error != nil {
		log.Printf("models.Database.Find() failed: %v \n", result.Error)
	}

	c.Data["json"] = users
	c.ServeJSON()
}
