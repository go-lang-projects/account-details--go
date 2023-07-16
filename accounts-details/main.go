package main

import (
	"fmt"
	"reflect"

	pb "ea.com/proto"
	"google.golang.org/protobuf/proto"
)

func generateAccount() *pb.Account {
	return &pb.Account{
		Id:       67,
		Name:     "Sourav",
		IsActive: true,
		Phones:   []int32{78788, 89989},
	}
}

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:          42,
		Name:        "My name",
		IsSimple:    true,
		SampleLists: []int32{1, 2, 3, 4, 5, 6},
	}
}

func generateTeam() *pb.Team {
	return &pb.Team{
		SingleMember: &pb.Details{Id: 42, Name: "Sachin Tend"},
		Members: []*pb.Details{
			{Id: 43, Name: "Sourav Roy"},
			{Id: 44, Name: "Rick Martin"},
		},
	}
}

func doOneOf(message interface{}) {
	switch x := message.(type) {
	case *pb.Identity_Id:
		fmt.Printf("This is an Id: %d\n", message.(*pb.Identity_Id).Id)
	case *pb.Identity_Name:
		fmt.Printf("This is a name : %s\n", message.(*pb.Identity_Name).Name)
	default:
		fmt.Printf("Identity has unexpected type: %T\n", x)
	}
}

func fetchUserIdMap() *pb.UserIdMap {
	message := &pb.UserIdMap{
		Ids: map[string]*pb.IdWrapper{
			"id1": {Id: 42},
			"id2": {Id: 43},
			"id3": {Id: 44},
		},
	}
	return message
}

func getEnumvalue() *pb.AccountStatusEnum {
	return &pb.AccountStatusEnum{
		ActiveStatusActive: pb.AccountStatus_ACTIVE_STATUS_ACTIVE,
		//ActiveStatusActive: 1,
	}
}

func doFile(p proto.Message) {
	path := "simple.bin"

	writeToFile(path, p)
	sm2 := &pb.Simple{}
	readFromFile(path, sm2)
	fmt.Println("Read the content:", sm2)
}

func doFromJSON(jsonString string, t reflect.Type) proto.Message {
	message := reflect.New(t).Interface().(proto.Message)
	fromJSON(jsonString, message)
	return message
}

func main() {
	fmt.Println(doSimple())
	fmt.Println(generateAccount())
	fmt.Println(generateTeam())
	fmt.Println(getEnumvalue())
	fmt.Println(fetchUserIdMap())
	doOneOf(&pb.Identity_Id{Id: 42})
	doOneOf(&pb.Identity_Name{Name: "Sourav"})
	fmt.Println(doFromJSON(toJSON(generateAccount()), reflect.TypeOf(pb.Account{})))
	fmt.Println(doFromJSON(toJSON(generateAccount()), reflect.TypeOf(pb.Account{})))
	fmt.Println(doFromJSON(`{"id": 42, "unknown": "test"}`, reflect.TypeOf(pb.Simple{})))
}
