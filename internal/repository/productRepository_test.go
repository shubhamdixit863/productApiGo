package repository

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
)

func InitailizeConnection() *mongo.Database {

	uri := "mongodb://localhost:27017/movie"
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected With Mongodb")
	return client.Database("movie", nil)

}

func TestProductRepository_GetAll(t *testing.T) {

	mongnd := InitailizeConnection()

	pr := ProductRepository{mongnd}

	all, err := pr.GetAll()
	if err != nil {
		return
	}
	assert.Nil(t, err)
	assert.Greater(t, len(all), 0)

}
