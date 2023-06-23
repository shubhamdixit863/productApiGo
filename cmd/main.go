package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"productApi/cmd/api"
	"productApi/internal/repository"
	"productApi/internal/services"
)

// all the code for starting of our project
// this would be called before the main function
func init() {
	// If you
	//fmt.Println("hiii")

}

func MongoConnection() *mongo.Database {
	//	uri := os.Getenv("MONGODB_URI")
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

func main() {

	db := MongoConnection()
	productRepository := &repository.ProductRepository{Db: db}
	productService := &services.ProductService{ProductRepository: productRepository} // Dependency injection
	handler := api.Handler{ProductService: productService}

	// Creating up a router

	r := mux.NewRouter()

	r.HandleFunc("/", handler.GetProducts)                       // root route
	r.HandleFunc("/product", handler.AddProduct).Methods("POST") // this api will only accept post request
	r.HandleFunc("/product/{id}", handler.GetProductById).Methods("GET")
	r.HandleFunc("/product/{id}", handler.DeleteProductById).Methods("DELETE")

	fmt.Println("Server running at port 8080")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalln(err)
	}

}
