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

	var infoPar []entity.Parcel

	for _, s := range *parcels {
		s.Type = "bought"
		infoPar = append(infoPar, s)
	}

	return &infoPar, nil
}

func (*gridService) AddParcel(parcel *entity.Parcel) error {
	parcelExists := repo.ParcelExists(parcel.H3Index)

	if parcelExists {
		return errors.New("parcel already exists")
	}

	return repo.AddParcel(parcel)
}
