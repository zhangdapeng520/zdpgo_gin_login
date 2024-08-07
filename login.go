package zdpgo_gin_login

import (
	gin "github.com/zhangdapeng520/zdpgo_gin"
	gorm "github.com/zhangdapeng520/zdpgo_gorm"
	password "github.com/zhangdapeng520/zdpgo_password"
)

// GetLoginHandler 获取登录的路由
func GetLoginHandler(db *gorm.DB, jwtKey, passwordSalt string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request userRequest
		err := c.ShouldBindJSON(&request)
		if err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}
		encryptedPassword, err := password.Sha256EncryptString(request.Password, passwordSalt)
		if err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}

		var user GinLoginUser
		db.Find(&GinLoginUser{
			Username: request.Username,
			Password: encryptedPassword,
		}).First(&user)

		// 生成Token
		token, err := GetToken(user.Id, user.Username, jwtKey)
		if err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"username": user.Username,
			"token":    token,
		})
	}
}
