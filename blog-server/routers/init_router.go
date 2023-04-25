package routers

import (
	"blog-server/global"
	"context"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitApiRouter() {
	if global.GlobalC.System.Env == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	// 处理路由分组
	InitRouterGroup(r)
	// 处理服务监听
	runServer(r)
}

func InitRouterGroup(router *gin.Engine) {
	// 定义公共前缀
	apiRouterGroup := router.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}
	// 查看系统信息
	routerGroupApp.SettingsRouter()
}

func runServer(router *gin.Engine) {
	srv := &http.Server{
		Addr:    fmt.Sprintf("%v:%d", global.GlobalC.System.Host, global.GlobalC.System.Port),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.S().Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	zap.S().Infoln("Listener Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.S().Fatal("Server Shutdown:", err)
	}
	zap.S().Infoln("Server exiting")
}
