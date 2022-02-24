package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

func Conn(dbName string, collName string) *mongo.Collection {
	con := ClientPool{}
	conn, err := con.GetClient()
	if err != nil {
		panic(err)
	}
	inv := conn.Client.Database("local").Collection("inventory")
	return inv
}

func TestFindOne(t *testing.T) {
	inv := Conn("local", "inventory")
	var result bson.M
	if err := inv.FindOne(context.TODO(), bson.D{{"qty", 77}}).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return
		}
		panic(err)
	}
	//defer con.SetClientReleasePool(conn)
	fmt.Println("---------------------------")
	fmt.Printf("%+v", result)
	fmt.Println("---------------------------")
}
func TestFind(t *testing.T) {
	inv := Conn("local", "inventory")

	reObj, err := inv.Find(
		context.TODO(),
		bson.M{
			"qty": bson.M{
				"$gt": 20,
				"$lt": 80,
			},
		},
		options.Find().SetLimit(10),
		options.Find().SetSort(
			bson.M{
				"qty": -1,
			},
		),
	)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return
		}
		panic(err)
	}
	var (
		inventory     Inventory
		inventorySlic = make([]Inventory, 0)
	)
	for reObj.Next(context.TODO()) {
		if err = reObj.Decode(&inventory); err != nil {
			panic(err)
		}
		inventorySlic = append(inventorySlic, inventory)
	}
	fmt.Printf("%+v--------%+v", inventorySlic, len(inventorySlic))

}

func TestInsert(t *testing.T) {
	inv := Conn("local", "inventory")
	//创建数据
	data := Inventory{
		Item:   "hello",
		Qty:    66,
		Status: "BB",
		Size:   map[string]interface{}{"h": 11, "ww": 11, "uom": 11},
		Tags:   []string{"block", "yyy"},
	}
	obj, err := inv.InsertOne(context.TODO(), data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", obj)
}

func TestInsertMany(t *testing.T) {
	inv := Conn("local", "inventory")
	data := make([]interface{}, 0)
	data = []interface{}{
		Inventory{
			Item:   "hello7",
			Qty:    772,
			Status: "CC",
			Size:   map[string]interface{}{"h": 117, "ww": 117, "uom": 117},
			Tags:   []string{"block7", "yyy7"},
		},
		Inventory{
			Item:   "hello8",
			Qty:    882,
			Status: "CC",
			Size:   map[string]interface{}{"h": 118, "ww": 118, "uom": 118},
			Tags:   []string{"block8", "yyy8"},
		},
	}
	obj, err := inv.InsertMany(context.TODO(), data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", obj)
}

func TestUpdateOne(t *testing.T) {
	inv := Conn("local", "inventory")
	obj, err := inv.UpdateOne(
		context.TODO(),
		bson.M{"qty": 77},
		bson.M{
			"$set": bson.M{
				"status": "CC77",
			},
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", obj)
}

func TestUpdateMany(t *testing.T) {
	inv := Conn("local", "inventory")
	obj, err := inv.UpdateMany(
		context.TODO(),
		bson.M{
			"item": "hello7",
		},
		bson.M{
			"$set": bson.M{
				"status": "CC7",
			},
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", obj)
}

func TestCount(t *testing.T) {
	inv := Conn("local", "inventory")
	obj, err := inv.CountDocuments(
		context.TODO(),
		bson.M{
			"qty": bson.M{
				"$gt": 30,
				"$lt": 60,
			},
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", obj)
}
