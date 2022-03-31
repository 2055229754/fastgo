package controllers

import (
	"fmt"

	"github.com/2055229754/fastgo/Controller"
)

type Index struct {
	Controller.Controller
}

func (this *Index) Show() {
	fmt.Println("index的Show方法")

	this.Ctx.Res.Header().Set("content-type", "text/json")
	this.Ctx.Res.Write([]byte("124"))
}
