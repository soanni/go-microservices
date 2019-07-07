package main

import (
	"context"
	pb "github.com/soanni/go-microservices/part3/vessel-service/proto/vessel"
)


type handler struct {
	repo repository 
}

func (s *handler) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return err
	}

	res.Vessel = vessel
	return nil
}

func (s *handler) CreateVessel(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	if err := s.repo.CreateVessel(req); err != nil {
		return err
	}
	res.Vessel = req
	res.Created = true
	return nil
}