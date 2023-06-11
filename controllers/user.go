package controllers

import (
	"blog/dao/mysql"
	"blog/logic"
	"blog/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("signup with invalid param", zap.Error(err))
		err, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvaildParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvaildParam, removeTopStruct(err.Translate(trans)))
		return
	}
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp(p) failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeSeverBusy)
	}
	ResponseSuccess(c, CodeSuccess)
}

func LoginHandler(c *gin.Context) {
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("login with invalid param", zap.Error(err))
		err, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvaildParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvaildParam, removeTopStruct(err.Translate(trans)))
		return
	}
	token, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login(p) failed, err:", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvaildPassword)
		return
	}
	ResponseSuccess(c, fmt.Sprintf("用户认证为:%s", token))

}
