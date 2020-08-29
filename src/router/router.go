package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jeferwang/WeWorkRobot/src/controller"
)

func Setup(r *gin.Engine) {
	r.GET("/callback", controller.Callback)
}
