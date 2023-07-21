package cmd

import (
	"example-project/pkg/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

var (
	apiCmd = cobra.Command{
		Use: "api",
		Run: func(cmd *cobra.Command, args []string) {
			// 这里 运行gin
			// TODO:: @gaoou
			r := gin.Default()
			r.GET("/project", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "example-project",
				})
			})
			//读取数据库配置，并赋值给变量
			username := viper.GetString("mysql.username")
			password := viper.GetString("mysql.password")
			addr := viper.GetString("mysql.addr")
			dbname := viper.GetString("mysql.dbname")
			// 初始化 db
			db, err := db.ConnMysql(username, password, addr, dbname)
			if err != nil {
				fmt.Println("连接数据库错误", err.Error())
				return
			}

			a := func() {}
			b := func(fn func()) {
				fn()
			}

			// 运行a函数
			a()

			// 当做函数参数传还给b
			b(a)

			r.GET("/api/classes", func(c *gin.Context) {

				type classes struct {
					Id   int
					Name string
				}

				var tables []classes
				err = db.Table("classes").Find(&tables).Error
				if err != nil {
					fmt.Println("show tables 错误", err.Error())
					return
				}
				c.JSON(http.StatusOK, tables)

			})

			err = r.Run()
			if err != nil {
				return
			} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
			// https://gin-gonic.com/zh-cn/docs/
			fmt.Println("api", time.Now())

		}}
)
