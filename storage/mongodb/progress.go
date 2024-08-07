package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	pb "progress-service/generated/progress"
	"time"
)

type ProgressRepo struct {
	coll *mongo.Collection
}

func NewProgressRepo(coll *mongo.Collection) *ProgressRepo {
	return &ProgressRepo{coll}
}

func (p *ProgressRepo) GetUserProgress(in *pb.LCodeUID) (*pb.GetUserProgressResponse, error) {
	var resp *pb.GetUserProgressResponse

	filter := bson.M{"user_id": in.UserId, "language_code": in.LanguageCode}

	err := p.coll.FindOne(context.TODO(), filter).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *ProgressRepo) GetDailyProgress(in *pb.UserID) (*pb.GetDailyProgressResponse, error) {
	var resp *pb.GetDailyProgressResponse

	filter := bson.M{"user_id": in.Id}

	err := p.coll.FindOne(context.TODO(), filter).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *ProgressRepo) GetWeeklyProgress(in *pb.UserID) (*pb.GetWeeklyProgressResponse, error) {
	var resp *pb.GetWeeklyProgressResponse

	var progress struct {
		DailyProgress []pb.Progress
	}

	startDate := time.Now().AddDate(0, 0, -7)
	endDate := time.Now()

	filter := bson.D{
		{"user_id", in.Id},
		{"date", bson.M{
			"$elemMatch": bson.M{
				"date": bson.D{
					{"$gte", startDate},
					{"$lte", endDate},
				},
			},
		},
		},
	}

	projection := bson.D{
		{"daily_progress", 1},
	}

	err := p.coll.FindOne(context.TODO(), filter, options.FindOne().SetProjection(projection)).
		Decode(&progress)
	if err != nil {
		return nil, err
	}

	var (
		xp     int64
		lesson int64
		vocab  int64
	)
	for _, pgrs := range progress.DailyProgress {
		xp += pgrs.Xp
		lesson += pgrs.LessonCompleted
		vocab += pgrs.VocabularyLearned
	}

	resp.Progress.Xp = xp
	resp.Progress.LessonCompleted = lesson
	resp.Progress.VocabularyLearned = vocab
	resp.WeekStart = startDate.String()
	resp.WeekEnd = endDate.String()

	return resp, nil
}

func (p *ProgressRepo) GetMonthlyProgress(in *pb.UserID) (*pb.GetMonthlyProgressResponse, error) {
	var resp *pb.GetMonthlyProgressResponse

	var progress struct {
		DailyProgress []pb.Progress
	}

	startDate := time.Now().AddDate(0, -1, 0)
	endDate := time.Now()

	filter := bson.D{
		{"user_id", in.Id},
		{"date", bson.M{
			"$elemMatch": bson.M{
				"date": bson.D{
					{"$gte", startDate},
					{"$lte", endDate},
				},
			},
		},
		},
	}

	projection := bson.D{
		{"daily_progress", 1},
	}

	err := p.coll.FindOne(context.TODO(), filter, options.FindOne().SetProjection(projection)).
		Decode(progress)
	if err != nil {
		return nil, err
	}

	var (
		xp     int64
		lesson int64
		vocab  int64
	)
	for _, pgrs := range progress.DailyProgress {
		xp += pgrs.Xp
		lesson += pgrs.LessonCompleted
		vocab += pgrs.VocabularyLearned
	}

	resp.Progress.Xp = xp
	resp.Progress.LessonCompleted = lesson
	resp.Progress.VocabularyLearned = vocab
	resp.Month = endDate.Month().String()

	return resp, nil
}

func (p *ProgressRepo) GetUserAchievement(in *pb.UserID) (*pb.GetUserAchievementsResponse, error) {
	resp := &pb.GetUserAchievementsResponse{}

	filter := bson.M{"user_id": in.Id}
	projection := bson.M{"achievements": 1}

	err := p.coll.FindOne(context.TODO(), filter, options.FindOne().SetProjection(projection)).
		Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *ProgressRepo) GetUserSkills(in *pb.LCodeUID) (*pb.GetUserSkillsResponse, error) {
	var resp *pb.GetUserSkillsResponse
	return resp, nil
}

func (p *ProgressRepo) GetStatistics(in *pb.LCodeUID) (*pb.GetStatisticsResponse, error) {
	var resp *pb.GetStatisticsResponse
	filter := bson.M{"user_id": in.UserId, "language_code": in.LanguageCode}

	var progress struct {
		DailyProgress []pb.Progress
	}

	err := p.coll.FindOne(context.TODO(), filter).Decode(&progress)
	if err != nil {
		return nil, err
	}

	// Подсчитываем общее количество выполненных упражнений
	var totalExercisesCompleted int64
	for _, pr := range progress.DailyProgress {
		totalExercisesCompleted += pr.LessonCompleted
	}

	// Создаем ответ
	resp.LanguageCode = in.LanguageCode
	resp.TotalExercisesCompleted = totalExercisesCompleted
	resp.TotalLessonsCompleted = int64(len(progress.DailyProgress))

	return resp, nil
}
