package router

import (
	"sort"
	"sync"

	"github.com/photowey/fastrouter/api/request"
	perrors "github.com/pkg/errors"
)

var (
	_       IRouter = (*Router)(nil)
	_router *Router
	_lock   sync.Mutex
)

func init() {
	_router = newRouter()
}

type IRouter interface {
	Register(method, path string, handler request.Handler) error
	Handlers() []request.Handler
	Dispatch(hctx request.Context) error
}

type Router struct {
	HandlerMapping  map[string]request.Handler // handler mapping
	RequestHandlers []request.Handler          // request handlers
}

func (r *Router) Register(method, path string, handler request.Handler) error {
	// TODO validate?
	requestMapping := BuildMapping(method, path)
	r.HandlerMapping[requestMapping] = handler

	return nil
}

func (r *Router) Handlers() []request.Handler {
	if len(_router.RequestHandlers) == 0 {
		_lock.Lock()
		defer _lock.Unlock()
		if len(_router.RequestHandlers) == 0 {
			r.initRhs()
		}
	}

	return r.RequestHandlers
}

func (r *Router) initRhs() {
	hs := make([]request.Handler, 0)
	for _, handler := range r.HandlerMapping {
		hs = append(hs, handler)
	}
	sort.SliceStable(hs, func(i, j int) bool {
		return hs[i].Order() < hs[j].Order()
	})

	r.RequestHandlers = hs
}

func (r *Router) Dispatch(hctx request.Context) error {
	handlers := r.Handlers()
	for _, handler := range handlers {
		if handler.Supports(hctx) {
			handler.Handle(hctx)
			return nil
		}
	}

	return perrors.Errorf("router: Dispatch, handler mapping not found: %s", hctx.BuildMapping())
}

func NewRouter() IRouter {
	if _router == nil {
		_lock.Lock()
		defer _lock.Unlock()
		if _router == nil {
			_router = newRouter()
		}
	}

	return _router
}

func newRouter() *Router {
	return &Router{
		HandlerMapping:  make(map[string]request.Handler, 0),
		RequestHandlers: make([]request.Handler, 0),
	}
}
