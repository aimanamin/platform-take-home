package server

import (
	"context"
	"errors"

	"github.com/rs/cors"
	"github.com/skip-mev/platform-take-home/observability/logging"
	"go.uber.org/zap"

	"fmt"
	"net"
	"net/http"

	"github.com/skip-mev/platform-take-home/api/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

// StartGRPCGateway starts the gRPC gateway server.
func StartGRPCGateway(ctx context.Context, host string, port int) error {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames:   true,
				EmitUnpopulated: true,
			},
		}))

	if err := types.RegisterTakeHomeServiceHandlerFromEndpoint(ctx, mux, "localhost:9008", opts); err != nil {
		return err
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return fmt.Errorf("error creating listener: %v", err)
	}

	corsMiddleware := cors.New(cors.Options{})
	handler := corsMiddleware.Handler(mux)

	// Create a new HTTP server
	server := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Check if the request is for the /health endpoint
			if r.URL.Path == "/health" {
				healthCheckHandler(w, r)
				return
			}
			// Otherwise, delegate to the gRPC gateway mux
			handler.ServeHTTP(w, r)
		}),
	}

	go func() {
		<-ctx.Done()
		if err := server.Shutdown(context.Background()); err != nil {
			logging.FromContext(ctx).Fatal("error shutting down http server", zap.Error(err))
		}
	}()

	if err := server.Serve(listener); !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("error serving http: %v", err)
	}

	return nil
}

// healthCheckHandler handles the /health endpoint.
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
