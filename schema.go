package zdpgo_gin_login

type userRequest struct {
	Username string `json:"username" binding:"required,gte=3,lte=36"`
	Password string `json:"password" binding:"required,gte=6,lte=128"`
}
