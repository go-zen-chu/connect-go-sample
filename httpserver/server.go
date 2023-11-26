package httpserver

import (
	"context"
	"log"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"connectrpc.com/connect"
	apisv1 "github.com/go-zen-chu/connect-go-sample/pkg/apis/v1"
	"github.com/go-zen-chu/connect-go-sample/pkg/apis/v1/apisv1connect"
)

//go:generate go run go.uber.org/mock/mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE

type HttpServer interface {
	Run() error
	Shutdown(ctx context.Context) error
}

type httpServer struct {
	srv *http.Server
	mux *http.ServeMux
}

func NewHttpServer(port string) HttpServer {
	return &httpServer{
		srv: &http.Server{
			Addr: port,
		},
		mux: http.NewServeMux(),
	}
}

func (h *httpServer) Run() error {
	path, handler := apisv1connect.NewGreetServiceHandler(&greetServer{})
	h.mux.Handle(path, handler)
	path, handler = apisv1connect.NewHealthServiceHandler(&healthServer{})
	h.mux.Handle(path, handler)
	h.mux.Handle("/test", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("test"))
		if err != nil {
			log.Printf("get error to /test : %v", err)
		}
	}))
	// by using h2c.NewHandler, we can use http2 protocol without TLS
	h.srv.Handler = h2c.NewHandler(h.mux, &http2.Server{})
	return h.srv.ListenAndServe()
}

func (h *httpServer) Shutdown(ctx context.Context) error {
	return h.srv.Shutdown(ctx)
}

type greetServer struct{}

func (g *greetServer) Greet(ctx context.Context, req *connect.Request[apisv1.GreetRequest]) (*connect.Response[apisv1.GreetResponse], error) {
	log.Println("recieved greet request")
	if req.Header() != nil {
		log.Println("Request headers: ", req.Header())
	}
	res := connect.NewResponse(&apisv1.GreetResponse{
		Greeting: "Hello, " + req.Msg.Name,
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}

type healthServer struct{}

func (h *healthServer) Health(ctx context.Context, req *connect.Request[apisv1.HealthRequest]) (*connect.Response[apisv1.HealthResponse], error) {
	log.Println("recieved health request")
	if req.Header() != nil {
		log.Println("Request headers: ", req.Header())
	}
	res := connect.NewResponse(&apisv1.HealthResponse{
		Status: "OK",
	})
	return res, nil
}
