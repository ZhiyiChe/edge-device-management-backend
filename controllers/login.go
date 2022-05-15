package controllers

import (
	"encoding/json"
	"errors"

	beego "github.com/beego/beego/v2/server/web"
	"gorm.io/gorm"

	"log"

	"edge-device-management-backend/models"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Post() {
	user := models.User{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &user); err != nil {
		log.Printf("json.Unmarshal() failed: %v \n", err)
	}

	tmpUser := models.User{}
	result := models.Database.Where("account = ?", user.Account).Take(&tmpUser)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) || tmpUser.Password != user.Password {
		// 如果数据库中没有该account，或者登录密码错误，登录失败
		c.Data["json"] = &CommonResponse{
			Code: -1,
			Data: "failure",
		}
	} else {
		// token鉴权
		// tokenStr := c.Ctx.Request.Header["Token"][0]
		// fmt.Println(tokenStr)
		// info, err := ValidateToken(tokenStr)
		// if err != nil {
		// 	log.Printf("ValidateToken() failed: %v \n", err)
		// }
		// fmt.Println(info)

		token, err := GenerateToken(&tmpUser, 0)
		if err != nil {
			log.Printf("GenerateToken() failed: %v \n", err)
		}

		c.Data["json"] = &CommonResponse{
			Code: 0,
			Data: token,
		}
	}

	c.ServeJSON()
}
