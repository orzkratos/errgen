// Package errgen provides compatibility wrapper around errgenkratos
// This package maintains API compatibility while delegating all functionality
// to the errgenkratos package
//
// errgen 提供 errgenkratos 的兼容包装
// 此包保持 API 兼容性，同时将所有功能委托给 errgenkratos 包
package errgen

import "github.com/orzkratos/errgenkratos"

// SetMetadataFieldName sets the global metadata field name
// Used to customize the metadata key that stores enum numeric values
//
// SetMetadataFieldName 设置全局元数据字段名
// 用于自定义存储枚举数值的元数据键名
func SetMetadataFieldName(fieldName string) {
	errgenkratos.SetMetadataFieldName(fieldName)
}

// GetMetadataFieldName gets the current metadata field name
// Returns the configured field name, defaults to empty string
//
// GetMetadataFieldName 获取当前元数据字段名
// 返回配置的字段名，默认为空字符串
func GetMetadataFieldName() string {
	return errgenkratos.GetMetadataFieldName()
}
