package repository

import (
	"context"
	"strconv"

	"cloud.google.com/go/firestore"
	"github.com/michals92/wonderland-go/entity"
)

type firestoreRepo struct{}

const (
	projectID            string = "nofoto-52aa6"
	parcelCollectionName string = "parcels"
)

type Repository interface {
	GetParcels(box *entity.BoundingBox) (*[]entity.Parcel, error)
	AddParcel(parcel *entity.Parcel) error
}

func NewFirestoreRepository() Repository {
	return &firestoreRepo{}
}

func (r *firestoreRepo) GetParcels(box *entity.BoundingBox) (*[]entity.Parcel, error) {

	ctx := context.Background()
	client, error := firestore.NewClient(ctx, projectID)

	if error != nil {
		return nil, error
	}

	defer client.Close()

	//TODO: - limit to selected bounding box
	var parcels []entity.Parcel
	documents, err := client.Collection(parcelCollectionName).Documents(ctx).GetAll()

	if err != nil {
		return nil, err
	}

	for _, doc := range documents {
		var parcel entity.Parcel
		_ = doc.DataTo(&parcel)
		parcels = append(parcels, parcel)
	}

	return &parcels, nil
}

func (r *firestoreRepo) AddParcel(parcel *entity.Parcel) error {
	ctx := context.Background()
	client, error := firestore.NewClient(ctx, projectID)

	if error != nil {
		return error
	}

	defer client.Close()

	parcelDoc := client.Collection(parcelCollectionName).Doc(strconv.Itoa(parcel.H3Index))
	_, error = parcelDoc.Set(ctx, &parcel)

	return error
}
