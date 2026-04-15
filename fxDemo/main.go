package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

// EchoHandler is an http.Handler that copies its request body
// back to the response.
type EchoHandler struct {
	log *zap.Logger
}

// Route is an http.Handler that knows the mux pattern
// under which it will be registered.
type Route interface {
	http.Handler

	// Pattern reports the path at which this is registered.
	Pattern() string
}

// NewEchoHandler builds a new EchoHandler.
func NewEchoHandler(log *zap.Logger) *EchoHandler {
	return &EchoHandler{
		log: log,
	}
}

func (*EchoHandler) Pattern() string {
	return "/echo"
} // ServeHTTP handles an HTTP request to the /echo endpoint.

func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := io.Copy(w, r.Body); err != nil {
		h.log.Warn("Failed to handle request", zap.Error(err))
	}
}

// NewServeMux builds a ServeMux that will route requests
// to the given EchoHandler.
func NewServeMux(route Route) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle(route.Pattern(), route)
	return mux
}

func main() {

	fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Provide(
			NewHttpServer,
			fx.Annotate(
				NewEchoHandler,
				fx.As(new(Route)),
			),
			NewServeMux,
			zap.NewExample,
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
	// We use fx.Provide to add an HTTP server to the application.
	// provide the server setup func to the Fx application above with fx.Provide
	// add an fx.Invoke that requests the constructed server.
}

func NewHttpServer(lc fx.Lifecycle, mux *http.ServeMux, log *zap.Logger) *http.Server {
	srv := &http.Server{Addr: ":8080", Handler: mux} // create server

	// append the onstart and onstop hooks to the fx lifecycle
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {

				ln, err := net.Listen("tcp", srv.Addr)
				if err != nil {
					return err
				}
				log.Info("Starting Server at ", zap.String("addr:", srv.Addr))
				fmt.Println("Starting HTTP server at", srv.Addr)
				go srv.Serve(ln)
				return nil

			},

			OnStop: func(ctx context.Context) error {
				return srv.Shutdown(ctx)
			},
		})
	return srv
}
