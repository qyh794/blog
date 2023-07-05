package postgresql

import "errors"

var (
	ErrorUserExist          = errors.New("用户已存在")
	ErrorUserNotExist       = errors.New("用户不存在")
	ErrorInvalidPassword    = errors.New("密码错误")
	ErrorInvalidCommunityID = errors.New("社区不存在")
	ErrorDelete             = errors.New("删除失败")
)
