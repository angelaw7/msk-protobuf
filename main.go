package main

import (
	model "angelaw7/msk-protobuf/model"
	"encoding/json"
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {

	// read data from JSON file
	data, err := ioutil.ReadFile("fetchjson.json")
	if err != nil {
		log.Fatal(err)
	}

	// JSON -> protobuf
	req := &model.FetchJSON{}
	if err := protojson.Unmarshal(data, req); err != nil {
		log.Fatal(err)
	}

	// protobuf -> wire format encoding
	data, err = proto.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}

	// wire format -> protobuf
	bookFromFile := &model.FetchJSON{}
	if err := proto.Unmarshal(data, bookFromFile); err != nil {
		log.Fatal(err)
	}

	// protobuf -> wire format encoding
	data, err = proto.Marshal(bookFromFile)
	if err != nil {
		log.Fatal(err)
	}

	// protobuf -> JSON
	data, err = json.Marshal(bookFromFile)
	if err != nil {
		log.Fatal(err)
	}

	// write JSON back to another file
	ioutil.WriteFile("next.json", data, 0600)
}
