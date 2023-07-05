package postgresql

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
	sqlStr := `
SELECT 
    "user_id" 
FROM 
    "user"
WHERE 
    username = $1`
	var userID int64
	if err := db.Get(&userID, sqlStr, username); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}
	if userID != 0 {
		return ErrorUserExist
	}
	return
}

// InsertUser 插入用户
func InsertUser(user *models.User) (err error) {
	user.Password = encryptPassword(user.Password)
	sqlStr := `
insert into
    "user" (user_id, username, password) 
values 
    ($1,$2,$3)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

// GetUserByID 通过id获取用户信息
func GetUserByID(userID int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `
select 
    "username" 
from 
    "user" 
where 
    "user_id" = $1`
	err = db.Get(user, sqlStr, userID)
	return
}

// Login 用户登录
func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := `
select 
    "user_id", "username", "password" 
from 
    "user"
where 
    "username" = $1`

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
