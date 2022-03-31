package controllers

import (
	"fmt"

	"github.com/2055229754/fastgo/Controller"
	"github.com/2055229754/fastgo/fast"
)

type Index struct {
	Controller.Controller
}

func (this *Index) Show() {
	fmt.Println(fast.FastApp)

	this.Ctx.Res.Header().Set("content-type", "text/json")
	this.Ctx.Res.Write([]byte("124"))
}
