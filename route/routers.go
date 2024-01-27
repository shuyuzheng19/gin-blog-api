package route

import (
	"common-web-framework/config"
	"common-web-framework/helper"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"net/http"
	"time"
)

type RouterSetup struct {
	apiGroup *gin.RouterGroup
}

type Router struct {
	//服务器监听地址
	addr string
	//gin服务
	service *gin.Engine
	//gin路由分组
	apiGroup *gin.RouterGroup
	//读取超时时间
	readTimeOut time.Duration
	//写入超时时间
	writeTimeOut time.Duration
}

// RunServer 运行Server
func (r Router) RunServer() {

	var server = &http.Server{
		Addr:         r.addr,
		Handler:      r.service,
		ReadTimeout:  r.readTimeOut,
		WriteTimeout: r.writeTimeOut,
	}

	if config.CONFIG.Server.Cron {
		cronJob.Start()
	}

	var err = server.ListenAndServe()

	helper.ErrorPanicAndMessage(err, "服务器启动失败")
}

func (r *Router) SetupController() *Router {
	var controllers = &RouterSetup{apiGroup: r.apiGroup}
	controllers.LoadUserController("user").
		LoadFileController("file").
		LoadTagController("tags").
		LoadCategoryController("category").
		LoadBlogController("blog").
		LoadTopicController("topics").
		LoadToolsController("tools").
		LoadDataBaseController("database")
	return r
}

func (r *Router) AddMiddles(f ...gin.HandlerFunc) *Router {
	r.apiGroup.Use(f...)
	return r
}

var cronJob *cron.Cron

func NewRouter(serverConfig config.ServerConfig) *Router {

	if serverConfig.Cron {
		cronJob = cron.New()
	}

	var service *gin.Engine

	if serverConfig.Release {
		gin.SetMode(gin.ReleaseMode)
		service = gin.New()
	} else {
		service = gin.Default()
	}

	var corsConfig = config.CONFIG.Cors

	if corsConfig.Enable {
		var cConfig = cors.Config{
			AllowMethods:     corsConfig.AllowMethods,
			AllowHeaders:     corsConfig.AllowHeaders,
			AllowCredentials: corsConfig.AllowCredentials,
			ExposeHeaders:    corsConfig.ExposeHeaders,
		}

		if corsConfig.AllOrigins {
			cConfig.AllowAllOrigins = true
		} else {
			cConfig.AllowOrigins = corsConfig.AllowOrigins
		}
		//加载cors中间件
		service.Use(cors.New(cConfig))
	}

	var group = service.Group(serverConfig.ApiPrefix)

	return &Router{
		addr:         serverConfig.Addr,
		service:      service,
		apiGroup:     group,
		writeTimeOut: time.Second * time.Duration(serverConfig.WriteTimeOut),
		readTimeOut:  time.Second * time.Duration(serverConfig.ReadTimeOut),
	}
}
