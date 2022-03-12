gen:
	make clean
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb --experimental_allow_proto3_optional

clean:
	rm pb/*.go

server:
	go run server/main.go -port 8080

client:
	go run client/main.go -address 0.0.0.0:8080

