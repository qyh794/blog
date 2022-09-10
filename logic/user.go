package logic

import (
	"blog/dao/mysql"
	"blog/models"
	"blog/pkg/jwt"
	"blog/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 判断用户是否已存在
	if err := mysql.UserIsExist(p.Username); err != nil {
		return err
	}
	// 如果存在返回错误
	// 不存在则在mysql创建一个用户
	userID := snowflake.GenID()
	user := &models.User{
		UserID: userID,
		Username: p.Username,
		Password: p.Password,
	}

	
	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) (token string, err error) {
	// 创建user示例，将获取到的参数放入示例
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 将user示例传入mysql进行查询
	if err := mysql.Login(user); err != nil {
		return "", err
	}
	// 查询成功需要返回的东西：1.token
	return jwt.GenToken(user.UserID, user.Username)
}