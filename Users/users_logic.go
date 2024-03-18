package Users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "registered user successfully",
	})
}
