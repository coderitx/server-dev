package routers

import (
	"blog-server/global"
	"context"
	"fmt"
	swaggerFiles "github.com/swaggo/files"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "blog-server/docs"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 定义公共前缀
	apiRouterGroup := router.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}
	// 查看系统信息
	routerGroupApp.SettingsRouter()
	// 上传图片相关
	routerGroupApp.ImagesRouter()
	// 广告相关
	routerGroupApp.AdvertRouter()
	// 菜单相关
	routerGroupApp.MenuRouter()
	// 用户相关
	routerGroupApp.UserRouter()
	// 标签相关
	routerGroupApp.TagRouter()
	// 消息相关
	routerGroupApp.MessageRouter()
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
