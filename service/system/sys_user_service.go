package system

import (
	"errors"
	"web-server/global"
	"web-server/model/system"
	"web-server/model/system/request"
	"web-server/utils"

	uuid "github.com/satori/go.uuid"
)

type SysUserService struct{}

// Login 登录
func (s *SysUserService) Login(u *system.SysUser) (userInfo *system.SysUser, err error) {
	var user system.SysUser
	err = global.PalDB.Where("username = ?", u.Username).First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
	}
	return &user, err
}

// GetUserInfo 获取用户信息
func (s *SysUserService) GetUserInfo(uuid uuid.UUID) (userInfo *system.SysUser, err error) {
	var user system.SysUser
	err = global.PalDB.Select("id, uuid, username, nick_name,header_img, created_at, updated_at, updated_by").First(&user, "uuid = ?", uuid).Error
	return &user, err
}

// ChangePassword 修改密码
func (s *SysUserService) ChangePassword(ID uint, u request.ChangePassword) (err error) {
	var user system.SysUser
	err = global.PalDB.Select("password").First(&user, "id = ?", ID).Error
	if err != nil {
		return err
	}
	if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
		return errors.New("密码错误")
	}
	newPassword := utils.BcryptHash(u.NewPassword)
	err = global.PalDB.Model(&user).Where("id = ?", ID).Update("password", newPassword).Error
	return err
}
