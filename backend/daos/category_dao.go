package daos

import (
	"context"
	"ecommerce-project/databases"
	"ecommerce-project/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type CategoryDAO struct {
	collection *mongo.Collection
}

const (
	CATEGORY_COLLECTION = "category"
)

func NewCategoryDAO() *CategoryDAO {
	return &CategoryDAO{
		collection: databases.DB.Collection(CATEGORY_COLLECTION),
	}
}

func (d *CategoryDAO) GetCategoryByName(name string) (*models.Category, error) {
	var category models.Category

	err := d.collection.FindOne(context.TODO(), bson.M{"name": name}).Decode(&category)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found category")
	}

	return &category, nil
}

func (d *CategoryDAO) CreateCategory(nc *models.Category) (*models.Category, error) {
	nc.ID = primitive.NewObjectID()
	if _, err := d.collection.InsertOne(context.TODO(), nc); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Cannot add category")
	}

	return nc, nil
}

func (d *CategoryDAO) UpdateCategory(ur *models.UpdateCategoryRequest, id string) error {
	update := bson.M{}

	update["name"] = ur.Name

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