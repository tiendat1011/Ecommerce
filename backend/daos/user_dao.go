package daos

import (
	"context"
	"ecommerce-project/databases"
	"ecommerce-project/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserDAO struct {
	collection *mongo.Collection
}

const (
	USER_COLLECTION = "users"
)

func NewUserDAO() *UserDAO {
	if databases.DB == nil {
		panic("Not connect db yet")
	}
	return &UserDAO{
		collection: databases.DB.Collection(USER_COLLECTION),
	}
}

func (d *UserDAO) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := d.collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (d *UserDAO) CreateUser(user *models.User) (*models.User, error) {
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if _, err := d.collection.InsertOne(context.TODO(), user); err != nil {
		return nil, err
	}

	return user, nil
}