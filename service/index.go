package service

import (
	"fmt"
	"heyChat/models"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetIndex
// @Tags 首页
// @Success 200 {string} welcome
// @Router /index [get]
func GetIndex(c *gin.Context) {
	ind, err := template.ParseFiles("index1.html")
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error loading template: %v", err))
		return
	}
	err = ind.Execute(c.Writer, "index1")
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error executing template: %v", err))
	}
	// c.JSON(200, gin.H{
	// 	"message": "welcome!!",
	// })
}

func ToRegister(c *gin.Context) {
	ind, err := template.ParseFiles("views/user/register.html")
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error loading template: %v", err))
		return
	}
	err = ind.Execute(c.Writer, "register")
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error executing template: %v", err))
	}
	// c.JSON(200, gin.H{
	// 	"message": "welcome!!",
	// })
}

func ToChat(c *gin.Context) {
	ind, err := template.ParseFiles("views/chat/index.html",
		"views/chat/head.html",
		"views/chat/foot.html",
		"views/chat/tabmenu.html",
		"views/chat/concat.html",
		"views/chat/group.html",
		"views/chat/profile.html",
		"views/chat/main.html",
		"views/chat/userinfo.html",
		"views/chat/createcom.html")
	if err != nil {
		// 模板解析失败时记录错误并返回500状态码
		c.String(http.StatusInternalServerError, "Error parsing templates: %v", err)
		return
	}

	// if err != nil {
	// 	panic(err)
	// }
	userId, _ := strconv.Atoi(c.Query("userId"))
	if err != nil {
		// 用户ID解析失败
		c.String(http.StatusBadRequest, "Invalid userId: %v", err)
		return
	}
	token := c.Query("token")
	user := models.UserBasic{}
	user.ID = uint(userId)
	user.Identity = token
	err = ind.Execute(c.Writer, user)
	if err != nil {
		// 模板执行失败时记录错误并返回500状态码
		c.String(http.StatusInternalServerError, "Error executing template: %v", err)
	}
	// c.JSON(200, gin.H{
	// 	"message": "ToChat Page!!",
	// })
}
