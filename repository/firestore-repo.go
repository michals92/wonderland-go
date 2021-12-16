package repository

import (
	"github.com/michals92/wonderland-go/entity"
)

type firestoreRepo struct{}

const (
	projectID          string = "nofoto-52aa6"
	userCollectionName string = "users"
)

type Repository interface {
	CreateUser(user *entity.User) error
}

func NewFirestoreRepository() Repository {
	return &firestoreRepo{}
}

func (r *firestoreRepo) CreateUser(user *entity.User) error {

	//TODO: - implement test object for user
	/*ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)

	if err != nil {
		return err
	}

	defer client.Close()

	_, err = client.Collection(userCollectionName).Doc(user.Email).Get(ctx)

	if err == nil {
		return errors.New("user with this email already exists")
	}

	_, err = client.Collection(userCollectionName).Doc(user.Email).Create(ctx, &user)*/
	return nil
}
