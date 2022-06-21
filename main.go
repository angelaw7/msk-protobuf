package main

import (
	model "angelaw7/msk-protobuf/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {

	// read data from JSON file (JSON wire-format encoding)
	data, err := ioutil.ReadFile("fetchjson.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("JSON encoded size: %d\n", len(data)) // 15,907,623

	// JSON -> ProtoMessage protobuf
	req := &model.FetchJSON{}
	if err := protojson.Unmarshal(data, req); err != nil {
		log.Fatal(err)
	}

	// protobuf -> ProtoMessage wire-format encoding
	data, err = proto.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("protobuf_encoding.txt", data, 0600)
	fmt.Printf("Protobuf encoded size: %d\n", len(data)) // 2,808,300

	// ProtoMessage wire-format encoding -> protobuf
	bookFromFile := &model.FetchJSON{}
	if err := proto.Unmarshal(data, bookFromFile); err != nil {
		log.Fatal(err)
	}

	// protobuf -> wire-format encoding (from ProtoMessage encoding rather than JSON encoding this time)
	data, err = proto.Marshal(bookFromFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Protobuf encoded size: %d\n", len(data)) // 2,808,300, same as before

	// protobuf -> JSON
	data, err = json.Marshal(bookFromFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("New JSON encoded size: %d\n", len(data)) // 6,357,384, smaller than previous JSON probably because null fields are now omitted

	// write JSON back to another file
	ioutil.WriteFile("new_json_file.json", data, 0600)
}
