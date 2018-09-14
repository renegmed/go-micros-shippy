// consignment-service/handler.go
package main

import (
	"context"
	"log"

	pb "github.com/renegmed/go-micros-shippy/consignment-service/proto/consignment"
	vesselProto "github.com/renegmed/go-micros-shippy/vessel-service/proto/vessel"
	mgo "gopkg.in/mgo.v2"
)

// Service should implement all of the methods to staisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures, etc.
// to give you a better idea.
type handler struct {
	session      *mgo.Session
	vesselClient vesselProto.VesselServiceClient
}

// If you use the master session, you are re-using the same socket.
// Which means your queries may become blocked by other queries and have
// to wait for operations to finish on this socket. Which is pointless
// in a language which supports concurrency.
func (s *handler) GetRepo() Repository {
	return &ConsignmentRepository{s.session.Clone()}
}

// CreateConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *handler) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	defer s.GetRepo().Close()

	// Here we call a client instance of our vessel service with our consignment weight,
	// and the amount of containers as the capacity value
	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
	if err != nil {
		return err
	}

	// We set the VesselId as the vessel we got back from our
	// vessel service
	req.VesselId = vesselResponse.Vessel.Id

	// Save our consignment
	err = s.GetRepo().Create(req)
	if err != nil {
		return err
	}

	// Return matching the `Response` message we created in our
	// protobuf definition.
	res.Created = true
	res.Consignment = req
	return nil
}

func (s *handler) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	defer s.GetRepo().Close()
	consignments, err := s.GetRepo().GetAll()
	if err != nil {
		return err
	}
	res.Consignments = consignments
	return nil
}
