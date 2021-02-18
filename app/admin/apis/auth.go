package apis

import (
	"github.com/gin-gonic/gin"

	"project/pkg/captcha"
	"project/utils"
	"project/utils/app"
)

func Captcha(c *gin.Context) {
	id, b64s, err := captcha.DriverMathFunc()
	utils.HasError(err, "验证码获取失败", 500)
	app.ResponseSuccess(c, gin.H{
		"code":    app.CodeSuccess,
		"message": app.CodeSuccess.Msg(),
		"img":     b64s,
		"uuid":    id,
	})
}
