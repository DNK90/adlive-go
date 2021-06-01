package services

import (
	"context"
	"github.com/dnk90/adlive/config"
	"github.com/dnk90/adlive/helpers/errors"
	"github.com/dnk90/adlive/helpers/log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dnk90/adlive/internal/stores"
	"github.com/dnk90/adlive/pb"
)

var ll = log.New()

const Version = "1.0.0"

type ReadinessCheck func() bool

func DefaultReadinessCheck() bool {
	return true
}

type service struct {
	isReady   bool
	cfg       *config.Config
	mainStore *stores.MainStore
	readinessCheck ReadinessCheck
}

func New(mainStore *stores.MainStore) pb.AdliveServer {
	return &service{
		isReady:   true,
		cfg:       config.Cfg,
		mainStore: mainStore,
	}
}

func (s *service) Version(context context.Context, req *pb.VersionRequest) (*pb.VersionResponse, error) {
	return &pb.VersionResponse{
		Version: Version,
	}, nil
}

func (s *service) Liveness(context context.Context, req *pb.LivenessRequest) (*pb.LivenessResponse, error) {
	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-osSignal:
		return nil, errors.Error(errors.Unavailable, "Server is shutting down")
	default:
		return &pb.LivenessResponse{Message: "OK"}, nil
	}
}
func (s *service) ToggleReadiness(context context.Context, req *pb.ToggleReadinessRequest) (*pb.ToggleReadinessResponse, error) {
	s.isReady = !s.isReady
	return &pb.ToggleReadinessResponse{Message: "OK"}, nil
}
func (s *service) Readiness(context context.Context, req *pb.ReadinessRequest) (*pb.ReadinessResponse, error) {
	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-osSignal:
		return nil, errors.Error(errors.Unavailable, "Server is shutting down")
	default:

		// err := s.mainStore.RelationalDatabaseCheck()

		if s.readinessCheck() == false {
			return nil, errors.Error(errors.Unavailable, "Server is not available, status: mainStore error")
		}

		if s.isReady {
			return &pb.ReadinessResponse{Message: "OK"}, nil
		}

		return nil, errors.Error(errors.Unavailable, "Server isn't ready, status: toggle off")
	}
}

