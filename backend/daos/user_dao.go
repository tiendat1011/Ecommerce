package daos

import (
	"context"
	"ecommerce-project/databases"
	"ecommerce-project/models"
	"ecommerce-project/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserDAO interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(id string) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	UpdateUser(ur *models.UpdateRequest, id string) error
	DeleteUser(id string) error
}

type userDAO struct {
	collection *mongo.Collection
}

const (
	USER_COLLECTION = "users"
)

func NewUserDAO() *userDAO {
	if databases.DB == nil {
		panic("Not connect db yet")
	}
	return &userDAO{
		collection: databases.DB.Collection(USER_COLLECTION),
	}
}

func (d *userDAO) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := d.collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (d *userDAO) GetUserById(id string) (*models.User, error) {
	var user models.User

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "invalid user ID format")
	}

	err = d.collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "cannot find the user")
	}

	return &user, nil
}

func (d *userDAO) CreateUser(user *models.User) (*models.User, error) {
	user.ID = primitive.NewObjectID()
	user.IsAdmin = false
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if _, err := d.collection.InsertOne(context.TODO(), user); err != nil {
		return nil, err
	}

	return user, nil
}

func (d *userDAO) GetAllUsers() ([]*models.User, error) {
	cursor, err := d.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var users []*models.User
	if err = cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (d *userDAO) UpdateUser(ur *models.UpdateRequest, id string) error {
	update := bson.M{}

	if ur.Username != "" {
		update["username"] = ur.Username
	}
	if ur.Email != "" {
		update["email"] = ur.Email
	}
	if ur.Password != "" {
		update["password"], _ = utils.HashPassword(ur.Password) // Hash password
	}

	update["is_admin"] = ur.IsAdmin

	if len(update) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "No fields to update")
	}

	updated := bson.M{"$set": update}

	objID, err := primitive.ObjectIDFromHex(id)

	result, err := d.collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, updated)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update user")
	}

	if result.MatchedCount == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Something went wrong when saving")
	}

	return nil
}

func (d *userDAO) DeleteUser(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "invalid user ID format")
	}

	result, err := d.collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update user")
	}

	if result.DeletedCount == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Something went wrong when deleting")
	}
	return nil
}
