package dto

//UserLoginDto 登录请求参数
type UserLoginDto struct {
	Username string `json:"username" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"` // 密码
	Code     string `json:"code" binding:"required"`     // 验证码
	UuId     string `json:"uuid" binding:"required"`     // 验证码id
}
