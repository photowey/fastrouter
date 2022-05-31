package fastrouter

import (
	"github.com/photowey/fastrouter/api/request"
	"github.com/photowey/fastrouter/api/router"
)

var _router router.IRouter

func init() {
	_router = router.NewRouter()
}

func Register(method, path string, handler request.Handler) error {
	return _router.Register(method, path, handler)
}

func Dispatch(hctx request.Context) error {
	return _router.Dispatch(hctx)
}
