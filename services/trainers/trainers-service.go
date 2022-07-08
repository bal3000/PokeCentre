package main

import (
	"context"
	"database/sql"
	"sync"

	"github.com/bal3000/PokeCentre/proto/trainers"
	"github.com/go-redis/redis/v9"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TrainersService struct {
	mu    sync.Mutex
	db    *sql.DB
	redis *redis.Client
	trainers.UnimplementedTrainersServiceServer
}

func NewTrainerService(db *sql.DB, redis *redis.Client) *TrainersService {
	return &TrainersService{
		db:    db,
		redis: redis,
	}
}

func (t *TrainersService) AddTrainer(ctx context.Context, request *trainers.AddTrainerRequest) (*trainers.AddTrainerResponse, error) {
	return nil, nil
}

func (t *TrainersService) UpdateTrainer(ctx context.Context, request *trainers.UpdateTrainerRequest) (*trainers.UpdateTrainerResponse, error) {
	return nil, nil
}

func (t *TrainersService) DeleteTrainer(ctx context.Context, request *trainers.DeleteTrainerRequest) (*trainers.DeleteTrainerResponse, error) {
	return nil, nil
}

func (t *TrainersService) GetTrainer(ctx context.Context, request *trainers.GetTrainerRequest) (*trainers.GetTrainerResponse, error) {
	return nil, nil
}

func (t *TrainersService) GetTrainerForPatient(ctx context.Context, request *trainers.GetTrainerForPatientRequest) (*trainers.GetTrainerForPatientResponse, error) {
	return nil, nil
}

func (t *TrainersService) GetAllTrainers(_ *emptypb.Empty, service trainers.TrainersService_GetAllTrainersServer) error {
	return nil
}
