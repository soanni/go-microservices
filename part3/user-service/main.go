package main

import (
	"log"

	pb "github.com/soanni/go-microservices/part3/user-service/proto/user"
	micro "github.com/micro/go-micro"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "host=pgsql port=5432 user=shippy dbname=users password=shippy sslmode=disable")
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	// create table 'users' based on model pb.User{}
	db.Table("users").CreateTable(&pb.User{})

	srv := micro.NewService(
		micro.Name("shippy.service.user"),
	)

	srv.Init()

	repository := &UserRepository{db}
	h := &handler{repository}

	pb.RegisterUserServiceHandler(srv.Server(), h)

	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}