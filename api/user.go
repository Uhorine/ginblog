package api

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"ginblog/utils/validate"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var code int

// AddUser 添加用户
func AddUser(c *gin.Context) {
	// todo 添加用户
	var data model.User
	_ = c.ShouldBindJSON(&data)
	// 传入数据验证
	msg, code := validate.Validate(&data)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": msg,
		})
		return
	}
	code = model.CheckUserByUserName(data.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetUserInfo 查询单个用户
func GetUserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var maps = make(map[string]any)
	data, code := model.GetUser(id)
	maps["username"] = data.Username
	maps["role"] = data.Role
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   1,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetUserList 查询用户列表
func GetUserList(c *gin.Context) {
	// todo 获得所有用户
	// 前端传经来的分页大小和页码
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	username := c.Query("username")
	if pageSize == 0 {
		// 通过 -1 消除 Limit 条件
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, total := model.GetUsers(pageSize, pageNum, username)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// EditUser 编辑用户
func EditUser(c *gin.Context) {
	// todo 编辑用户
	id, _ := strconv.Atoi(c.Param("id"))
	var data model.User
	_ = c.ShouldBindJSON(&data)
	// 根据用户名判断是否有重名
	code = model.CheckUpdateUserByUserName(id, data.Username)
	if code == errmsg.SUCCESS {
		model.EditUser(uint(id), &data)
	}
	if code == errmsg.ErrorUsernameUsed {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// ChangePassword 修改密码
func ChangePassword(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code := model.ChangePassword(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// DeleteUser 删除用户 ==> 软删除 被软删除的数据是不会被查出来的
func DeleteUser(c *gin.Context) {
	// todo 删除用户
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
