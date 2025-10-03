PROTO_FILES=$(shell find . -name "*.proto")

.PHONY: install
install:
	go install github.com/orzkratos/errgen/cmd/protoc-gen-go-errors@latest
	@echo "protoc-gen-go-errors 安装完成!"

.PHONY: generate
generate:
	protoc --proto_path=. \
	       --proto_path=../../proto3ps \
	       --go_out=paths=source_relative:. \
	       --go-errors_out=include_nested=true,paths=source_relative:. \
	       $(PROTO_FILES)
	@echo "错误代码生成完成（包含顶级和嵌套枚举）!"

.PHONY: clean
clean:
	rm -f *.pb.go *_errors.pb.go
	@echo "清理生成文件完成!"

.PHONY: help
help:
	@echo "Available targets:"
	@echo "  install  - Install protoc-gen-go-errors plugin"
	@echo "  generate - Generate Go code from proto files (include nested enums)"
	@echo "  clean    - Remove generated .pb.go files"
	@echo "  help     - Show this help message"