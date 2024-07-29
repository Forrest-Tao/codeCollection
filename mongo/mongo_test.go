package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"testing"
)

var mongoClient *mongo.Client
var once sync.Once

func init() {
	once.Do(func() {
		client, err := mongo.Connect(context.Background(),
			options.Client().
				ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{
				Username: "root",
				Password: "root",
			}))
		if err != nil {
			panic(err)
		}
		mongoClient = client
	})
	err := mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
}

type Person struct {
	Id   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}

// TestInsertOne 使用 bson.D
func TestInsertOne(t *testing.T) {
	println("开始 TestInsertOne")
	defer func() {
		println("开始执行 defer")
		err := mongoClient.Disconnect(context.TODO())
		if err != nil {
			panic(err)
		}
	}()
	collection := mongoClient.Database("test").Collection("Person")
	person := bson.D{
		{"_id", "1"},
		{"name", "John"},
	}
	insertResult, err := collection.InsertOne(context.Background(), person)
	if err != nil {
		panic(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

// TestInsertMany 使用 []interface{}
func TestInsertMany(t *testing.T) {
	println("开始 TestInsertMany")
	defer func() {
		println("开始执行 defer")
		err := mongoClient.Disconnect(context.TODO())
		if err != nil {
			panic(err)
		}
	}()
	collection := mongoClient.Database("test").Collection("Person")
	persons := []interface{}{
		Person{"2", "Bill"},
		Person{"3", "Cat"},
		Person{"4", "Dog"},
		Person{"5", "Jack"},
	}
	insertResult, err := collection.InsertMany(context.Background(), persons)
	if err != nil {
		panic(err)
	}
	fmt.Println("Inserted multiple documents: ", insertResult.InsertedIDs)
}

// TestInsertMany2 使用bson.A
func TestInsertMany2(t *testing.T) {
	println("开始 TestInsertMany2")
	defer func() {
		println("开始执行 defer")
		err := mongoClient.Disconnect(context.TODO())
		if err != nil {
			panic(err)
		}
	}()
	collection := mongoClient.Database("test").Collection("Person")
	persons := bson.A{
		bson.D{
			{"_id", "10"},
			{"name", "John"},
		},
		bson.D{
			{"_id", "21"},
			{"name", "Bill"},
		},
	}
	insertResult, err := collection.InsertMany(context.Background(), persons)
	if err != nil {
		panic(err)
	}
	fmt.Println("Inserted multiple documents: ", insertResult.InsertedIDs)
}

func TestFind(t *testing.T) {
	println("开始 TestFind")
	defer func() {
		println("开始执行 defer")
		err := mongoClient.Disconnect(context.TODO())
		if err != nil {
			panic(err)
		}
	}()
	collection := mongoClient.Database("test").Collection("Person")
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}
	for cursor.Next(context.Background()) {
		var p Person
		if err = cursor.Decode(&p); err != nil {
			panic(err)
		}
		println(p.Name)
	}
}

func TestUpdateOne(t *testing.T) {
	println("开始 TestUpdateOne")
	defer func() {
		println("开始执行 defer")
		err := mongoClient.Disconnect(context.TODO())
		if err != nil {
			panic(err)
		}
	}()
	collection := mongoClient.Database("test").Collection("Person")
	filter := bson.D{
		{"name", "Bill"},
	}
	update := bson.D{
		{"$set", bson.D{{"name", "BILL"}}},
	}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		panic(err)
	}
}

func TestDeleteMany(t *testing.T) {
	println("开始 TestDeleteOne")
	defer func() {
		println("开始执行 defer")
		err := mongoClient.Disconnect(context.TODO())
		if err != nil {
			panic(err)
		}
	}()
	collection := mongoClient.Database("test").Collection("Person")
	filter := bson.D{
		{"name", bson.D{
			{"$in", []string{"BILL", "John"}},
		}},
	}

	one, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		panic(err)
	}
	println(one.DeletedCount)
}
