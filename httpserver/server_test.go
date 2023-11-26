package httpserver

import (
	"context"
	"errors"
	"log"
	"net/http"
	"reflect"
	"testing"

	"connectrpc.com/connect"
	apisv1 "github.com/go-zen-chu/connect-go-sample/pkg/apis/v1"
)

func Test_httpServer_Run(t *testing.T) {
	type fields struct {
		srv *http.Server
		mux *http.ServeMux
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "if new server and mux given, it should run server",
			fields: fields{
				srv: &http.Server{
					Addr: ":28080", // need to use unused port
				},
				mux: http.NewServeMux(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &httpServer{
				srv: tt.fields.srv,
				mux: tt.fields.mux,
			}
			go func() {
				if err := h.Run(); (err != nil) != tt.wantErr {
					if !errors.Is(http.ErrServerClosed, err) {
						t.Errorf("httpServer.Run() error = %v, wantErr %v", err, tt.wantErr)
					}
					log.Println("closing test server...")
				}
			}()
			if err := h.Shutdown(context.Background()); err != nil {
				t.Errorf("httpServer.Shutdown() error = %v", err)
			}
		})
	}
}

func Test_httpServer_Shutdown(t *testing.T) {
	type fields struct {
		srv *http.Server
		mux *http.ServeMux
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &httpServer{
				srv: tt.fields.srv,
				mux: tt.fields.mux,
			}
			if err := h.Shutdown(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("httpServer.Shutdown() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_greetServer_Greet(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[apisv1.GreetRequest]
	}
	tests := []struct {
		name    string
		g       *greetServer
		args    args
		want    *connect.Response[apisv1.GreetResponse]
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &greetServer{}
			got, err := g.Greet(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("greetServer.Greet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("greetServer.Greet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_healthServer_Health(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[apisv1.HealthRequest]
	}
	tests := []struct {
		name    string
		h       *healthServer
		args    args
		want    *connect.Response[apisv1.HealthResponse]
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &healthServer{}
			got, err := h.Health(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("healthServer.Health() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("healthServer.Health() = %v, want %v", got, tt.want)
			}
		})
	}
}
