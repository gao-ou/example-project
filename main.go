package main

import "example-project/cmd"

//
//type student struct {
//	Id      int
//	ClassId int
//	Name    string
//	Gender  string
//	Score   int
//}

func main() {
	cmd.Execute()
	////dsn := "root:12345678@tcp(127.0.0.1:3306)/chenshiweifuqindeshujuku?charset=utf8mb4&parseTime=True&loc=Local"
	//config.InitConfig("/Users/apple/Desktop/gaoou/woshinibaba/gorm.io/")
	//
	//username := viper.GetString("mysql.username")
	//password := viper.GetString("mysql.password")
	//addr := viper.GetString("mysql.addr")
	//dbname := viper.GetString("mysql.dbname")
	//db, err := db.ConnMysql(username, password, addr, dbname)
	//if err != nil {
	//	fmt.Println("连接数据库错误", err.Error())
	//	return
	//}
	//
	//type classes struct {
	//	Id   int
	//	Name string
	//}
	//
	//var tables []classes
	//err = db.Table("classes").Find(&tables).Error
	//if err != nil {
	//	fmt.Println("show tables 错误", err.Error())
	//	return
	//}
	//
	//fmt.Println(tables)
}
