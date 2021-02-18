package models

import (
	"errors"
	"project/common/global"
	"project/utils"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SysUser struct {
	*BaseModel
	Username     string `json:"username"`
	Password     string `json:"password"`
	DeptId       int    `json:"dept_id"`        //部门id
	PostId       int    `json:"post_id"`        //
	RoleId       int    `json:"role_id"`        //
	NickName     string `json:"nick_name"`      //
	Phone        string `json:"phone"`          //
	Email        string `json:"email"`          //
	AvatarPath   string `json:"avatar_path"`    //头像路径
	Avatar       string `json:"avatar"`         //
	Sex          string `json:"sex"`            //
	Status       string `json:"status"`         //
	Remark       string `json:"remark"`         //
	Salt         string `json:"salt"`           //
	Gender       []byte `json:"gender"`         //性别（0为男默认，1为女）
	IsAdmin      []byte `json:"is_admin"`       //是否为admin账号
	Enabled      []byte `json:"enabled"`        //状态：1启用（默认）、0禁用
	PwdResetTime int64  `json:"pwd_reset_time"` //修改密码的时间
	CreateBy     int    `json:"create_by"`      //
	UpdateBy     int    `json:"update_by"`      //
}

type LoginM struct {
	UserName
	PassWord
}

type UserName struct {
	Username string `json:"username"`
}

type PassWord struct {
	// 密码
	Password string `json:"password"`
}

var (
	ErrorUserNotExist     = errors.New("用户不存在")
	ErrorInvalidPassword  = errors.New("用户名或密码错误")
	ErrorServerBusy       = errors.New("服务器繁忙")
	ErrorUserIsNotEnabled = errors.New("用户未激活")
)

func (SysUser) TableName() string {
	return "sys_user"
}

// Login 查询用户是否存在，并验证密码
func (u *SysUser) Login() error {
	oPassword := u.Password
	err := global.Eloquent.Table(u.TableName()).Where("username = ?", u.Username).First(u).Error
	if err == gorm.ErrRecordNotFound {
		zap.L().Error("用户不存在", zap.Error(err))
		return ErrorUserNotExist
	}
	if err != nil {
		zap.L().Error("服务器繁忙", zap.Error(err))
		return ErrorServerBusy
	}
	if u.Password != utils.EncodeMD5(oPassword) {
		zap.L().Error("user account or password is error")
		return ErrorInvalidPassword
	}
	if u.Enabled[0] == 0 {
		return ErrorUserIsNotEnabled
	}
	return nil
}
