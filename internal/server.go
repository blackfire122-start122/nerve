package internal

import (
	"github.com/gin-gonic/gin"
)

func SetRouters(router *gin.Engine) {
	router.GET("/", Home)
	router.GET("/register", Register)
	router.GET("/login", LogIn)
	router.GET("/profile", Profile)
	router.POST("/api/user/register", RegisterUser)
	router.POST("/api/user/login", LoginUser)
	router.POST("/api/user/changeData", ChangeUser)
	router.GET("/api/user/getUser", GetUser)
	router.GET("/api/user/tasks", UserTasks)
}
