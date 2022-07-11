package main

import (
	"context"
	"sync"

	"github.com/bal3000/PokeCentre/proto/trainers"
	"github.com/bal3000/PokeCentre/services/trainers/data"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TrainersService struct {
	mu    sync.Mutex
	model data.TrainersModeler
	trainers.UnimplementedTrainersServiceServer
}

func NewTrainerService(model data.TrainersModeler) *TrainersService {
	return &TrainersService{
		model: model,
	}
}

func (t *TrainersService) AddTrainer(ctx context.Context, request *trainers.AddTrainerRequest) (*trainers.AddTrainerResponse, error) {
	m := &data.Trainer{
		Name:      request.Name,
		Email:     request.Email,
		Address:   request.Address,
		Phone:     request.Phone,
		NhsNumber: request.NhsNumber,
	}

	err := t.model.Insert(ctx, m)
	if err != nil {
		return nil, err
	}

	return &trainers.AddTrainerResponse{
		Id:        m.ID,
		Name:      m.Name,
		Phone:     m.Phone,
		Email:     m.Email,
		Address:   m.Address,
		NhsNumber: m.NhsNumber,
		CreatedAt: timestamppb.New(m.CreatedAt),
		UpdatedAt: timestamppb.New(m.UpdatedAt),
	}, nil
}

func (t *TrainersService) UpdateTrainer(ctx context.Context, request *trainers.UpdateTrainerRequest) (*trainers.UpdateTrainerResponse, error) {
	m := &data.Trainer{
		ID:        request.Id,
		Name:      request.Name,
		Email:     request.Email,
		Address:   request.Address,
		Phone:     request.Phone,
		NhsNumber: request.NhsNumber,
	}

	err := t.model.Update(ctx, request.Id, m)
	if err != nil {
		return &trainers.UpdateTrainerResponse{
			Success: false,
		}, err
	}

	return &trainers.UpdateTrainerResponse{
		Success: true,
	}, nil
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
