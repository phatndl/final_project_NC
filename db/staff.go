package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// AddStudent function
func AddStudent(student *Student) (interface{}, error) {
	collection := Client.Database(DbName).Collection(ColName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, student)
	return res, err
}

func UpdateStudent(student *Student) (interface{}, error) {
	collection := Client.Database(DbName).Collection(ColName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	fmt.Printf("[Update Student] - student: student ")
	filter := bson.M{"email": student.Email}
	update := bson.M{"$set": student}
	res, err := collection.UpdateOne(ctx, filter, update)
	return res, err
}
