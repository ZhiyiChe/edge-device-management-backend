package controllers

import (
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"log"

	"edge-device-management-backend/models"
)

type AddUserController struct {
	MainController
}

func (c *AddUserController) Post() {
	user := models.User{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &user); err != nil {
		log.Printf("json.Unmarshal() failed: %v \n", err)
	}

	fmt.Println(user)

	tmpUser := models.User{}
	result := models.Database.Where("account = ?", user.Account).Take(&tmpUser)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// 如果数据库中没有该account，进行注册
		result = models.Database.Create(&user)
		if result.Error != nil {
			log.Printf("models.Database.Create() failed: %v \n", result.Error)
		}
		c.Data["json"] = &CommonResponse{
			Code: 0,
			Data: "success",
		}
	} else {
		// 如果数据库中已有该account，注册失败
		c.Data["json"] = &CommonResponse{
			Code: -1,
			Data: "failure",
		}
	}

	c.ServeJSON()
}
