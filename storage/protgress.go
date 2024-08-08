package storage

import (
	pb "progress-service/generated/progress"
)

type Mongodb interface {
	GetUserProgress(in *pb.LCodeUID) (*pb.GetUserProgressResponse, error)
	GetDailyProgress(in *pb.UserID) (*pb.GetDailyProgressResponse, error)
	GetWeeklyProgress(in *pb.UserID) (*pb.GetWeeklyProgressResponse, error)
	GetMonthlyProgress(in *pb.UserID) (*pb.GetMonthlyProgressResponse, error)
	GetUserAchievement(in *pb.UserID) (*pb.GetUserAchievementsResponse, error)
	//GetLeaders(in *pb.LanguageCode) (*pb.GetLeadersResponse, error)
	GetUserSkills(in *pb.LCodeUID) (*pb.GetUserSkillsResponse, error)
	GetStatistics(in *pb.LCodeUID) (*pb.GetStatisticsResponse, error)
}
