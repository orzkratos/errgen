PROTO_FILES=$(shell find . -name *.proto)

.PHONY: generate
generate:
	protoc --proto_path=. \
	       --proto_path=../../.. \
	       --proto_path=../../proto3ps \
	       --go_out=paths=source_relative:. \
	       --go-errors_out=paths=source_relative:. \
	       $(PROTO_FILES)

.PHONY: clean
clean:
	rm -f *.pb.go *_errors.pb.go

.PHONY: help
help:
	@echo "Available targets:"
	@echo "  generate - Generate Go code from proto files using protoc-gen-go-errors"
	@echo "  clean    - Remove generated .pb.go files"
	@echo "  help     - Show this help message"