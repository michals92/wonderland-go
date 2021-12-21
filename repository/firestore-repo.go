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
	ParcelExists(parcelId int) bool
	AddPinnedNft(pin *entity.PinNft) error
	RemovePinnedNft(index int) error
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

func (r *firestoreRepo) ParcelExists(parcelId int) bool {
	ctx := context.Background()
	client, error := firestore.NewClient(ctx, projectID)

	if error != nil {
		return false
	}

	defer client.Close()

	parcelDoc, error := client.Collection(parcelCollectionName).Doc(strconv.Itoa(parcelId)).Get(ctx)

	if error != nil {
		return false
	}

	return parcelDoc != nil
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

func (r *firestoreRepo) AddPinnedNft(pin *entity.PinNft) error {
	ctx := context.Background()
	client, error := firestore.NewClient(ctx, projectID)

	if error != nil {
		return error
	}

	defer client.Close()

	_, error = client.Collection(parcelCollectionName).Doc(strconv.Itoa(pin.H3Index)).Update(ctx, []firestore.Update{
		{
			Path:  "PinnedNFT",
			Value: pin.PinnedNFT,
		},
	})

	return error
}

func (r *firestoreRepo) RemovePinnedNft(index int) error {
	ctx := context.Background()
	client, error := firestore.NewClient(ctx, projectID)

	if error != nil {
		return error
	}

	defer client.Close()

	_, error = client.Collection(parcelCollectionName).Doc(strconv.Itoa(index)).Update(ctx, []firestore.Update{
		{
			Path:  "PinnedNFT",
			Value: firestore.Delete,
		},
	})

	return error
}
