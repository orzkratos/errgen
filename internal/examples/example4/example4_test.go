package example4

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Test GetOrderResponse nested enum with separator
// 测试 GetOrderResponse 嵌套枚举（带分隔符）
func TestGetOrderErrors(t *testing.T) {
	// Test order not found
	erk := ErrorGetOrderResponse_OrderNotFound("order ORD-123 not found")
	require.True(t, IsGetOrderResponse_OrderNotFound(erk))
	t.Log(erk)

	// Test order expired
	erk = ErrorGetOrderResponse_OrderExpired("order expired")
	require.True(t, IsGetOrderResponse_OrderExpired(erk))
	t.Log(erk)

	// Test access denied
	erk = ErrorGetOrderResponse_AccessDenied("access denied")
	require.True(t, IsGetOrderResponse_AccessDenied(erk))
	t.Log(erk)
}

// Test function name with underscore separator
// 测试下划线分隔符的函数名
func TestSeparatedFunctionNames(t *testing.T) {
	erk := ErrorGetOrderResponse_Unknown("unknown error")
	t.Log("Function: ErrorGetOrderResponse_Unknown")
	require.NotNil(t, erk)

	erk = ErrorGetOrderResponse_AccessDenied("access denied")
	t.Log("Function: ErrorGetOrderResponse_AccessDenied")
	require.NotNil(t, erk)
}
