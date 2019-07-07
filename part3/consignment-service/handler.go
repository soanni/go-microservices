package main

import (
	"context"
	"log"
	"fmt"

	pb "github.com/soanni/go-microservices/part3/consignment-service/proto/consignment"
	vesselProto "github.com/soanni/go-microservices/part3/vessel-service/proto/vessel"
)

type handler struct {
	repo repository
	vesselClient vesselProto.VesselServiceClient
}

func (s *handler) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	fmt.Println("*** Inside handler.CreateConsignment ***")
	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity: int32(len(req.Containers)),
	})

	if err != nil {
		return err
	}
	log.Printf("Found veSSSel: %s \n", vesselResponse.Vessel.Name)
	log.Printf("Found vessel id: %s \n", vesselResponse.Vessel.Id)

	req.VesselId = vesselResponse.Vessel.Id

	if err = s.repo.Create(req); err != nil {
		return err
	}

	res.Created = true
	res.Consignment = req

	return nil
}

func (s *handler) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	fmt.Println("*** Inside handler.GetConsignments ***")
	consignments, err := s.repo.GetAll()
	if err != nil {
		return err
	}
	res.Consignments = consignments
	return nil
}
