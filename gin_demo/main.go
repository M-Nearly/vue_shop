package main

import (
	util "./util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		fmt.Println(method)
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// 处理请求
		c.Next()
	}
}

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.Use(Cors())
	loginGroup := r.Group("/api/private/v1")
	{
		loginGroup.POST("/login", func(c *gin.Context) {
			var login Login
			if err := c.ShouldBind(&login); err == nil {
				if login.Username == "admin" && login.Password == "123456" {
					token, _ := util.GenerateToken(login.Username, login.Password)
					c.JSON(http.StatusOK, gin.H{
						"status":   200,
						"username": login.Username,
						"password": login.Password,
						"token":    token,
					})
				} else {
					c.JSON(http.StatusOK, gin.H{
						"status":   201,
						"username": login.Username,
						"password": login.Password,
					})
				}

				//username := c.PostForm("username")
				//password := c.PostForm("password")

			}
		})

		r.Run(":8888")
	}
}
