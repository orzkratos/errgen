// Package errgen provides compatibility wrapper for errgenkratos.
// This package maintains API compatibility while delegating all functionality
// to the errgenkratos package.
package errgen

import "github.com/orzkratos/errgenkratos"

// SetMetadataFieldName sets the global metadata field name
func SetMetadataFieldName(fieldName string) {
	errgenkratos.SetMetadataFieldName(fieldName)
}

// GetMetadataFieldName returns the current metadata field name
func GetMetadataFieldName() string {
	return errgenkratos.GetMetadataFieldName()
}
