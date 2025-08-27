package example3_test

import (
	"fmt"
	"testing"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/orzkratos/errgen/internal/examples/example3"
	"github.com/orzkratos/errgenkratos"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	errgenkratos.SetMetadataFieldName("error_code_num")
	m.Run()
}

func TestErrorInvalidRequest(t *testing.T) {
	erk := example3.ErrorInvalidRequest("invalid request format: %s", "missing required field")
	t.Log(erk)

	require.NotNil(t, erk)
	require.Equal(t, int32(400), erk.Code)
	require.Equal(t, example3.BusinessError_INVALID_REQUEST.String(), erk.Reason)
	require.Equal(t, "invalid request format: missing required field", erk.Message)
}

func TestErrorOrderNotFound(t *testing.T) {
	erk := example3.ErrorOrderNotFound("order #%s not found", "ORD-12345")
	t.Log(erk)

	require.NotNil(t, erk)
	require.Equal(t, int32(404), erk.Code)
	require.Equal(t, example3.BusinessError_ORDER_NOT_FOUND.String(), erk.Reason)
	require.Equal(t, "order #ORD-12345 not found", erk.Message)
}

func TestErrorPaymentFailed(t *testing.T) {
	erk := example3.ErrorPaymentFailed("payment failed: %s", "insufficient funds")
	t.Log(erk)

	require.NotNil(t, erk)
	require.Equal(t, int32(402), erk.Code)
	require.Equal(t, example3.BusinessError_PAYMENT_FAILED.String(), erk.Reason)
	require.Equal(t, "payment failed: insufficient funds", erk.Message)
}

func TestErrorInsufficientInventory(t *testing.T) {
	erk := example3.ErrorInsufficientInventory("insufficient inventory for product %s", "PROD-456")
	t.Log(erk)

	require.NotNil(t, erk)
	require.Equal(t, int32(409), erk.Code)
	require.Equal(t, example3.BusinessError_INSUFFICIENT_INVENTORY.String(), erk.Reason)
	require.Equal(t, "insufficient inventory for product PROD-456", erk.Message)
}

func TestErrorOrderExpired(t *testing.T) {
	erk := example3.ErrorOrderExpired("order expired after %d minutes", 30)
	t.Log(erk)

	require.NotNil(t, erk)
	require.Equal(t, int32(410), erk.Code)
	require.Equal(t, example3.BusinessError_ORDER_EXPIRED.String(), erk.Reason)
	require.Equal(t, "order expired after 30 minutes", erk.Message)
}

func TestIsInvalidRequest(t *testing.T) {
	erk := example3.ErrorInvalidRequest("test invalid request")
	t.Log(erk)
	require.True(t, example3.IsInvalidRequest(erk))

	orderErk := example3.ErrorOrderNotFound("test order")
	t.Log(orderErk)
	require.False(t, example3.IsInvalidRequest(orderErk))

	stdErk := errors.New(400, "INVALID_REQUEST", "standard invalid request")
	t.Log(stdErk)
	require.True(t, example3.IsInvalidRequest(stdErk))

	require.False(t, example3.IsInvalidRequest(nil))
}

func TestIsOrderNotFound(t *testing.T) {
	erk := example3.ErrorOrderNotFound("test order not found")
	t.Log(erk)
	require.True(t, example3.IsOrderNotFound(erk))

	invalidErk := example3.ErrorInvalidRequest("test invalid")
	t.Log(invalidErk)
	require.False(t, example3.IsOrderNotFound(invalidErk))

	stdErk := errors.New(404, "ORDER_NOT_FOUND", "standard order not found")
	t.Log(stdErk)
	require.True(t, example3.IsOrderNotFound(stdErk))
}

func TestIsPaymentFailed(t *testing.T) {
	erk := example3.ErrorPaymentFailed("test payment failed")
	t.Log(erk)
	require.True(t, example3.IsPaymentFailed(erk))

	orderErk := example3.ErrorOrderNotFound("test order")
	t.Log(orderErk)
	require.False(t, example3.IsPaymentFailed(orderErk))

	stdErk := errors.New(402, "PAYMENT_FAILED", "standard payment failed")
	t.Log(stdErk)
	require.True(t, example3.IsPaymentFailed(stdErk))
}

func TestIsInsufficientInventory(t *testing.T) {
	erk := example3.ErrorInsufficientInventory("test insufficient inventory")
	t.Log(erk)
	require.True(t, example3.IsInsufficientInventory(erk))

	paymentErk := example3.ErrorPaymentFailed("test payment")
	t.Log(paymentErk)
	require.False(t, example3.IsInsufficientInventory(paymentErk))

	stdErk := errors.New(409, "INSUFFICIENT_INVENTORY", "standard insufficient inventory")
	t.Log(stdErk)
	require.True(t, example3.IsInsufficientInventory(stdErk))
}

func TestIsOrderExpired(t *testing.T) {
	erk := example3.ErrorOrderExpired("test order expired")
	t.Log(erk)
	require.True(t, example3.IsOrderExpired(erk))

	inventoryErk := example3.ErrorInsufficientInventory("test inventory")
	t.Log(inventoryErk)
	require.False(t, example3.IsOrderExpired(inventoryErk))

	stdErk := errors.New(410, "ORDER_EXPIRED", "standard order expired")
	t.Log(stdErk)
	require.True(t, example3.IsOrderExpired(stdErk))
}

func TestMetadataFieldName(t *testing.T) {
	erk := example3.ErrorInvalidRequest("test with metadata")
	t.Log(erk)

	require.NotNil(t, erk.Metadata)

	businessId, exists := erk.Metadata["error_code_num"]
	require.True(t, exists)

	expectedId := fmt.Sprintf("%d", example3.BusinessError_INVALID_REQUEST.Number())
	require.Equal(t, expectedId, businessId)
}
