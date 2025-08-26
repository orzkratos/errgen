PROTO_FILES=$(shell find . -name *.proto)

.PHONY: generate-branded
generate-branded:
	protoc --proto_path=. \
	       --proto_path=../../cmd/protoc-gen-orzkratos-errors \
	       --go_out=paths=source_relative:. \
	       --orzkratos-errors_out=paths=source_relative:. \
	       $(PROTO_FILES)

.PHONY: generate-compatible  
generate-compatible:
	protoc --proto_path=. \
	       --proto_path=../../cmd/protoc-gen-go-errors \
	       --go_out=paths=source_relative:. \
	       --go-errors_out=paths=source_relative:. \
	       $(PROTO_FILES)

.PHONY: generate
generate: generate-compatible

.PHONY: clean
clean:
	rm -f *.pb.go *_errors.pb.go