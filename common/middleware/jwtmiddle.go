package middleware

import (
	"project/utils/config"
	"strings"

	"project/common/api"
	"project/pkg/jwt"
	"project/utils/app"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			app.ResponseError(c, app.CodeLoginExpire)
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == config.JwtConfig.Header) {
			app.ResponseError(c, app.CodeInvalidToken)
			c.Abort()
			return
		}
		//待添加用户缓存
		// 将当前请求的user_id信息保存到请求的上下文c上
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			c.Abort()
			return
		}
		r := new(api.UserMessage)
		r.UserId = mc.UserID
		r.Username = mc.Username
		c.Set(api.CtxUserIdAndName, r)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
