package model

import (
	"encoding/base64"
	"ginblog/utils/errmsg"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	// gorm.Model是⼀个包含了ID, CreatedAt, UpdatedAt, DeletedAt四个字段的Golang结构体
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required" label:"角色码"`
}

// CheckUserByUserName 查询用户是否存在
func CheckUserByUserName(username string) int {
	var user User
	db.Select("id").Where("username = ?", username).First(&user)
	if user.ID > 0 {
		return errmsg.ErrorUsernameUsed
	}
	return errmsg.SUCCESS
}

// CheckUpdateUserByUserName 更新查询
func CheckUpdateUserByUserName(id int, username string) int {
	var user User
	db.Select("id").Where("username = ?", username).First(&user)
	if user.ID == uint(id) {
		return errmsg.SUCCESS
	}
	if user.ID > 0 {
		return errmsg.ErrorUsernameUsed
	}
	return errmsg.SUCCESS
}

// CreateUser 新增用户
func CreateUser(data *User) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ErrorUserCreateFail
	}
	return errmsg.SUCCESS
}

// GetUser 查询单个用户
func GetUser(id int) (User, int) {
	var user User
	err := db.Limit(1).Where("ID = ?", id).First(&user).Error
	if err != nil {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCESS
}

// GetUsers 查询用户列表
func GetUsers(pageSize int, pageNum int, username string) ([]User, int64) {
	var users []User
	var total int64
	var err error
	// 分页偏移
	offset := (pageNum - 1) * pageSize
	if pageNum == -1 && pageSize == -1 {
		offset = -1
	}
	if username != "" {
		db.Where("username LIKE ?", "%"+username+"%").Limit(pageSize).Offset(offset).Find(&users)
		db.Model(&users).Where("username LIKE ?", "%"+username+"%").Count(&total)
		return users, total
	}
	db.Limit(pageSize).Offset(offset).Find(&users)
	db.Model(&users).Count(&total)
	if err != nil {
		return users, 0
	}
	return users, total
}

// EditUser 编辑用户
func EditUser(id uint, data *User) int {
	var user User
	var userMap = make(map[string]interface{})
	userMap["username"] = data.Username
	userMap["role"] = data.Role
	err := db.Model(&user).Where("id = ?", id).Updates(userMap).Error
	if err != nil {
		return errmsg.ErrorUserEditFail
	}
	return errmsg.SUCCESS
}

// ChangePassword 修改密码
func ChangePassword(id int, data *User) int {
	err := db.Select("password").Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ErrorUserDeleteFail
	}
	return errmsg.SUCCESS
}

// ScryptPassword 密码加密
func ScryptPassword(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{4, 128, 1, 32, 2, 8, 16, 64}
	// 官方自带 scrypt 库
	HashPassword, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	FinalPassword := base64.StdEncoding.EncodeToString(HashPassword)
	return FinalPassword
}

// BeforeSave 事务钩子 自动调用
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	u.Password = ScryptPassword(u.Password)
	return
}

// CheckLogin 登录验证
func CheckLogin(username string, password string) int {
	var user User
	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ErrorUserDntExist
	}
	if ScryptPassword(password) != user.Password {
		return errmsg.ErrorPasswordWrong
	}
	if user.Role != 1 {
		return errmsg.ErrorInsufficientPermissions
	}
	return errmsg.SUCCESS
}
