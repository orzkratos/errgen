package errgen

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// ProtoErrorEnum constraint for protobuf generated enum types
type ProtoErrorEnum interface {
	String() string
	Number() protoreflect.EnumNumber
}

// Config for error generation
type Config struct {
	MetadataFieldName string
}

var defaultConfig = &Config{
	MetadataFieldName: "", // default: no metadata
}

// SetMetadataFieldName sets the global metadata field name
func SetMetadataFieldName(fieldName string) {
	defaultConfig.MetadataFieldName = fieldName
}

// GetMetadataFieldName returns the current metadata field name
func GetMetadataFieldName() string {
	return defaultConfig.MetadataFieldName
}

// IsError checks if the given error matches the specified enum and HTTP code
func IsError[E ProtoErrorEnum](err error, enum E, httpCode int) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == enum.String() && e.Code == int32(httpCode)
}

// NewError creates a new error with the specified parameters
func NewError[E ProtoErrorEnum](httpCode int, enum E, format string, args ...interface{}) *errors.Error {
	erk := errors.New(httpCode, enum.String(), fmt.Sprintf(format, args...))
	if defaultConfig.MetadataFieldName != "" {
		erk = erk.WithMetadata(map[string]string{
			defaultConfig.MetadataFieldName: fmt.Sprintf("%d", enum.Number()),
		})
	}
	return erk
}
