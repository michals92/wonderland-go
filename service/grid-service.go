package service

import (
	"context"

	//"github.com/go-redis/redis/v7"
	"github.com/michals92/wonderland-go/entity"
	"github.com/michals92/wonderland-go/repository"
)

type gridService struct{}

//var client *redis.Client
var ctx = context.Background()

type GridService interface {
	GetGrid(box *entity.BoundingBox) (*[]entity.Parcel, error)
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

func (*gridService) GetGrid(box *entity.BoundingBox) (*[]entity.Parcel, error) {
	parcels, error := repo.GetParcels(box)
	if error != nil {
		return nil, error
	}

	return parcels, nil
}

func (*gridService) AddParcel(parcel *entity.Parcel) error {
	return repo.AddParcel(parcel)
}
