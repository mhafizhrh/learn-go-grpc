package main

import (
	"bytes"
	"fmt"
	"learn-go-grpc/model/model"
	"log"

	"github.com/golang/protobuf/jsonpb"
)

func main() {
	user1 := &model.User{
		Id:       "u001",
		Name:     "Punklord Hacker",
		Password: "sUp3Rzecr3To",
		Gender:   model.UserGender_FEMALE,
	}

	userList := &model.UserList{
		List: []*model.User{user1},
	}

	// garage1 := &model.Garage{
	// 	Id:   "g001",
	// 	Name: "Sprightly Vonwacq",
	// 	Coordinate: &model.GarageCoordinate{
	// 		Latitude:  0.34567,
	// 		Longitude: 0.12890,
	// 	},
	// }

	// garageList := &model.GarageList{
	// 	List: []*model.Garage{garage1},
	// }

	// garageListByUser := &model.GarageListByUser{
	// 	List: map[string]*model.GarageList{
	// 		user1.Id: garageList,
	// 	},
	// }

	fmt.Printf("Original: %#v\n", user1)
	fmt.Printf("String: %s\n", user1.String())

	buf := new(bytes.Buffer)
	err := (&jsonpb.Marshaler{}).Marshal(buf, userList)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("JSON String: %v\n", buf.String())
}
