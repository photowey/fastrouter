package router

import (
	"github.com/photowey/fastrouter/api/apiconstant"
	"github.com/photowey/fastrouter/api/dispatcher"
	"github.com/photowey/fastrouter/api/filter"
	"github.com/photowey/fastrouter/api/interceptor"
	"github.com/photowey/fastrouter/api/request"
	"github.com/valyala/fasthttp"
)

var (
	_       Router = (*router)(nil)
	_router *router
)

func init() {
	_router = newRouter()
}

type Router interface {
	route(rctx *request.Context, ctx *fasthttp.RequestCtx) error
}
type router struct {
	handler dispatcher.Handler
}

func (r *router) route(rctx *request.Context, ctx *fasthttp.RequestCtx) error {
	return _router.handler.Route(rctx, ctx)
}

func Register(method, path string, handler request.Handler) {
	_router.handler.Register(method, path, handler)
}

func UnRegister(method, path string, handler request.Handler) error {
	return _router.handler.UnRegister(method, path, handler)
}

func RegisterFilter(name string, filterx filter.Filter) {
	_router.handler.RegisterFilter(name, filterx)
}

func UnRegisterFilter(name string, filterx filter.Filter) error {
	return _router.handler.UnRegisterFilter(name, filterx)
}

func RegisterInterceptor(name string, interceptorx interceptor.Interceptor) {
	_router.handler.RegisterInterceptor(name, interceptorx)
}

func UnRegisterInterceptor(name string, interceptorx interceptor.Interceptor) error {
	return _router.handler.UnRegisterInterceptor(name, interceptorx)
}

func Route(ctx *fasthttp.RequestCtx) error {
	ctx.SetContentType(apiconstant.ContentTypeApplicationJSON)
	ctx.SetStatusCode(fasthttp.StatusOK)
	rctx := request.NewRequestContext(ctx)

	return _router.route(rctx, ctx)
}

func newRouter() *router {
	return &router{
		handler: dispatcher.NewHandler(),
	}
}
