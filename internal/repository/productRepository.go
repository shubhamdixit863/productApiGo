package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"productApi/internal/entity"
)

type ProductRepository struct {
	// Db connection would go

	Db *mongo.Database
}

func (pr *ProductRepository) SaveProduct(product entity.Product) error {

	_, err := pr.Db.Collection("movie").InsertOne(context.TODO(), product)
	if err != nil {
		return err
	}

	return nil

}

func (pr *ProductRepository) GetAll() ([]entity.Product, error) {
	var products []entity.Product
	fmt.Println("Service")
	filter := bson.D{}
	find, err := pr.Db.Collection("movie").Find(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for find.Next(context.TODO()) {

		// We will parse the data -----

		// We can unmarshall this string back to our struct
		b := entity.Product{}
		fmt.Println(find.Current.String())
		err := find.Decode(&b)

		//err := json.Unmarshal([]byte(find.Current.String()), &b)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		products = append(products, b)

	}

	fmt.Println(products)
	return products, nil
}
