package logic

import (
	"fmt"
	"mybluebell/dao/mysql"
	"mybluebell/models"
	"mybluebell/pkg/jwt"
	"mybluebell/pkg/snowflake"
)

func SignUp(p *models.ParamSignup) (err error) {
	// 1.判断用户存不存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// 2.生成UUID
	userID := snowflake.GenID()

	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	return mysql.InsertUser(user)

}

func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}

	// 2.查询用户信息，验证密码是否正确
	err = mysql.Login(user)
	if err != nil {
		return nil, err
	}
	fmt.Println("logic   Login  ", user.UserID, user.UserID, user.Password)
	// 生成jwt
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}
