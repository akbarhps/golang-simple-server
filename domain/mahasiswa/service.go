package mahasiswa

import (
	"context"
	uuid "github.com/gofrs/uuid"
	"time"
)

type Service interface {
	Create(ctx context.Context, req *createRequest) (*Mahasiswa, error)
}

type serviceImpl struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &serviceImpl{repository}
}

func (s *serviceImpl) Create(ctx context.Context, req *createRequest) (*Mahasiswa, error) {
	id, _ := uuid.NewV4()
	date, _ := time.Parse("2006-01-02", req.DateOfBirth)

	mhs := &Mahasiswa{
		ID:          id.String(),
		Name:        req.Name,
		NIK:         req.NIK,
		Religion:    req.Religion,
		DateOfBirth: date,
		Gender:      req.Gender,
		Address:     req.Address,
	}

	s.repository.Create(ctx, mhs)

	return mhs, nil
}
