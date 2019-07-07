package main

import (
	"context"

	pb "github.com/soanni/go-microservices/part3/vessel-service/proto/vessel"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	CreateVessel(*pb.Vessel) error
}

type VesselsRepository struct {
	collection *mongo.Collection
}

func (repo *VesselsRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	var vessel *pb.Vessel
	// Here we define a more complex query than our consignment-service's
	// GetAll function. Here we're asking for a vessel who's max weight and
	// capacity are greater than or equal to the given capacity and weight.
	// We're also using the `One` function here as that's all we want.
	err := repo.collection.FindOne(context.Background(), bson.M{
		"capacity":  bson.M{"$gte": spec.Capacity},
		"maxweight": bson.M{"$gte": spec.MaxWeight},
	}).Decode(&vessel)
	if err != nil {
		return nil, err
	}
	return vessel, nil
}

func (repo *VesselsRepository) CreateVessel(vessel *pb.Vessel) error {
	_, err := repo.collection.InsertOne(context.Background(), vessel)
	return err	
}