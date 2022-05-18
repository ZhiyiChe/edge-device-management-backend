package controllers

import (
	"encoding/json"
	"fmt"

	"log"

	"edge-device-management-backend/models"
)

type UpdateUserController struct {
	MainController
}

func (c *UpdateUserController) Post() {
	user := models.User{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &user); err != nil {
		log.Printf("json.Unmarshal() failed: %v \n", err)
	}

	fmt.Println(user)

	tmpUser := models.User{}
	result := models.Database.Where("account = ?", user.Account).Take(&tmpUser)
	if result.Error != nil {
		// 如果数据库中没有该account，或发生其他错误
		log.Printf("models.Database.Where() failed: %v \n", result.Error)
		c.Data["json"] = &CommonResponse{
			Code: -1,
			Data: "failure",
		}
		goto END
	} else {
		user.Id = tmpUser.Id
		result = models.Database.Save(&user)
		if result.Error != nil {
			log.Printf("models.Database.Save() failed: %v \n", result.Error)
			c.Data["json"] = &CommonResponse{
				Code: -1,
				Data: "failure",
			}
			goto END
		}

		c.Data["json"] = &CommonResponse{
			Code: 0,
			Data: "success",
		}
	}

END:
	c.ServeJSON()
}
