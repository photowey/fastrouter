package apiconstant

import (
	"github.com/valyala/fasthttp"
)

const (
	ContentTypeApplicationJSON string = "application/json; charset=utf-8"
	ContentTypeApplicationForm string = "application/x-www-form-urlencoded; charset=utf-8"
	AcceptAll                  string = "*/*"
)

const (
	HeaderAccept        string = "Accept"
	HeaderContentType   string = "Content-Type"
	HeaderContentLength string = "Content-Length"
	HeaderUserAgent     string = "User-Agent"
	HeaderConnection    string = "Connection"

	HeaderContentTypeValue string = "application/json"
	HeaderConnectionValue  string = "Keep-Alive"
)

const (
	HeaderAuthorization        string = "Authorization"
	AuthorizationBasicTemplate string = "%s%s"
	AuthorizationBasicPrefix   string = "Basic "
)

var ValidateMethods = []string{
	fasthttp.MethodPost,
	fasthttp.MethodPut,
	fasthttp.MethodPatch,
	fasthttp.MethodDelete,
}
