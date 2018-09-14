// vessel-service/repository.go
package main

import (
	pb "github.com/renegmed/go-micros-shippy/vessel-service/proto/vessel"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName           = "shippy"
	vesselCollection = "vessels"
)

type Repository interface {
	Create(*pb.Vessel) error
	GetAll() ([]*pb.Vessel, error)
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Close()
}

type VesselRepository struct {
	session *mgo.Session
}

// Create a new consignment
func (repo *VesselRepository) Create(vessel *pb.Vessel) error {
	return repo.collection().Insert(vessel)
}

// GetAll consignments
func (repo *VesselRepository) GetAll() ([]*pb.Vessel, error) {
	var vessels []*pb.Vessel
	// Find normally takes a query, but as we want everything, we can nil this.
	// We then bind our consignments variable by passing it as an argument to .All().
	// That sets consignments to the result of the find query.
	// There's also a `One()` function for single results.
	err := repo.collection().Find(nil).All(&vessels)
	return vessels, err
}

// FindAvailable - checks a specification against a map of vessels,
// if capacity and max weight are below a vessels capacity and max weight,
// then return that vessel.
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	var vessel *pb.Vessel

	// Here we define  a more complex query than our consignment-services's
	// GetAll function. Here we're asking for a vessel who's max weight and
	// capacity are greater than and equal to the given capacity and weight.
	// We're also using the `One` function here as that's all we want.

	err := repo.collection().Find(bson.M{
		"capacity":  bson.M{"$gte": spec.Capacity},
		"maxweight": bson.M{"$gte": spec.MaxWeight},
	}).One(&vessel)
	if err != nil {
		return nil, err
	}
	return vessel, nil
}

func (repo *VesselRepository) Close() {
	repo.session.Close()
}

func (repo *VesselRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(vesselCollection)
}
