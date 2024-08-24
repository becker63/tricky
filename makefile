generate:
	protoc --proto_path=pb pb/*.proto --go_out=. --go-grpc_out=.
