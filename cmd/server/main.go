package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dnk90/adlive/cmd/server/interceptors"

	"github.com/dnk90/adlive/config"
	l "github.com/dnk90/adlive/helpers/log"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"github.com/dnk90/adlive/pb"
	"google.golang.org/grpc"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	"go.opencensus.io/trace"
)

var (
	ll = l.New()
)

func main() {
	// load configs
	cfg := config.Cfg

	// setup google cloud tracing
	if cfg.Tracing.Enabled {
		exporter, err := stackdriver.NewExporter(stackdriver.Options{})
		if err != nil {
			ll.Fatal("failed to new exporter", l.Error(err))
		}
		trace.RegisterExporter(exporter)
	}

	// grpc server
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_prometheus.UnaryServerInterceptor,
			interceptors.ForwardMetadataUnaryServerInterceptor(),
			interceptors.LogUnaryServerInterceptor(),
			interceptors.AuthInterceptor(),
		)),
	)

	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.Register(s)
	// Register Prometheus metrics handler.

	svc := registerService()
	pb.RegisterAdliveServer(s, svc)

	// handle signal
	_, ctxCancel := context.WithCancel(context.Background())
	go func() {
		osSignal := make(chan os.Signal, 1)
		signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM)
		<-osSignal
		ctxCancel()
		// Wait for maximum 15s
		go func() {
			var durationSec time.Duration = 15
			if cfg.Environment == "D" {
				durationSec = 1
			}
			timer := time.NewTimer(durationSec * time.Second)
			<-timer.C
			ll.Fatal("Force shutdown due to timeout!")
		}()
	}()

	go func() {
		gw := NewServer()
		ll.Info("HTTP server start listening", l.Int("HTTP address", cfg.HTTPAddress))
		err := gw.RunGRPCGateway()
		if err != nil {
			ll.Fatal("error listening to address", l.Int("address", cfg.HTTPAddress), l.Error(err))
			return
		}
	}()

	ll.Info("GRPC server start listening", l.Int("GRPC address", cfg.GRPCAddress))
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPCAddress))
	if err != nil {
		ll.Fatal("error listening to address", l.Int("address", cfg.GRPCAddress), l.Error(err))
		return
	}

	ll.Fatal("error while serving", l.Error(s.Serve(listener)))
}
