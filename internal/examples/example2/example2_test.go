package example2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNestedEnums(t *testing.T) {
	// Test nested enum in message
	erk := ErrorGetInfoRespUserNotFound("user not found")
	require.True(t, IsGetInfoRespUserNotFound(erk))
	t.Log(erk)
}

func TestTopLevelEnums(t *testing.T) {
	// Test top-level enum
	erk := ErrorUnknown("unknown error")
	require.True(t, IsUnknown(erk))
	t.Log(erk)

	erk = ErrorServiceUnavailable("service unavailable")
	require.True(t, IsServiceUnavailable(erk))
	t.Log(erk)
}

func TestEnumIsolation(t *testing.T) {
	// Test isolation between enums
	topErk := ErrorUnknown("top level error")
	nestedErk := ErrorGetInfoRespUserNotFound("nested error")

	require.False(t, IsGetInfoRespUserNotFound(topErk))
	require.False(t, IsUnknown(nestedErk))
}
