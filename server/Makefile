server:
	go run main.go

proto:
	rm -f pb/*.go 
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=../proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=react_go \
	../proto/*.proto --experimental_allow_proto3_optional
	statik -src=./doc/swagger -dest=./doc

evans:
	chmod +x ./evans
	./evans --host localhost --port 9090 --package pb --service ReactGo -r repl

.PHONY: server proto evans