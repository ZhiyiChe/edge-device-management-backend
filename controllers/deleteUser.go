package controllers

import (
	beego "github.com/beego/beego/v2/server/web"

	"log"

	"edge-device-management-backend/models"
)

type DeleteUserController struct {
	beego.Controller
}

func (c *DeleteUserController) Get() {
	values, err := c.Input()
	if err != nil {
		log.Printf("c.Input() failed: %v \n", err)
	}

	account := values.Get("account")

	user := models.User{}
	result := models.Database.Where("account = ?", account).Take(&user)
	if result.Error != nil {
		// 如果数据库中没有该account，或发生其他错误
		log.Printf("models.Database.Where() failed: %v \n", result.Error)
		c.Data["json"] = &CommonResponse{
			Code: -1,
			Data: "failure",
		}
		goto END
	} else {
		result = models.Database.Delete(&user)
		if result.Error != nil {
			log.Printf("models.Database.Delete() failed: %v \n", result.Error)
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
