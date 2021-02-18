package api

import (
	"errors"

	"github.com/gin-gonic/gin"
)

const (
	CtxUserIdAndName = "user"
)

type UserMessage struct {
	UserId   int
	Username string
}

var ErrorUserNotLogin = errors.New("用户未登录")

// GetUserMessage 获取当前登录的用户ID和用户名
func GetUserMessage(c *gin.Context) (*UserMessage, error) {
	res, ok := c.Get(CtxUserIdAndName)
	if !ok {
		err := ErrorUserNotLogin
		return nil, err
	}
	userMessage := res.(*UserMessage)
	return userMessage, nil
}
