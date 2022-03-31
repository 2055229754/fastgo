package fast

import (
	"github.com/2055229754/fastgo/Config"
	"github.com/2055229754/fastgo/Router"
)

var (
	APP_CONF_PATH = "conf/app.json"
)

type App struct {
	Config *Config.ConfigContainer
	Route  *Router.RouteContainer
}

//初始化项目APP
func New() *App {
	app := &App{}
	//注册配置
	app.Config = registConfig()
	//注册路由
	app.Route = registRoute()
	return app
}

func registConfig() *Config.ConfigContainer {
	c := Config.NewConfigContainer()
	return c
}
func registRoute() *Router.RouteContainer {
	r := &Router.Route
	return r
}

func (fa *App) Start() {
	fa.Route.Run()
}
