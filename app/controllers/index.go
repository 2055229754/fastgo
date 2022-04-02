package controllers

import (
	"errors"

	"github.com/2055229754/fastgo/Controller"
	"github.com/2055229754/fastgo/app/models"
)

var ErrAbort = errors.New("user stop run")

type Index struct {
	Controller.Controller
}

func (this *Index) Show() {
	admin := new(models.Admin)
	admin.Show()

	this.Ctx.Res.Header().Set("content-type", "text/json")
	this.Ctx.Res.Write([]byte("124"))
}
