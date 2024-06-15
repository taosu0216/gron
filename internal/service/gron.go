package service

import (
	"context"
	"encoding/json"
	"gron/internal/biz"
	"gron/pkg/pb"
)

type GronService struct {
	pb.UnimplementedGronServer
	timerUC *biz.GronUseCase
}

func (s *GronService) EnableTimer(ctx context.Context, req *pb.EnableTimerRequest) (*pb.EnableTimerResp, error) {
	err := s.timerUC.EnableTimer(ctx, req.TimerId)
	return nil, nil
}

func (s *GronService) CreateTimer(ctx context.Context, req *pb.CreateTimerRequest) (*pb.CreateTimerResp, error) {
	param, err := json.Marshal(req.NotifyHTTPParam)
	if err != nil {
		return nil, err
	}
	timert, err := s.timerUC.CreateTimer(ctx, &biz.Timer{
		App:             req.App,
		Name:            req.Name,
		Status:          int(req.Status),
		Cron:            req.Cron,
		NotifyHTTPParam: string(param),
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateTimerResp{
		Code:    0,
		Message: "ok",
		Data: &pb.CreateTimerRespData{
			TimerId: timert.TimerId,
		},
	}, nil
}
