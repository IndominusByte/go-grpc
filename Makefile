PACKAGE := $(shell head -1 go.mod | awk '{print $$2}')

# make proto DIR="greet/proto" DIR_BIN="greet"
proto:
	PATH="${PATH}:${HOME}/go/bin" \
			 protoc -I$(DIR) --go_opt=module=${PACKAGE} --go_out=paths=source_relative:./$(DIR) \
			 --go-grpc_opt=module=${PACKAGE} --go-grpc_out=. $(DIR)/*.proto
	go build -v -o bin/$(DIR_BIN)/server $(DIR_BIN)/server/*.go
	go build -v -o bin/$(DIR_BIN)/client $(DIR_BIN)/client/*.go
