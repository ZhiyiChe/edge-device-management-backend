package controllers

import (
	"log"

	"edge-device-management-backend/models"
)

type QueryUserController struct {
	MainController
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
