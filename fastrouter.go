package fastrouter

//
// Fastrouter
//

import (
	"github.com/photowey/fastrouter/api/filter"
	"github.com/photowey/fastrouter/api/interceptor"
	"github.com/photowey/fastrouter/api/request"
	"github.com/photowey/fastrouter/api/router"
)

func Register(method, path string, handler request.Handler) {
	router.Register(method, path, handler)
}

func UnRegister(method, path string, handler request.Handler) error {
	return router.UnRegister(method, path, handler)
}

func RegisterFilter(name string, filterx filter.Filter) {
	router.RegisterFilter(name, filterx)
}

func UnRegisterFilter(name string, filterx filter.Filter) error {
	return router.UnRegisterFilter(name, filterx)
}

func RegisterInterceptor(name string, interceptorx interceptor.Interceptor) {
	router.RegisterInterceptor(name, interceptorx)
}

func UnRegisterInterceptor(name string, interceptorx interceptor.Interceptor) error {
	return router.UnRegisterInterceptor(name, interceptorx)
}
