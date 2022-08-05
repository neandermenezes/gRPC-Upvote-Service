build:
	protoc -Iproto --go_out=proto/pb --go_opt=paths=source_relative --go-grpc_out=proto/pb --go-grpc_opt=paths=source_relative proto/post.proto
	go build -o bin/server ./server
	go build -o bin/client ./client

prune:
	rm -f proto/pb/*.go