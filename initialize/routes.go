package initialize

import (
	"github.com/gin-gonic/gin"
	"watchAlert/globals"
	"watchAlert/routers"
	"watchAlert/routers/v1"
)

func InitRoute() {

	globals.Logger.Sugar().Info("服务启动")

	ginEngine := gin.Default()
	allRouter(ginEngine)

	err := ginEngine.Run(":" + globals.Config.Server.Port)
	if err != nil {
		globals.Logger.Sugar().Error("服务启动失败:", err)
		return
	}
}

func allRouter(engine *gin.Engine) {

	routers.HealthCheck(engine)
	v1.AlertEventMsg(engine)
	v1.Web(engine)

}
