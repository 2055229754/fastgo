package Debug

import (
	"fmt"
	"html/template"
	"net/http"
	"runtime"
)

func Except(w http.ResponseWriter, r *http.Request, msg string) {
	_, file, line, _ := runtime.Caller(3)
	var data = make(map[string]interface{})
	data["msg"] = msg
	data["file"] = file
	data["line"] = line
	t, _ := template.ParseFiles("./Template/fast_exception.tpl")
	// 渲染数据
	t.Execute(w, data)
	fmt.Println(msg)
	return
}

func Error(msg string) {
	fmt.Println(msg)
	return
}
