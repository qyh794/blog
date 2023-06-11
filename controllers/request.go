package controllers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)
var ErrorUserNotLogin = errors.New("用户未登录")

func getCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get("userID")
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	// 用类型断言转换
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

func getPageInfo(c  *gin.Context) (int, int) {
	pageStr := c.Query("page")
	sizeStr := c.Query("size")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		size = 10
	}
	return page, size
}