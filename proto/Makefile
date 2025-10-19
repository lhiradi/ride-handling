# /proto/Makefile
PROTO_DIR := .
OUT := .

PROTO_FILES := common/v1/common.proto driver/v1/driver.proto trip/v1/trip.proto matching/v1/matching.proto

# Generate Go code for all protos in this folder.
# Requires:
#   - protoc (the protobuf compiler)
#   - protoc-gen-go  (go install google.golang.org/protobuf/cmd/protoc-gen-go@latest)
#   - protoc-gen-go-grpc (go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest)
gen:
	@protoc -I $(PROTO_DIR) \
		--go_out=$(OUT) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT) --go-grpc_opt=paths=source_relative \
		$(PROTO_FILES)

lint:
	buf lint

.PHONY: gen lint
