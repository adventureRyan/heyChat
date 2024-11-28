package service

import (
	"fmt"
	"heyChat/models"
	"heyChat/utils"
	"math/rand"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// GetUserList
// @Summary 用户列表
// @Tags 用户模块
// @Success 200 {string} json{"code", "message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	c.JSON(200, gin.H{
		"message": data,
	})
}

// CreateUser
// @Summary 新增用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")
	salt := fmt.Sprintf("%06d", rand.Int31())
	data := models.FindUserByName(user.Name)
	if data.Name != "" {
		c.JSON(-1, gin.H{
			"message": "用户名已经注册",
		})
		return
	}
	if password != repassword {
		c.JSON(-1, gin.H{
			"message": "两次密码不一致",
		})
		return
	}
	// user.PassWord = password
	user.PassWord = utils.MakePassword(password, salt)
	user.Salt = salt
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"message": "新增用户成功",
	})
}

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @param id query string false "用户名"
// @Success 200 {string} json{"code","message"}
// @Router /user/DeleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(200, gin.H{
		"message": "删除用户成功",
	})
}

// UpdateUser
// @Summary 修改用户
// @Tags 用户模块
// @param id formData string false "id"
// @param name formData string false "用户名"
// @param password formData string false "密码"
// @param phone formData string false "电话"
// @param email formData string false "邮箱"
// @Success 200 {string} json{"code","message"}
// @Router /user/UpdateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"message": "修改参数不匹配",
		})
	} else {
		models.UpdateUser(user)
		c.JSON(200, gin.H{
			"message": "修改用户成功",
		})
	}

}

// FindUserByName
// @Summary 登录
// @Tags 用户模块
// @param name query string false "name"
// @param password query string false "password"
// @Success 200 {string} json{"code","message"}
// @Router /user/FindUserByName [post]
func FindUserByName(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	user.PassWord = c.Query("password")
	res_user := models.FindUserByName(user.Name)
	if res_user.Name == "" {
		c.JSON(200, gin.H{
			"message": "没有这个用户",
		})
		return
	}
	cal_pass := utils.MakePassword(user.PassWord, res_user.Salt)
	if cal_pass != res_user.PassWord {
		c.JSON(200, gin.H{
			"message": "密码错误",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "登录成功",
	})
	return
}
