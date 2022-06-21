# msk-protobuf

A repo to create a prototype Golang ProtocolBuffer for MPath Sample

## Installation & Usage

### Local Development
To start local development, make sure you have `fetchjson.json` in the `msk-protobuf` directory and run the following:

```
git clone https://github.com/angelaw7/msk-protobuf.git
cd msk-protobuf
go run main.go
```

You may have to also run `export GO111MODULE=on` to enable module mode, and/or change the GOPATH using `export GOPATH=<directory>`

### Editing the Protocol Buffer
1. Make appropriate changes to `msk.proto`
2. Compile `msk.proto` to generate `msk.pb.go` by running:
```
protoc --go_out=. --go_opt=paths=source_relative model/msk.proto
```