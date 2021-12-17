package repository

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"github.com/michals92/wonderland-go/entity"
)

type firestoreRepo struct{}

const (
	projectID            string = "nofoto-52aa6"
	parcelCollectionName string = "parcels"
)

type Repository interface {
	GetParcels(box *entity.BoundingBox) ([]entity.Parcel, error)
	AddParcel(parcel *entity.Parcel) error
}

func NewFirestoreRepository() Repository {
	return &firestoreRepo{}
}

func (r *firestoreRepo) GetParcels(box *entity.BoundingBox) ([]entity.Parcel, error) {

	ctx := context.Background()
	client, error := firestore.NewClient(ctx, projectID)

	if error != nil {
		return nil, error
	}

	defer client.Close()

	//TODO: - obtain parcels in selected area
	return nil, errors.New("parcels repo not impelemented")
}

func (r *firestoreRepo) AddParcel(parcel *entity.Parcel) error {
	ctx := context.Background()
	client, error := firestore.NewClient(ctx, projectID)

	if error != nil {
		return error
	}

	defer client.Close()

	parcelDoc := client.Collection(parcelCollectionName).Doc(parcel.H3Index)
	_, error = parcelDoc.Set(ctx, &parcel)

	return error
}
