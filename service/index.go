package service

import (
	"fmt"
	"html/template"
	"net/http"

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
