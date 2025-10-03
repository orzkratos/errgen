[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/orzkratos/errgen/release.yml?branch=main&label=BUILD)](https://github.com/orzkratos/errgen/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/orzkratos/errgen)](https://pkg.go.dev/github.com/orzkratos/errgen)
[![Coverage Status](https://img.shields.io/coveralls/github/orzkratos/errgen/main.svg)](https://coveralls.io/github/orzkratos/errgen?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/orzkratos/errgen.svg)](https://github.com/orzkratos/errgen/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/orzkratos/errgen)](https://goreportcard.com/report/github.com/orzkratos/errgen)

# errgen

Backward-compatible package providing the classic protoc-gen-go-errors command, built on errgenkratos implementation.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->
## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Overview

This project maintains backward compatible interface with existing protoc-gen-go-errors, delegating implementation to [errgenkratos](https://github.com/orzkratos/errgenkratos). 

**New projects should use [errgenkratos](https://github.com/orzkratos/errgenkratos) to access the newest capabilities and enhancements.**

## Installation

```bash
go install github.com/orzkratos/errgen/cmd/protoc-gen-go-errors@latest
```

## Usage

### Standard Generation (Top Enums)
```bash
protoc --go-errors_out=paths=source_relative:./your_output_dir your_proto_files.proto
```

### With Nested Enum Support

Enable `include_nested=true` to generate helpers when enums nest inside messages:

```bash
protoc --go-errors_out=include_nested=true,paths=source_relative:./your_output_dir your_proto_files.proto
```

### With Underscore Naming

```bash
protoc --go-errors_out=include_nested=true,separate_named=true,paths=source_relative:./your_output_dir your_proto_files.proto
```

### Generated Code Example

```go
// Generated from protoc-gen-go-errors

// User not found
func IsUserNotFound(err error) bool {
    return errgenkratos.IsError(err, ErrorReason_USER_NOT_FOUND, 404)
}

// User not found  
func ErrorUserNotFound(format string, args ...interface{}) *errors.Error {
    return errgenkratos.NewError(404, ErrorReason_USER_NOT_FOUND, format, args...)
}
```

### Setting Custom Metadata Field

Configure metadata field in the business code:

```go
import "github.com/orzkratos/errgen"

func init() {
    // Set custom metadata field name
    errgen.SetMetadataFieldName("codeNum")
}
```

## Examples

Check out the [examples](internal/examples/) DIR:
- [example1](internal/examples/example1/) - Basic top enum error generation
- [example2](internal/examples/example2/) - Single file with nested enum support
- [example3](internal/examples/example3/) - Multi-file project with service definitions
- [example4](internal/examples/example4/) - Nested enum using underscore_names

## Migration to errgenkratos

To access new features and active development, migrate to [errgenkratos](https://github.com/orzkratos/errgenkratos):

```bash
# Install the new plugin
go install github.com/orzkratos/errgenkratos/cmd/protoc-gen-orzkratos-errors@latest

# Update the protoc command
protoc --orzkratos-errors_out=paths=source_relative:./your_output_dir your_proto_files.proto

# Update the imports
import "github.com/orzkratos/errgenkratos"
```

## Requirements

- Go 1.25.0+
- protoc (ProtoBuf code gen)
- Kratos v2.8.0+

---

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-26 07:39:27.188023 +0000 UTC -->

## 📄 License

MIT License. See [LICENSE](LICENSE).

---

## 🤝 Contributing

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Found a mistake?** Open an issue on GitHub with reproduction steps
- 💡 **Have a feature idea?** Create an issue to discuss the suggestion
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes and use significant commit messages
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/orzkratos/errgen.svg?variant=adaptive)](https://starchart.cc/orzkratos/errgen)