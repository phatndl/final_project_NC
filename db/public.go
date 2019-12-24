package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllStudents() (interface{}, error) {
	collection := Client.Database(DbName).Collection(ColName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Printf("get all students error: %v", err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	var students []Student
	err = res.All(ctx, &students)
	if err != nil {
		fmt.Printf("get all students error: %v", err)
	}

	fmt.Println("[Get Students] - students: ", students)
	return &students, err
}

func SearchStudentSimple(req *StudentSearchRequest) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{}
	bs, err := json.Marshal(req)
	err = json.Unmarshal(bs, &filter)
	if err != nil {
		log.Fatalf("log error: %v", err)
		return nil, err
	}

	for key, value := range filter {
		strCovert := fmt.Sprintf("%s", value)
		valueWithRegex := primitive.Regex{Pattern: strCovert, Options: "i"}
		fmt.Printf("strConvert: %v\n", valueWithRegex)
		filter[key] = valueWithRegex
	}

	fmt.Println("filter: ", filter)
	cur := Client.Database(DbName).Collection(ColName)
	res, err := cur.Find(ctx, filter)
	if err != nil {
		fmt.Printf("find error: %v", err)
		return nil, err
	}
	var students []Student
	err = res.All(ctx, &students)
	if err != nil {
		fmt.Printf("res all error: %v", err)
		return nil, err
	}

	return students, err

}
