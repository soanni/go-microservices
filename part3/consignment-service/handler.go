package main

import (
	"context"
	"log"

	pb "github.com/soanni/go-microservices/part3/consignment-service/proto/consignment"
	vesselProto "github.com/soanni/go-microservices/part3/vessel-service/proto/vessel"
)

type handler struct {
	repo repository
	vesselClient vesselProto.VesselServiceClient
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity: int32(len(req.Containers)),
	})

	if err != nil {
		return err
	}
	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
	req.VesselId = vesselResponse.Vessel.Id

	if err = s.repo.Create(req); err != nil {
		return err
	}

	res.Created = true
	res.Consignment = req

	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments, err := s.repo.GetAll()
	if err != nil {
		return err
	}
	res.Consignments = consignments
	return nil
}
