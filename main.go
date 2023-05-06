package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	"path/filepath"
)

func main() {
	r := gin.Default()

	// 加载模板文件
	r.LoadHTMLGlob(filepath.Join("templates", "*.gohtml"))

	// 显示初始化安装页面
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "install.gohtml", gin.H{})
	})

	// 处理安装请求
	r.POST("/install", func(c *gin.Context) {
		ipAddress := c.PostForm("ipAddress")

		// 执行安装命令
		cmd := exec.Command("ping", "-c", "3", ipAddress)
		err := cmd.Run()

		if err != nil {
			// 安装失败
			c.JSON(http.StatusOK, gin.H{
				"status": "error",
				"msg":    "安装失败",
			})
		} else {
			// 安装成功
			c.JSON(http.StatusOK, gin.H{
				"status": "success",
				"msg":    "初始化安装完成",
			})
		}
	})

	// 启动HTTP服务
	if err := r.Run(":8080"); err != nil {
		fmt.Println("HTTP server启动失败：", err)
	}
}
