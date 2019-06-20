package main

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/soanni/go-microservices/part3/vessel-service/proto/vessel"
	"go.mongodb.org/mongo-driver/mongo"
	import "go.mongodb.org/mongo-driver/bson"
)

type repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

type VesselRepository struct {
	vessels []*pb.Vessel
	collection *mongo.Collection
}

func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	filter := bson.
	for _, vessel := range repo.vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, errors.New("No vessel found by this spec")
}