package main

import (
	"fmt"

	_ "github.com/2055229754/fastgo/app"
	"github.com/2055229754/fastgo/fast"
)

func main() {
	fa := fast.New()
	for _, v := range fa.Route.Routes {
		fmt.Println("route", v)
	}
	fa.Start()
}
