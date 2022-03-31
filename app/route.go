package app

import (
	"github.com/2055229754/fastgo/Router"
	"github.com/2055229754/fastgo/app/controllers"
)

func init() {
	router := Router.New()
	router.Add("/index/show", &controllers.Index{}, "get", "Show")
	router.Add("/index/a", &controllers.Index{}, "get", "a")
}
