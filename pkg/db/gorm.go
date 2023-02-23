package db

import (
	"fmt"                  //打印输出用的包
	"gorm.io/driver/mysql" //gorm的my sql
	"gorm.io/gorm"         //gorm自己的包
)

// ConnMysql 连接数据库 mysql
func ConnMysql(username string, password string, addr string, dbname string) (*gorm.DB, error) {

	// DB  database  数据库
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	// "用户名:密码@tcp(IP地址:端口号)/数据库名称?charset=utf8mb4&parseTime=True&loc=Local"
	var dsn = username + ":" + password + "@tcp(" + addr + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "root:12345678@tcp(127.0.0.1:3306)/chenshiweifuqindeshujuku?charset=utf8mb4&parseTime=True&loc=Local"
	// db是gorm gormDB结构体的变量，err 是gorm.Open返回的错误
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) //gorm包名  链接mysql 返回一个连接器 把连接器给gorm.Open 利用gorm打开连接器
	//连接后有两个返回值 db和err db是一个结构体 err是一个错误
	// 判断是否有错误
	if err != nil {
		fmt.Printf("连接数据库错误:%s\n", err.Error())
		return nil, err
	}

	return db, nil
}
