package wine

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gopub/log"
	pathutil "github.com/gopub/wine/internal/path"
	"github.com/gopub/wine/internal/template"
)

var acceptEncodings = [2]string{"gzip", "defalte"}
var ShortHandlerNameFlag = true

const (
	endpointPath = "_endpoints"
)

// Server implements web server
type Server struct {
	*Router
	*TemplateManager

	header  http.Header
	timeout time.Duration //timeout for each request, default value is 20s
	server  *http.Server

	paramsParser ParamsParser

	faviconHandlerList  *handlerList
	notfoundHandlerList *handlerList
	optionsHandlerList  *handlerList

	logger log.Logger

	BeginHandler Handler
}

// NewServer returns a server
func NewServer(config *Config) *Server {
	if config == nil {
		config = DefaultConfig()
	}

	s := &Server{
		Router:          NewRouter(),
		TemplateManager: NewTemplateManager(),
		header:          config.ResponseHeaders,
		timeout:         config.RequestTimeout,
		paramsParser:    config.ParamsParser,
		logger:          config.Logger,
	}

	s.faviconHandlerList = newHandlerList([]Handler{HandlerFunc(handleFavIcon)})
	s.notfoundHandlerList = newHandlerList([]Handler{HandlerFunc(handleNotFound)})
	s.optionsHandlerList = newHandlerList([]Handler{HandlerFunc(s.handleOptions)})

	s.AddTemplateFuncMap(template.FuncMap)

	if config != nil {
		s.handlers = config.Handlers
	}
	return s
}

// Run starts server
func (s *Server) Run(addr string) error {
	if s.server != nil {
		logger.Panic("Server is running")
	}

	logger.Info("Running at", addr, "...")
	s.server = &http.Server{Addr: addr, Handler: s}
	err := s.server.ListenAndServe()
	if err != nil {
		s.server = nil
	}
	return err
}

// RunTLS starts server with tls
func (s *Server) RunTLS(addr, certFile, keyFile string) error {
	if s.server != nil {
		logger.Panic("Server is running")
	}

	s.Router.Print()
	logger.Info("Running at", addr, "...")
	s.server = &http.Server{Addr: addr, Handler: s}
	err := s.server.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		s.server = nil
	}
	return err
}

// Shutdown stops server
func (s *Server) Shutdown() error {
	return s.server.Shutdown(context.Background())
}

func (s *Server) logHTTP(rw http.ResponseWriter, req *http.Request, startAt time.Time) {
	var statGetter statusGetter
	if cw, ok := rw.(*CompressedResponseWriter); ok {
		cw.Close()
		statGetter = cw.ResponseWriter.(statusGetter)
	}

	if statGetter == nil {
		statGetter = rw.(statusGetter)
	}

	info := fmt.Sprintf("%s %s %s %s | return %d in %v",
		req.RemoteAddr,
		req.UserAgent(),
		req.Method,
		req.RequestURI,
		statGetter.Status(),
		time.Since(startAt))

	if statGetter.Status() >= http.StatusBadRequest {
		if statGetter.Status() != http.StatusUnauthorized {
			s.logger.Errorf("%s header: %v form: %v", info, req.Header, req.PostForm)
		} else {
			s.logger.Errorf("%s header: %v", info, req.Header)
		}
	} else {
		s.logger.Info(info)
	}
}

// ServeHTTP implements for http.Handler interface, which will handle each http request
func (s *Server) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	startAt := time.Now()
	if logger.Level() > log.DebugLevel {
		defer func() {
			if e := recover(); e != nil {
				logger.Error(e, req)
			}
		}()
	}

	// Add compression to responseWriter
	rw = createCompressedWriter(NewResponseWriter(rw), req)
	defer s.logHTTP(rw, req, startAt)

	path := pathutil.NormalizedRequestPath(req)
	method := strings.ToUpper(req.Method)
	handlers, pathParams := s.Match(method, path)

	if handlers.Empty() {
		if method == http.MethodOptions {
			handlers = s.optionsHandlerList
		} else if path == "favicon.ico" {
			handlers = s.faviconHandlerList
		} else {
			handlers = s.notfoundHandlerList
		}
	}

	ctx, cancel := context.WithTimeout(req.Context(), s.timeout)
	defer cancel()

	parsedReq, err := NewRequest(req, s.paramsParser)
	if err != nil {
		logger.Errorf("Parse failed: %v", err)
		resp := Text(http.StatusBadRequest, fmt.Sprintf("Parse request failed: %v", err))
		resp.Respond(ctx, rw)
		return
	}
	parsedReq.params.AddMapObj(pathParams)

	for k, v := range s.header {
		rw.Header()[k] = v
	}

	ctx = withTemplate(ctx, s.templates)
	ctx = withResponseWriter(ctx, rw)
	var resp Responsible
	if s.BeginHandler != nil && path != endpointPath {
		resp = s.BeginHandler.HandleRequest(ctx, parsedReq, handlers.Head().Invoke)
	} else {
		resp = handlers.Head().Invoke(ctx, parsedReq)
	}

	if resp == nil {
		resp = handleNotImplemented(ctx, parsedReq, nil)
	}
	resp.Respond(ctx, rw)
}

func createCompressedWriter(rw http.ResponseWriter, req *http.Request) http.ResponseWriter {
	// Add compression to responseWriter
	ae := req.Header.Get("Accept-Encoding")
	for _, enc := range acceptEncodings {
		if strings.Contains(ae, enc) {
			rw.Header().Set("Content-Encoding", enc)
			cw, err := NewCompressedResponseWriter(rw, enc)
			if err != nil {
				log.Errorf("NewCompressedResponseWriter failed: %v", err)
			}
			return cw
		}
	}

	return rw
}

func (s *Server) handleOptions(ctx context.Context, req *Request, next Invoker) Responsible {
	path := pathutil.NormalizedRequestPath(req.request)
	var allowedMethods []string
	for routeMethod := range s.Router.methodTrees {
		if handlers, _ := s.Match(routeMethod, path); !handlers.Empty() {
			allowedMethods = append(allowedMethods, routeMethod)
		}
	}

	if len(allowedMethods) > 0 {
		allowedMethods = append(allowedMethods, http.MethodOptions)
	}

	return ResponsibleFunc(func(ctx context.Context, rw http.ResponseWriter) {
		if len(allowedMethods) > 0 {
			val := []string{strings.Join(allowedMethods, ",")}
			rw.Header()["Allow"] = val
			rw.Header()["Access-Control-Allow-Methods"] = val
			rw.WriteHeader(http.StatusNoContent)
		} else {
			rw.WriteHeader(http.StatusNotFound)
		}
	})
}
