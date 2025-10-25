package example1_test

import (
	"fmt"
	"testing"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/orzkratos/errgen/internal/examples/example1"
	"github.com/orzkratos/errkratos/newerk"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	newerk.SetReasonCodeFieldName("numeric_reason_code_enum")
	m.Run()
}

func TestErrorUnknown(t *testing.T) {
	erk := example1.ErrorUnknown("unknown error occurred: %s", "system failure")
	t.Log(erk)

	require.NotNil(t, erk)
	require.Equal(t, int32(500), erk.Code)
	require.Equal(t, example1.ErrorReason_UNKNOWN.String(), erk.Reason)
	require.Equal(t, "unknown error occurred: system failure", erk.Message)
}

func TestErrorUserNotFound(t *testing.T) {
	erk := example1.ErrorUserNotFound("user with ID %d not found", 12345)
	t.Log(erk)

	require.NotNil(t, erk)
	require.Equal(t, int32(404), erk.Code)
	require.Equal(t, example1.ErrorReason_USER_NOT_FOUND.String(), erk.Reason)
	require.Equal(t, "user with ID 12345 not found", erk.Message)
}

func TestErrorInvalidParameter(t *testing.T) {
	erk := example1.ErrorInvalidParameter("invalid parameter: %s", "email format")
	t.Log(erk)

	require.NotNil(t, erk)
	require.Equal(t, int32(400), erk.Code)
	require.Equal(t, example1.ErrorReason_INVALID_PARAMETER.String(), erk.Reason)
	require.Equal(t, "invalid parameter: email format", erk.Message)
}

func TestIsUnknown(t *testing.T) {
	erk := example1.ErrorUnknown("test unknown")
	t.Log(erk)
	require.True(t, example1.IsUnknown(erk))

	userErk := example1.ErrorUserNotFound("test user")
	t.Log(userErk)
	require.False(t, example1.IsUnknown(userErk))

	stdErk := errors.New(500, "UNKNOWN", "standard unknown")
	t.Log(stdErk)
	require.True(t, example1.IsUnknown(stdErk))

	require.False(t, example1.IsUnknown(nil))
}

func TestIsUserNotFound(t *testing.T) {
	erk := example1.ErrorUserNotFound("test user not found")
	t.Log(erk)
	require.True(t, example1.IsUserNotFound(erk))

	unknownErk := example1.ErrorUnknown("test unknown")
	t.Log(unknownErk)
	require.False(t, example1.IsUserNotFound(unknownErk))

	stdErk := errors.New(404, "USER_NOT_FOUND", "standard user not found")
	t.Log(stdErk)
	require.True(t, example1.IsUserNotFound(stdErk))
}

func TestIsInvalidParameter(t *testing.T) {
	erk := example1.ErrorInvalidParameter("test invalid param")
	t.Log(erk)
	require.True(t, example1.IsInvalidParameter(erk))

	unknownErk := example1.ErrorUnknown("test unknown")
	t.Log(unknownErk)
	require.False(t, example1.IsInvalidParameter(unknownErk))

	stdErk := errors.New(400, "INVALID_PARAMETER", "standard invalid param")
	t.Log(stdErk)
	require.True(t, example1.IsInvalidParameter(stdErk))
}

func TestMetadataFieldName(t *testing.T) {
	erk := example1.ErrorUnknown("test with metadata")
	t.Log(erk)

	require.NotNil(t, erk.Metadata)

	reason, exists := erk.Metadata["numeric_reason_code_enum"]
	require.True(t, exists)

	expectedReason := fmt.Sprintf("%d", example1.ErrorReason_UNKNOWN.Number())
	require.Equal(t, expectedReason, reason)
}
