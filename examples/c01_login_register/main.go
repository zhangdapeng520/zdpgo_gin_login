package main

import (
	gin "github.com/zhangdapeng520/zdpgo_gin"
	ginLogin "github.com/zhangdapeng520/zdpgo_gin_login"
	gorm "github.com/zhangdapeng520/zdpgo_gorm"
	_ "github.com/zhangdapeng520/zdpgo_mysql"
)

var (
	jwtKey       = "zhangdapeng.com"
	passwordSalt = "zhangdapeng.com"
)

func main() {
	db, err := gorm.Open(
		"mysql",
		"root:root@tcp(127.0.0.1:3306)/blog?charset=utf8",
	)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&ginLogin.GinLoginUser{})
	defer db.Close()

	r := gin.Default()

	r.POST("/user/register/", ginLogin.GetRegisterHandler(db, passwordSalt))
	r.POST("/user/login/", ginLogin.GetLoginHandler(db, jwtKey, passwordSalt))

	r.Run(":8888")
}
