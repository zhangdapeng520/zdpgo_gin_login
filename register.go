package zdpgo_gin_login

import (
	gin "github.com/zhangdapeng520/zdpgo_gin"
	gorm "github.com/zhangdapeng520/zdpgo_gorm"
	password "github.com/zhangdapeng520/zdpgo_password"
)

// GetRegisterHandler 获取注册的路由
func GetRegisterHandler(db *gorm.DB, passwordSalt string) gin.HandlerFunc {
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

		db.Create(&GinLoginUser{
			Username: request.Username,
			Password: encryptedPassword,
		})
		c.JSON(200, nil)
	}
}
