package mysql

import (
	"blog/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"

	"go.uber.org/zap"
)

const secret = "abc.com"

// UserIsExist 判断用户是否存在
func UserIsExist(username string) (err error) {
	// 查表
	sqlStr := "select count(user_id) from user where username=?"
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

// InsertUser 插入用户
func InsertUser(user *models.User) (err error) {
	user.Password = encryptPassword(user.Password)
	sqlStr := "insert into user(user_id, username, password) values(?,?,?)"
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

// GetUserByID 通过id获取用户信息
func GetUserByID(userID int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := "select username from user where user_id = ?"
	err = db.Get(user, sqlStr, userID)
	return
}


// Login 用户登录
func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := "select user_id, username, password from user where username = ?"
	
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		zap.L().Error("111", zap.Error(err))
		return err
	}
	if user.Password != encryptPassword(oPassword) {
		return ErrorInvalidPassword
	}
	return
}

// encryptPassword 密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}