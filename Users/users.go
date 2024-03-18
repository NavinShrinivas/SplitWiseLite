package Users

import (
	"github.com/gin-gonic/gin"
)

func UserApiRoutes(user_router *gin.RouterGroup) {
	user_router.GET("/register", UserRegister)
}
