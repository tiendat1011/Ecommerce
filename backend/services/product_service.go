package services

import "ecommerce-project/daos"

type ProductService interface {
	
}

type productSerivce struct {
	productDAO daos.ProductDAO
}

func NewProductService(productDAO daos.ProductDAO) *productSerivce {
	return &productSerivce{
		productDAO: productDAO,
	}
}
