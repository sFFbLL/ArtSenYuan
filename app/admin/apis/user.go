package apis

import (
	"errors"
	"project/app/admin/models"
	"project/app/admin/models/dto"
	"project/app/admin/service"
	"project/utils/app"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func LoginHandler(c *gin.Context) {
	// 1.获取参数 校验参数
	p := new(dto.UserLoginDto)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误， 直接返回响应
		zap.L().Error("Login failed", zap.String("username", p.Username), zap.Error(err))
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	/*
	   	// 2.业务逻辑处理
	   	//TODO 方便postman测试 (模拟前端数据)
	   	//p.Password, _ = utils.RsaPubEncode(p.Password)
	   	//value, err := utils.RsaPriDecode(p.Password)
	   	//if err != nil {
	   	//	zap.L().Error("ras decode fail", zap.Error(err))
	   	//	app.ResponseError(c, app.CodeLoginFailResCode)
	   	//	return
	   	//}
	   	//p.Password = value
	   私钥 公钥 解密环节
	*/
	u := new(service.User)
	data, err := u.Login(p)
	if err != nil {
		c.Error(err)
		zap.L().Error("get login user info message failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, models.ErrorInvalidPassword) || errors.Is(err, models.ErrorUserNotExist) {
			app.ResponseError(c, app.CodeLoginFailResCode)
			return
		} else if errors.Is(err, models.ErrorUserIsNotEnabled) {
			app.ResponseError(c, app.CodeUserIsNotEnabled)
		}
		app.ResponseError(c, app.CodeSeverError)
		return
	}

	// 3.返回响应
	app.ResponseSuccess(c, data)
}
