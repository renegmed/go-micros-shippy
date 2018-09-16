// user-server/repository.go
package main

import (
	"log"

	"github.com/jinzhu/gorm"
	pb "github.com/renegmed/go-micros-shippy/user-service/proto/user"
)

type Repository interface {
	GetAll() ([]*pb.User, error)
	Get(id string) (*pb.User, error)
	Create(user *pb.User) error
	GetByEmailAndPassword(user *pb.User) (*pb.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func (repo *UserRepository) GetAll() ([]*pb.User, error) {
	log.Println("Start getting all users ....")
	var users []*pb.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) Get(id string) (*pb.User, error) {
	log.Println("Start getting a user ....")
	var user *pb.User
	user.Id = id
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) Create(user *pb.User) error {
	log.Println("Start creating a user....")
	if err := repo.db.Create(&user).Error; err != nil {
		return err
	}
	log.Printf("Create user %v\n", user)
	return nil
}

func (repo *UserRepository) GetByEmailAndPassword(user *pb.User) (*pb.User, error) {
	log.Println("Start getting email and password ....")
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
