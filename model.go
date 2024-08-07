package zdpgo_gin_login

type GinLoginUser struct {
	Id       int    `json:"id"`
	Username string `json:"username" gorm:"column:username;unique"`
	Password string `json:"password,omitempty"`
}
