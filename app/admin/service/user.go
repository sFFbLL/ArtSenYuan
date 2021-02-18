package service

import (
	"project/pkg/jwt"

	"project/app/admin/models"
	"project/app/admin/models/dto"
)

type User struct {
}

// Login 返回json web token
func (u *User) Login(p *dto.UserLoginDto) (token string, err error) {
	user := new(models.SysUser)
	user.Username = p.Username
	user.Password = p.Password

	if err = user.Login(); err != nil {
		return
	}

	token, err = jwt.GenToken(user.ID, user.Username)
	if err != nil {
		return
	}
	return
}
