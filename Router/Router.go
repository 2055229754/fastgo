package Router

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/2055229754/fastgo/Config"
	"github.com/2055229754/fastgo/Controller"
	"github.com/2055229754/fastgo/Request"
	"github.com/julienschmidt/httprouter"
)

var HTTPMETHOD = map[string]bool{
	"GET":       true,
	"POST":      true,
	"PUT":       true,
	"DELETE":    true,
	"PATCH":     true,
	"OPTIONS":   true,
	"HEAD":      true,
	"TRACE":     true,
	"CONNECT":   true,
	"MKCOL":     true,
	"COPY":      true,
	"MOVE":      true,
	"PROPFIND":  true,
	"PROPPATCH": true,
	"LOCK":      true,
	"UNLOCK":    true,
}

type RouteContainer struct {
	Routes    []*RouteInfo
	GroupName string
}
type RouteInfo struct {
	Pattern    string
	Controller Controller.ControllerInterface
	Method     string
	Action     string
	Group      string
}

var GroupName string
var Route RouteContainer

func New() *RouteContainer {

	return &Route
}

func (rc *RouteContainer) Add(pattern string, c Controller.ControllerInterface, methods string, action string) *RouteInfo {
	var routeinfo RouteInfo
	routeinfo.Pattern = pattern
	routeinfo.Controller = c
	routeinfo.Method = methods
	routeinfo.Action = action
	routeinfo.Group = GroupName
	rc.Routes = append(Route.Routes, &routeinfo)
	return &routeinfo
}

func (rc *RouteContainer) Group(pattern string, child ...*RouteInfo) {
	rc.GroupName = pattern
	for _, v := range child {
		v.Group = pattern
	}
}

func (rc *RouteContainer) Run() {
	router := httprouter.New()
	for _, v := range rc.Routes {
		pattern := v.Pattern
		method := v.Method
		// controller := v.Controller
		group := v.Group
		// action := v.Action
		if group != "" {
			pattern = "/" + group + pattern
		}
		method = strings.ToUpper(method)
		if !HTTPMETHOD[method] {
			panic(method + "禁止使用")
		}
		switch method {
		case "GET":
			router.GET(pattern, RouteHandle)
		case "POST":
			router.POST(pattern, RouteHandle)
		}

	}
	fmt.Println("端口", ":"+strconv.Itoa(Config.AppConf.Http_port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(Config.AppConf.Http_port), router))
}
func RouteHandle(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//根据请求执行具体控制器方法
	uri := r.URL.Path
	method := r.Method
	ctx := Request.Request{
		Req:    r,
		Res:    w,
		Params: params,
	}
	for _, v := range Route.Routes {
		var fullurl string
		if v.Group == "" {
			fullurl = v.Pattern
		} else {
			fullurl = "/" + v.Group + v.Pattern
		}

		if fullurl == uri && method == strings.ToUpper(v.Method) {
			executeHandler(&ctx, v.Controller, v.Action)
		}
	}
}

func executeHandler(ctx *Request.Request, c Controller.ControllerInterface, action string) {
	var execController Controller.ControllerInterface
	reflectType := reflect.TypeOf(c)
	reflectVal := reflect.ValueOf(c)
	var ok bool
	execController, ok = reflectVal.Interface().(Controller.ControllerInterface)
	if !ok {
		panic("controller is not ControllerInterface")
	}
	controllerName := strings.Split(reflectType.String(), ".")
	execController.Construct(ctx, controllerName[1], action)
	fn := reflectVal.MethodByName(action)
	if !fn.IsValid() {

	} else {
		fn.Call(nil)
	}

}
