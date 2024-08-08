package service

import (
	"context"
	"log/slog"
	pb "progress-service/generated/progress"
	"progress-service/storage"
)

type Service struct {
	pb.UnimplementedProgressServiceServer
	Logger *slog.Logger
	storage.Storage
}

func NewService(logger *slog.Logger, st storage.Storage) *Service {
	return &Service{Logger: logger, Storage: st}
}

func (s *Service) GetUserProgress(ctx context.Context, in *pb.LCodeUID) (*pb.GetUserProgressResponse, error) {
	res, err := s.Storage.NewMongoStorage().GetUserProgress(in)
	if err != nil {
		s.Logger.Error("error in Get UserProgress", "error", err)
		return nil, err
	}
	return res, nil
}

func (s *Service) GetDailyProgress(ctx context.Context, in *pb.UserID) (*pb.GetDailyProgressResponse, error) {
	res, err := s.Storage.NewMongoStorage().GetDailyProgress(in)
	if err != nil {
		s.Logger.Error("error in Get Daily Progress", "error", err)
		return nil, err
	}
	return res, nil
}

func (s *Service) GetWeeklyProgress(ctx context.Context, in *pb.UserID) (*pb.GetWeeklyProgressResponse, error) {
	res, err := s.Storage.NewMongoStorage().GetWeeklyProgress(in)
	if err != nil {
		s.Logger.Error("error in Get Weekly Progress", "error", err)
		return nil, err
	}
	return res, nil
}

func (s *Service) GetMonthlyProgress(ctx context.Context, in *pb.UserID) (*pb.GetMonthlyProgressResponse, error) {
	res, err := s.Storage.NewMongoStorage().GetMonthlyProgress(in)
	if err != nil {
		s.Logger.Error("error in Get Monthly Progress", "error", err)
		return nil, err
	}
	return res, nil
}

func (s *Service) GetUserAchievement(ctx context.Context, in *pb.UserID) (*pb.GetUserAchievementsResponse, error) {
	res, err := s.Storage.NewMongoStorage().GetUserAchievement(in)
	if err != nil {
		s.Logger.Error("error in Get User Achievement", "error", err)
		return nil, err
	}
	return res, nil
}

func (s *Service) GetUserSkills(ctx context.Context, in *pb.LCodeUID) (*pb.GetUserSkillsResponse, error) {
	res, err := s.Storage.NewMongoStorage().GetUserSkills(in)
	if err != nil {
		s.Logger.Error("error in Get User Skills", "error", err)
		return nil, err
	}
	return res, nil
}

func (s *Service) GetStatistics(ctx context.Context, in *pb.LCodeUID) (*pb.GetStatisticsResponse, error) {
	res, err := s.Storage.NewMongoStorage().GetStatistics(in)
	if err != nil {
		s.Logger.Error("error in Get Statistics", "error", err)
		return nil, err
	}
	return res, nil
}
