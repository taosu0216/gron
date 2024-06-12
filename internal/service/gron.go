package service

import (
	"context"
	"gron/internal/biz"
	"gron/pkg/pb"
)

type GronService struct {
	pb.UnimplementedGronServer
	timerUC *biz.GronUseCase
}

func (s *GronService) EnableTimer(ctx context.Context, req *pb.EnableTimerRequest) (*pb.EnableTimerResp, error) {
	return nil, nil
}

func (s *GronService) CreateTimer(ctx context.Context, req *pb.CreateTimerRequest) (*pb.CreateTimerResp, error) {
	return nil, nil
}
