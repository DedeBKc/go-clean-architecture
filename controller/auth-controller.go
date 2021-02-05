package controller

import "github.com/gin-gonic/gin"

// AuthController untuk melakukan aksi apa
type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	// masukkan service yg dibutuhkan
}

// NewAuthController membuat instance dari AuthController
func NewAuthController() AuthController {
	return &authController{}
}

func (c *authController) Login(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"status": "Hello Login"})
}

func (c *authController) Register(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"status": "Hello Register"})
}
