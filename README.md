# zdpgo_gin_login

适配gin框架的登录注册功能组件，通过本框架轻松拥有登录注册相关的API接口

## 安装

```bash
go get github.com/zhangdapeng520/zdpgo_gin_login
```

## 使用

### 基本用法

```go
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
```

## 版本

### v0.1.0

- 增加注册和登录两个接口的路由获取方式

### v0.1.1

- 登录接口返回用户名