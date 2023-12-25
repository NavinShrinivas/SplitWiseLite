package main

import (
	"github.com/gin-gonic/gin"
   "SplitWiseLite/users"
)

func V1_Routes(v1_router *gin.RouterGroup) {
	users_grp := v1_router.Group("/users")
   {
      users.UserApiRoutes(users_grp)
   }

}
