package main

import (
	"context"
	"log"
	"fmt"

	pb "github.com/soanni/go-microservices/part3/consignment-service/proto/consignment"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type repository interface {
	Create(*pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
}

type ConsignmentsRepository struct {
	collection *mongo.Collection
}

func (repository *ConsignmentsRepository) Create(consignment *pb.Consignment) error {
	fmt.Println("*** Inside repository.Create ***")
	_, err := repository.collection.InsertOne(context.Background(), consignment)
	return err
}

func (repository *ConsignmentsRepository) GetAll() ([]*pb.Consignment, error) {
	fmt.Println("*** Inside repository.GetAll ***")
	cur, err := repository.collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	var consignments []*pb.Consignment

	for cur.Next(context.Background()) {
		fmt.Println("*** Inside repository.GetAll.cursor ***")
		var consignment *pb.Consignment
		if err := cur.Decode(&consignment); err != nil {
			return nil, err
		}

		consignments = append(consignments, consignment)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return consignments, err
}
