package internal

import (
	//. "Nerve/pkg/logging"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(c *gin.Context) {
	//ErrorLogger.Println(err.Error())
	c.JSON(http.StatusOK, nil)
}
