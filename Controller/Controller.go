package Controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/2055229754/fastgo/Debug"
	"github.com/2055229754/fastgo/Request"
)

type Controller struct {
	// context data
	Ctx  *Request.Request
	Data map[interface{}]interface{}

	// route controller info
	controllerName string
	actionName     string
	methodMapping  map[string]func() //method:routertree
	AppController  interface{}

	// template data
	TplName        string
	ViewPath       string
	Layout         string
	LayoutSections map[string]string // the key is the section name and the value is the template name
	EnableRender   bool
}

type ControllerInterface interface {
	Construct(ctx *Request.Request, controllerName, actionName string)
	BeforAction()
	Get()
	Post()
	Delete()
	Put()
	Head()
	Patch()
	Options()
	AfaterAction()
	URLMapping()
}

// Init generates default values of controller operations.
func (c *Controller) Construct(ctx *Request.Request, controllerName, actionName string) {
	rootPath, _ := filepath.Abs("")
	c.ViewPath = fmt.Sprintf("%s\\app\\views\\%s\\", rootPath, strings.ToLower(controllerName))
	c.Layout = ""
	c.TplName = fmt.Sprintf("%s.html", strings.ToLower(actionName))
	c.controllerName = controllerName
	c.actionName = actionName
	c.Ctx = ctx
	c.EnableRender = true
	c.Data = map[interface{}]interface{}{}
	c.methodMapping = make(map[string]func())
	fmt.Println("Controller的构造函数")
}

// Prepare runs after Init before request function execution.
func (c *Controller) BeforAction() {}

// Finish runs after request function execution.

// Get adds a request function to handle GET request.
func (c *Controller) Get() {
}

// Post adds a request function to handle POST request.
func (c *Controller) Post() {
}

// Delete adds a request function to handle DELETE request.
func (c *Controller) Delete() {
}

// Put adds a request function to handle PUT request.
func (c *Controller) Put() {
}

// Head adds a request function to handle HEAD request.
func (c *Controller) Head() {
}

// Patch adds a request function to handle PATCH request.
func (c *Controller) Patch() {
}

// Options adds a request function to handle OPTIONS request.
func (c *Controller) Options() {
}

// HandlerFunc call function with the name
func (c *Controller) HandlerFunc(fnname string) bool {
	if v, ok := c.methodMapping[fnname]; ok {
		v()
		return true
	}
	return false
}
func (c *Controller) AfaterAction() {}

// URLMapping register the internal Controller router.
func (c *Controller) URLMapping() {}

// Mapping the method to function
func (c *Controller) Mapping(method string, fn func()) {
	c.methodMapping[method] = fn
}

// Render sends the response with rendered template bytes as text/html type.
func (c *Controller) Render() error {
	if !c.EnableRender {
		return nil
	}
	c.ParseRender()
	return nil
}
func (c *Controller) JsonOut(jsondata map[string]interface{}) {
	c.Ctx.Res.Header().Set("Content-Type", "application/json;charset=UTF-8")
	c.Ctx.Res.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(c.Ctx.Res).Encode(jsondata); err != nil {
		Debug.Error(err.Error())
	}
}
func (c *Controller) Assign(key string, value interface{}) {
	c.Data[key] = value
}
func (c *Controller) SetTpl(tpl string) {
	c.TplName = tpl + ".html"
}
func (c *Controller) ParseRender() {
	// 解析模板
	tmppath := c.ViewPath + c.TplName
	t, err := template.ParseFiles(tmppath)
	fmt.Println("渲染页面")
	if err != nil {
		Debug.Except(c.Ctx.Res, c.Ctx.Req, fmt.Sprintf("视图文件加载失败, err%v", err))
		return
	}
	// 渲染数据
	err = t.Execute(c.Ctx.Res, c.Data)
	if err != nil {
		Debug.Except(c.Ctx.Res, c.Ctx.Req, fmt.Sprintf("模板渲染失败, err%v", err))
		return
	}
}
