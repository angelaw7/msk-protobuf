package main

import (
	model "angelaw7/msk-protobuf/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {
	jsonItem := &model.FetchJSON{
		SampleCount: 1,
		Disclaimer:  "Hello",
		Results: []*model.Result{
			{Tex: "Bob"},
		},
	}

	data, err := proto.Marshal(jsonItem)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
	ioutil.WriteFile("jsonItem.protobuf", data, 0600)

	data, err = json.Marshal(jsonItem)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("jsonItem.json", data, 0600)

	data, err = ioutil.ReadFile("jsonItem.protobuf")
	if err != nil {
		log.Fatal(err)
	}

	jsonItemFromFile := model.FetchJSON{}
	if err := proto.Unmarshal(data, &jsonItemFromFile); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("jsonItem from protobuf file %+v\n", jsonItemFromFile)

}
