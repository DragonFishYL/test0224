package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	bson2 "gopkg.in/mgo.v2/bson"
	"log"
	"os"
)

type Inventory struct {
	Id     bson2.ObjectId `bson:"_id,omitempty"`
	Item   string
	Qty    interface{}
	Status string
	Size   map[string]interface{}
	Tags   []string
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable.")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database("local").Collection("inventory")
	//fmt.Printf("%+v", coll)
	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{"qty", 25}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return
		}
		panic(err)
	}
	fmt.Println("---------------------------")
	fmt.Printf("%+v", result)
	var (
		inventoryStruct Inventory
	)
	inventoryJson, err := json.Marshal(result)
	if err != nil {
		return
	}
	if err := json.Unmarshal([]byte(inventoryJson), &inventoryStruct); err != nil {
		return
	}
	fmt.Println("---------------------------")
	fmt.Printf("%+v", inventoryStruct)
}
