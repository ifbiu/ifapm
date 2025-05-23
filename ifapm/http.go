package ifapm

import (
	"context"
	"net/http"
)

type HttpServer struct {
	mux *http.ServeMux
	*http.Server
}

func NewHttpServer(addr string) *HttpServer {
	mux := http.NewServeMux()
	server := &http.Server{Addr: addr, Handler: mux}
	s := &HttpServer{mux, server}
	globalStarters = append(globalStarters, s)
	globalClosers = append(globalClosers, s)
	return s
}

func (h *HttpServer) Handler(pattern string, handler http.Handler) {
	h.mux.Handle(pattern, handler)
}

func (h *HttpServer) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	h.mux.HandleFunc(pattern, handler)
}

func (h *HttpServer) Start() {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
}

func (h *HttpServer) Close() {
	err := h.Server.Shutdown(context.Background())
	if err != nil {
		panic(err)
	}
}
