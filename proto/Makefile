# /proto/Makefile
# Generate Go files into a shared location that matches the go_package option.
# This makes them importable via the Go module path.
# The command should be run from the project root directory (e.g., `make -f proto/Makefile gen`)
PROTO_DIR := proto
OUT_DIR := .

.PHONY: gen
gen:
	protoc -I $(PROTO_DIR) \
		--go_out=$(OUT_DIR) \
		--go-grpc_out=$(OUT_DIR) \
		$(PROTO_DIR)/common/v1/common.proto \
		$(PROTO_DIR)/driver/v1/driver.proto \
		$(PROTO_DIR)/trip/v1/trip.proto \
		$(PROTO_DIR)/matching/v1/matching.proto
