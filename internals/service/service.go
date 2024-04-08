package service

import (
	"tools/internals/repository"
)

type Service struct {
	rp *repository.Repository
}

func NewService(rp *repository.Repository) *Service {
	return &Service{rp: rp}
}
