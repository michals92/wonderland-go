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
	GetGrid(userInfo *entity.UserInfo) (*[]entity.Parcel, error)
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

func (*gridService) GetGrid(userInfo *entity.UserInfo) (*[]entity.Parcel, error) {
	parcels, error := repo.GetParcels(&userInfo.BoundingBox)
	if error != nil {
		return nil, error
	}

	var infoPar []entity.Parcel

	for _, s := range *parcels {

		if s.Owner == userInfo.Wallet {
			s.Type = "mine"
		} else {
			s.Type = "bought"
		}

		infoPar = append(infoPar, s)
	}

	return &infoPar, nil
}

func (*gridService) AddParcel(parcel *entity.Parcel) error {
	parcelExists := repo.ParcelExists(parcel.H3Index)

	if parcelExists {
		return errors.New("parcel already exists")
	}

	if parcel.Name == "" {
		return errors.New("empty parcel name")
	}

	if parcel.Owner == "" {
		return errors.New("empty parcel owner")
	}

	if parcel.H3Index == 0 {
		return errors.New("bad index")
	}

	return repo.AddParcel(parcel)
}
