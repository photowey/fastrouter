package dispatcher

import (
	"sort"
	"sync"

	"github.com/photowey/fastrouter/api/apiconstant"
	"github.com/photowey/fastrouter/api/filter"
	"github.com/photowey/fastrouter/api/interceptor"
	"github.com/photowey/fastrouter/api/request"
	"github.com/photowey/fastrouter/internal/pkg/pathx"
	"github.com/photowey/fastrouter/pkg/jsonx"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
)

var (
	_     Handler = (*handler)(nil)
	_lock sync.Mutex
)

type Handler interface {
	Register(method, path string, handler request.Handler)
	UnRegister(method, path string, handler request.Handler) error
	RegisterFilter(name string, filterx filter.Filter)
	UnRegisterFilter(name string, filterx filter.Filter) error
	RegisterInterceptor(name string, interceptorx interceptor.Interceptor)
	UnRegisterInterceptor(name string, interceptorx interceptor.Interceptor) error
	handlers() []request.Handler
	Route(rctx *request.Context, ctx *fasthttp.RequestCtx) error
	dispatch(rctx *request.Context) error
}

type handler struct {
	handlerMapping  map[string]request.Handler         // handler mapping
	requestHandlers []request.Handler                  // request handlers
	filters         map[string]filter.Filter           // filter
	interceptors    map[string]interceptor.Interceptor // interceptors
}

func (r *handler) Register(method, path string, handler request.Handler) {
	_lock.Lock()
	defer _lock.Unlock()
	if method == "" {
		panic("handler.registry: web.http.handler.registry: Register method is blank")
	}
	if handler == nil {
		panic("handler.registry: web.http.handler.registry: Register handler is nil")
	}

	requestMapping := pathx.BuildMapping(method, path)
	if _, dup := r.handlerMapping[requestMapping]; dup {
		panic("handler.registry: web.http.handler.registry: Register called twice for path: " + requestMapping)
	}
	r.handlerMapping[requestMapping] = handler
}

func (r *handler) UnRegister(method, path string, handler request.Handler) error {
	_lock.Lock()
	defer _lock.Unlock()
	if method == "" {
		return errors.New("handler.registry: web.http.handler.registry: UnRegister method is blank")
	}
	if handler == nil {
		return errors.New("handler.registry: web.http.handler.registry: UnRegister handler is nil")
	}

	requestMapping := pathx.BuildMapping(method, path)
	if _, ok := r.handlerMapping[requestMapping]; ok {
		delete(r.handlerMapping, requestMapping)
	}

	r.initRhs()

	return nil
}

func (r *handler) RegisterFilter(name string, filterx filter.Filter) {
	if name == "" {
		panic("filter.registry: web.http.filter.registry: Register filter name is blank")
	}
	if filterx == nil {
		panic("filter.registry: web.http.filter.registry: Register filter is nil")
	}

	if _, dup := r.filters[name]; dup {
		panic("handler.registry: web.http.filter.registry: RegisterFilter called twice for filter: " + name)
	}

	r.filters[name] = filterx
}

func (r *handler) UnRegisterFilter(name string, filterx filter.Filter) error {
	if name == "" {
		return errors.New("filter.registry: web.http.filter.registry: UnRegister filter name is blank")
	}
	if filterx == nil {
		return errors.New("filter.registry: web.http.filter.registry: UnRegister filter is nil")
	}

	if _, ok := r.filters[name]; ok {
		delete(r.filters, name)
	}

	return nil
}

func (r *handler) RegisterInterceptor(name string, interceptorx interceptor.Interceptor) {
	if name == "" {
		panic("interceptor.registry: web.http.interceptor.registry: Register interceptor name is blank")
	}
	if interceptorx == nil {
		panic("interceptor.registry: web.http.interceptor.registry: Register interceptor is nil")
	}

	if _, dup := r.filters[name]; dup {
		panic("handler.registry: web.http.interceptor.registry: RegisterInterceptor called twice for filter: " + name)
	}

	r.interceptors[name] = interceptorx
}

func (r *handler) UnRegisterInterceptor(name string, interceptorx interceptor.Interceptor) error {
	if name == "" {
		return errors.New("filter.registry: web.http.interceptor.registry: UnRegister interceptor name is blank")
	}
	if interceptorx == nil {
		return errors.New("filter.registry: web.http.interceptor.registry: UnRegister interceptor is nil")
	}

	if _, ok := r.interceptors[name]; ok {
		delete(r.interceptors, name)
	}

	return nil
}

// ----------------------------------------------------------------

/*
	0.parse
	1.do filter
	2.pre handle
	3. dispatch
	4. post handle
	5.header
	6.response status code
	7.response
*/
func (r *handler) Route(rctx *request.Context, ctx *fasthttp.RequestCtx) error {
	// 0.parse
	method := rctx.Method
	if ArrayContains(apiconstant.ValidateMethods, method) {
		body := ctx.PostBody()
		if body != nil {
			bodyMap := jsonx.ToStringMap(string(body))
			rctx.BodyMap = bodyMap
		}
		// TODO form body && headers
	}

	// 1.do filter
	filters := r.sortedFilters()
	if len(filters) > 0 {
		for _, filterx := range filters {
			err := filterx.DoFilter(rctx)
			if err != nil {
				flush(rctx, ctx)
				return err
			}
		}
	}

	// 2.pre handle
	interceptors := r.sortedInterceptors()
	length := len(interceptors)
	if length > 0 {
		for _, interceptorx := range interceptors {
			ok := interceptorx.PreHandle(rctx)
			if !ok {
				flush(rctx, ctx)
				return nil
			}
		}
	}

	// 3. handle
	err := r.dispatch(rctx)

	// 4. post handle
	if length > 0 {
		for i := length - 1; i >= 0; i-- {
			interceptorx := interceptors[i]
			interceptorx.PostHandle(rctx)
		}
	}

	// 5.header
	if rctx.Header.Length() > 0 {
		values := rctx.Header.Values()
		for k, vs := range values {
			for _, v := range vs {
				ctx.Response.Header.Add(k, v)
			}
		}
	}

	// 6.response status code
	// 7.response
	flush(rctx, ctx)

	return err
}

// ----------------------------------------------------------------

func (r *handler) dispatch(rctx *request.Context) error {
	handlers := r.handlers()
	for _, handlerx := range handlers {
		if handlerx.Supports(rctx) {
			handlerx.Handle(rctx)
			return nil
		}
	}

	return errors.Errorf("dispatcher.handler: Dispatch, handler mapping not found: %s", rctx.BuildMapping())
}

// ----------------------------------------------------------------

func (r *handler) handlers() []request.Handler {
	if len(r.requestHandlers) == 0 {
		_lock.Lock()
		defer _lock.Unlock()
		if len(r.requestHandlers) == 0 {
			r.initRhs()
		}
	}

	return r.requestHandlers
}

// ----------------------------------------------------------------

func (r *handler) initRhs() {
	hs := make([]request.Handler, 0)
	for _, handlerx := range r.handlerMapping {
		hs = append(hs, handlerx)
	}
	sort.SliceStable(hs, func(i, j int) bool {
		return hs[i].Order() < hs[j].Order()
	})

	r.requestHandlers = hs
}

// ----------------------------------------------------------------

func (r *handler) sortedFilters() []filter.Filter {
	filters := make([]filter.Filter, 0, len(r.filters))

	for _, v := range r.filters {
		filters = append(filters, v)
	}

	sort.SliceStable(filters, func(i, j int) bool {
		return filters[i].Order() < filters[j].Order()
	})

	return filters
}

func (r *handler) sortedInterceptors() []interceptor.Interceptor {
	interceptors := make([]interceptor.Interceptor, 0, len(r.interceptors))

	for _, v := range r.interceptors {
		interceptors = append(interceptors, v)
	}

	sort.SliceStable(interceptors, func(i, j int) bool {
		// 升序 -> order 越小,优先级越高
		return interceptors[i].Order() < interceptors[j].Order()
	})

	return interceptors
}

// ----------------------------------------------------------------

func NewHandler() Handler {
	return &handler{
		handlerMapping:  make(map[string]request.Handler, 0),
		requestHandlers: make([]request.Handler, 0),
		filters:         make(map[string]filter.Filter, 0),
		interceptors:    make(map[string]interceptor.Interceptor, 0),
	}
}

// ----------------------------------------------------------------

func flush(rctx *request.Context, ctx *fasthttp.RequestCtx) {
	status(rctx, ctx)
	response(rctx, ctx)
}

func status(rctx *request.Context, ctx *fasthttp.RequestCtx) {
	if rctx.StatusCode != fasthttp.StatusOK {
		ctx.SetStatusCode(rctx.StatusCode)
	}
}

func response(rctx *request.Context, ctx *fasthttp.RequestCtx) {
	if rctx.HandleResponse {
		if rctx.Body != nil {
			ctx.SetBody(rctx.Body)
		}
	}
}
