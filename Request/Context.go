package Request

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Request struct {
	Req    *http.Request
	Res    http.ResponseWriter
	Params httprouter.Params
}

var Ctx Request
