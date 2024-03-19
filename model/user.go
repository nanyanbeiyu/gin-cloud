package model

import (
	"go_cloud/util"
	"go_cloud/util/errmsg"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName  string `gorm:"type:varchar(20)" json:"userName"`
	PassWord  string `gorm:"type:varchar(100)" json:"passWord"`
	OpenID    string `gorm:"type:varchar(255)"`
	ImagePath string `gorm:"type:varchar(255)"`
}

// 用户登陆校验
func CheckLogin(userName string, passWord string) (*User, int) {
	user, code := GetUserByName(userName)
	if code != errmsg.Success {
		return nil, code
	}
	if !util.CheckPassword(passWord, user.PassWord) {
		return user, errmsg.ERROR_USER_PASSWORD_ERROR
	}
	return user, errmsg.Success
}

// 添加用户
func Register(userName, passWord string) (int, *User) {
	// 查询用户名是否存在
	_, code := GetUserByName(userName)
	if code != errmsg.ERROR_USER_NOT_EXIST {
		return errmsg.ERROR_USER_ADD_EXIST, nil
	}
	// 生成openid
	openid := util.Snowflake()
	// 加密密码
	passWord, _ = util.HashPassword(passWord)
	user := User{
		UserName: userName,
		PassWord: passWord,
		OpenID:   openid,
	}
	err := db.Create(&user).Error
	if err != nil {
		return errmsg.ERROR_USER_ADD, nil
	}
	return errmsg.Success, &user
}

// 通过用户名查询用户
func GetUserByName(userName string) (*User, int) {
	var user User
	db.Find(&user, "user_name = ?", userName)
	if user.ID == 0 {
		return nil, errmsg.ERROR_USER_NOT_EXIST
	}
	return &user, errmsg.Success
}
