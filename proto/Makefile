# /proto/Makefile
PROTO_DIR := .
OUT := ../services

GEN_SERVICES := rider-svc driver-svc trip-svc matching-svc

gen:
	@for svc in $(GEN_SERVICES); do \
		protoc -I $(PROTO_DIR) \
			--go_out=$(OUT)/$$svc/internal/gen --go_opt=paths=source_relative \
			--go-grpc_out=$(OUT)/$$svc/internal/gen --go-grpc_opt=paths=source_relative \
			common/v1/common.proto driver/v1/driver.proto trip/v1/trip.proto matching/v1/matching.proto; \
	done

lint:
	buf lint

.PHONY: gen lint
