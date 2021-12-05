package controller

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/models"
	"errors"
	"strconv"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// SignupHandler 处理注册请求的函数
// @Summary 处理注册请求的函数
// @Description 处理注册请求的函数
// @Tags 注册
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param object query models.ParamSignUp false "注册信息"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseSuccess
// @Router /posts2 [get]
func SignupHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	//var p models.ParamSignUp
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回相应
		zap.L().Error("Signup with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 2. 业务处理
	if err := logic.Signup(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, nil)
}

// LoginHandler 登陆请求的函数
// @Summary 处理登陆请求的函数
// @Description 处理登陆请求的函数
// @Tags 登陆
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param object query models.ParamLogin false "登陆信息"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseSuccess
// @Router /posts2 [get]
func LoginHandler(c *gin.Context) {
	// 1. 获取请求参数及参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回相应
		zap.L().Error("Login with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 2. 业务逻辑处理
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}

	// 3. 返回相应
	ResponseSuccess(c, gin.H{
		"user_id":   strconv.FormatInt(user.UserID, 10), //id值大于1<<53-1，int64类型的最大值是1<<64-1
		"user_name": user.Username,
		"token":     user.Token,
	})
}
