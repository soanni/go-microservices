package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/soanni/go-microservices/part3/vessel-service/proto/vessel"
	micro "github.com/micro/go-micro"
)

const (
	defaultHost = "mongodb://datastore:27017"
)

func createDummyData(repo *VesselsRepository) {
	vessels := []*pb.Vessel{
		&pb.Vessel{Id: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500},
	}
	for _, v := range vessels {
		repo.CreateVessel(v)
	}
}

func main() {
	srv := micro.NewService(
		micro.Name("shippy.service.vessel"),
	)

	srv.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(uri)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.TODO())

	vesselsCollection := client.Database("shippy").Collection("vessels")
	repository := &VesselsRepository{vesselsCollection}
	h := &handler{repository}

	createDummyData(repository)

	pb.RegisterVesselServiceHandler(srv.Server(), h)

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}