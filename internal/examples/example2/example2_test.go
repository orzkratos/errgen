package example2_test

import (
	"fmt"
	"testing"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/orzkratos/errgen/internal/examples/example2"
	"github.com/orzkratos/errgenkratos"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	errgenkratos.SetMetadataFieldName("api_error_code")
	m.Run()
}

func TestErrorServiceUnavailable(t *testing.T) {
	erk := example2.ErrorServiceUnavailable("service %s is temporarily unavailable", "payment")
	t.Log(erk)

	require.NotNil(t, erk)
	require.Equal(t, int32(500), erk.Code)
	require.Equal(t, example2.ApiError_SERVICE_UNAVAILABLE.String(), erk.Reason)
	require.Equal(t, "service payment is temporarily unavailable", erk.Message)
}

func TestErrorAuthFailed(t *testing.T) {
	erk := example2.ErrorAuthFailed("authentication failed for user %s", "admin")
	t.Log(erk)

	require.NotNil(t, erk)
	require.Equal(t, int32(401), erk.Code)
	require.Equal(t, example2.ApiError_AUTH_FAILED.String(), erk.Reason)
	require.Equal(t, "authentication failed for user admin", erk.Message)
}

func TestErrorPermissionDenied(t *testing.T) {
	erk := example2.ErrorPermissionDenied("access denied to resource %s", "/admin/users")
	t.Log(erk)

	require.NotNil(t, erk)
	require.Equal(t, int32(403), erk.Code)
	require.Equal(t, example2.ApiError_PERMISSION_DENIED.String(), erk.Reason)
	require.Equal(t, "access denied to resource /admin/users", erk.Message)
}

func TestErrorRateLimitExceeded(t *testing.T) {
	erk := example2.ErrorRateLimitExceeded("rate limit exceeded: %d requests per minute", 1000)
	t.Log(erk)

	require.NotNil(t, erk)
	require.Equal(t, int32(429), erk.Code)
	require.Equal(t, example2.ApiError_RATE_LIMIT_EXCEEDED.String(), erk.Reason)
	require.Equal(t, "rate limit exceeded: 1000 requests per minute", erk.Message)
}

func TestIsServiceUnavailable(t *testing.T) {
	erk := example2.ErrorServiceUnavailable("test service down")
	t.Log(erk)
	require.True(t, example2.IsServiceUnavailable(erk))

	authErk := example2.ErrorAuthFailed("test auth")
	t.Log(authErk)
	require.False(t, example2.IsServiceUnavailable(authErk))

	stdErk := errors.New(500, "SERVICE_UNAVAILABLE", "standard service unavailable")
	t.Log(stdErk)
	require.True(t, example2.IsServiceUnavailable(stdErk))

	require.False(t, example2.IsServiceUnavailable(nil))
}

func TestIsAuthFailed(t *testing.T) {
	erk := example2.ErrorAuthFailed("test auth failed")
	t.Log(erk)
	require.True(t, example2.IsAuthFailed(erk))

	serviceErk := example2.ErrorServiceUnavailable("test service")
	t.Log(serviceErk)
	require.False(t, example2.IsAuthFailed(serviceErk))

	stdErk := errors.New(401, "AUTH_FAILED", "standard auth failed")
	t.Log(stdErk)
	require.True(t, example2.IsAuthFailed(stdErk))
}

func TestIsPermissionDenied(t *testing.T) {
	erk := example2.ErrorPermissionDenied("test permission denied")
	t.Log(erk)
	require.True(t, example2.IsPermissionDenied(erk))

	authErk := example2.ErrorAuthFailed("test auth")
	t.Log(authErk)
	require.False(t, example2.IsPermissionDenied(authErk))

	stdErk := errors.New(403, "PERMISSION_DENIED", "standard permission denied")
	t.Log(stdErk)
	require.True(t, example2.IsPermissionDenied(stdErk))
}

func TestIsRateLimitExceeded(t *testing.T) {
	erk := example2.ErrorRateLimitExceeded("test rate limit")
	t.Log(erk)
	require.True(t, example2.IsRateLimitExceeded(erk))

	authErk := example2.ErrorAuthFailed("test auth")
	t.Log(authErk)
	require.False(t, example2.IsRateLimitExceeded(authErk))

	stdErk := errors.New(429, "RATE_LIMIT_EXCEEDED", "standard rate limit")
	t.Log(stdErk)
	require.True(t, example2.IsRateLimitExceeded(stdErk))
}

func TestMetadataFieldName(t *testing.T) {
	erk := example2.ErrorServiceUnavailable("test with metadata")
	t.Log(erk)

	require.NotNil(t, erk.Metadata)

	apiCode, exists := erk.Metadata["api_error_code"]
	require.True(t, exists)

	expectedCode := fmt.Sprintf("%d", example2.ApiError_SERVICE_UNAVAILABLE.Number())
	require.Equal(t, expectedCode, apiCode)
}
