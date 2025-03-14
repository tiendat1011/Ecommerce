package daos

import (
	"ecommerce-project/databases"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

const (
	PRODUCT_COLLECTION = "products"
)

type ProductDAO interface {

}

type productDAO struct {
	collection *mongo.Collection
}

func NewProductDAO() *productDAO {
	return &productDAO{
		collection: databases.DB.Collection(PRODUCT_COLLECTION),
	}
}