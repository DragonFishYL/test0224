package main

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"sync"
)

var mu sync.RWMutex

const (
	MaxConnection     = 10
	InitialConnection = 4
	Available         = false
	Used              = true
)

type MongoData struct {
	Client *mongo.Client
	Pos    int
	Flag   bool
}

type ClientPool struct {
	ClientList [MaxConnection]MongoData
	size       int
}

//初始化 最好放在main
func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func (cpool *ClientPool) DbConnect() (*mongo.Client, error) {
	uri := os.Getenv("MONGODB")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable.")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return client, nil
}

func (cpool *ClientPool) AllowConnectToPool(pos int) error {
	var err error
	cpool.ClientList[pos].Client, err = cpool.DbConnect()
	if err != nil {
		return err
	}
	cpool.ClientList[pos].Flag = Used
	cpool.ClientList[pos].Pos = pos
	return nil
}

// GetClientToPool apply
func (cpool *ClientPool) GetClientToPool(pos int) {
	cpool.ClientList[pos].Flag = Used
}

// PutClientBackPool back
func (cpool *ClientPool) PutClientBackPool(pos int) {
	cpool.ClientList[pos].Flag = Available
}

// GetClient 获取数据库对象
func (cpool *ClientPool) GetClient() (*MongoData, error) {
	mu.RLock()
	pos := cpool.size
	for i := 0; i < pos; i++ {
		if cpool.ClientList[i].Flag == Available {
			return &cpool.ClientList[i], nil
		}
	}
	mu.RUnlock()
	mu.Lock()
	defer mu.Unlock()
	if pos < MaxConnection {
		if err := cpool.AllowConnectToPool(pos); err != nil {
			return nil, err
		}
		cpool.size++
		return &cpool.ClientList[pos], nil
	}
	return nil, nil
}

// SetClientReleasePool 回收资源
func (cpool *ClientPool) SetClientReleasePool(mongoData *MongoData) {
	mu.Lock()
	cpool.PutClientBackPool(mongoData.Pos)
	mu.Unlock()
}

// DeferClientClosePool 释放资源
func (cpool *ClientPool) DeferClientClosePool(mongoData *MongoData) error {
	defer func() {
		if err := mongoData.Client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	return nil
}
