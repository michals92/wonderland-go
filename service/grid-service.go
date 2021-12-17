package service

import (
	"context"
	"errors"

	//"github.com/go-redis/redis/v7"
	"github.com/michals92/wonderland-go/entity"
	"github.com/michals92/wonderland-go/repository"
)

type gridService struct{}

//var client *redis.Client
var ctx = context.Background()

type GridService interface {
	GetGrid(box *entity.BoundingBox) ([]entity.Parcel, error)
	AddParcel(parcel *entity.Parcel) error
}

var (
	repo repository.Repository
)

func NewGridService(repository repository.Repository) GridService {
	repo = repository
	//	initRedis()
	return &gridService{}
}

func (*gridService) GetGrid(box *entity.BoundingBox) ([]entity.Parcel, error) {
	//TODO: implement user
	return nil, errors.New("get grid not implemented")
}

func (*gridService) AddParcel(parcel *entity.Parcel) error {
	return repo.AddParcel(parcel)
}
