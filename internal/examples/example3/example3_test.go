package example3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Test top-level and nested error generation
// 测试顶层和嵌套错误生成
func TestErrorGeneration(t *testing.T) {
	// Top-level error from errors.proto
	topErk := ErrorUnknown("top error")
	require.True(t, IsUnknown(topErk))
	t.Log(topErk)

	// Nested error from user_service.proto
	nestedErk := ErrorGetUserResponseUserNotFound("user not found")
	require.True(t, IsGetUserResponseUserNotFound(nestedErk))
	t.Log(nestedErk)

	// Verify isolation
	require.False(t, IsGetUserResponseUserNotFound(topErk))
	require.False(t, IsUnknown(nestedErk))
}

// Test same-name enum behavior
// 测试同名枚举的行为
func TestSameNameEnum(t *testing.T) {
	topErk := ErrorUnknown("top")
	nestedErk := ErrorGetUserResponseUnknown("nested")

	// Same reason="UNKNOWN" makes them equivalent
	require.True(t, IsUnknown(topErk))
	require.True(t, IsUnknown(nestedErk))
}
