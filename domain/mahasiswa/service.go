package mahasiswa

import (
	"context"
	uuid "github.com/gofrs/uuid"
	"time"
)

type Service interface {
	Create(ctx context.Context, req *createRequest) (*Mahasiswa, error)
	GetByID(ctx context.Context, id string) (*Mahasiswa, error)
}

type serviceImpl struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &serviceImpl{repository}
}

func (s *serviceImpl) Create(ctx context.Context, req *createRequest) (*Mahasiswa, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	date, err := time.Parse("2006-01-02", req.DateOfBirth)
	if err != nil {
		return nil, err
	}

	mhs := &Mahasiswa{
		ID:          id.String(),
		Name:        req.Name,
		NIK:         req.NIK,
		Religion:    req.Religion,
		DateOfBirth: date,
		Gender:      req.Gender,
		Address:     req.Address,
	}

	if err = s.repository.Create(ctx, mhs); err != nil {
		return nil, err
	}

	return mhs, nil
}

func (s *serviceImpl) GetByID(ctx context.Context, id string) (*Mahasiswa, error) {
	return s.repository.GetByID(ctx, id)
}
