package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	// Db connection would go

	Db *mongo.Database
}

func (pr *ProductRepository) GetAll() {

	fmt.Println("I am called from repo")

	find, err := pr.Db.Collection("movie").Find(context.TODO(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(find)
}
