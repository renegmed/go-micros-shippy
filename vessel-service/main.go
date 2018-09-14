// vessel-service/main.go
package main

import (
	"fmt"
	"log"
	"os"

	micro "github.com/micro/go-micro"
	pb "github.com/renegmed/go-micros-shippy/vessel-service/proto/vessel"
)

const (
	defaultHost = "localhost:27017"
)

func createDummyData(repo Repository) {
	defer repo.Close()
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500},
	}
	for _, v := range vessels {
		repo.Create(v)
	}
}

func main() {
	fmt.Println("Starting vessel server now...")

	// Database host from the environment variables
	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)

	// Mgo creates a 'master' session, we need to end that session
	// before the main function closes.
	defer session.Close()

	if err != nil {
		// We're wrapping the error returned from our CreateSession
		// here to add some context to the error
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}

	repo := &VesselRepository{session.Copy()}
	createDummyData(repo)

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	srv.Init()

	// Register our implementation with
	pb.RegisterVesselServiceHandler(srv.Server(), &service{session})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
