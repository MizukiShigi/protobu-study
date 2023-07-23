package main

import (
	"fmt"
	"log"
	"protobuf-lesson/pb"

	"github.com/golang/protobuf/jsonpb"
)

func main() {
	employee := &pb.Employee{
		Id:          1,
		Name:        "Suzuki",
		Email:       "test@example.com",
		Occupation:  pb.Occupation_ENGINEER,
		PhoneNumber: []string{"090-1234-5678", "080-1234,5678"},
		Project:     map[string]*pb.Company_Project{"ProjextX": &pb.Company_Project{}},
		Profile: &pb.Employee_Text{
			Text: "My name is Suzuki",
		},
		Birth: &pb.Date{
			Year:  2000,
			Month: 1,
			Day:   1,
		},
	}

	// バイナリ
	// binData, err := proto.Marshal(employee)
	// if err != nil {
	// 	log.Fatalln("Can't serialize", err)
	// }
	// if err := ioutil.WriteFile("test.bin", binData, 0666); err != nil {
	// 	log.Fatalln("Can't write", err)
	// }
	// in, err := ioutil.ReadFile("test.bin")
	// if err != nil {
	// 	log.Fatalln("Can't read file")
	// }
	// readEmployee := &pb.Employee{}

	// err = proto.Unmarshal(in, readEmployee)
	// if err != nil {
	// 	log.Fatalln("Can't desirialize")
	// }
	// fmt.Println(readEmployee)

	//  構造体 → JSON
	m := jsonpb.Marshaler{}
	out, err := m.MarshalToString(employee)
	if err != nil {
		log.Fatalln("can't marjal json", err)
	}
	// fmt.Println(out)

	// JSON → 構造体
	readEmployee := &pb.Employee{}
	if err := jsonpb.UnmarshalString(out, readEmployee); err != nil {
		log.Fatalln("can't unmarshal from json", err)
	}
	fmt.Println(readEmployee)
}
