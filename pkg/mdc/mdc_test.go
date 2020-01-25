package mdc

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContext(t *testing.T) {
	ctx := Context(nil)
	assert.NotNil(t, ctx)
	assert.NotNil(t, ctx.Value(DiagnosticMap{}))

	ctx = context.WithValue(context.Background(), "key", "value")
	assert.Nil(t, ctx.Value(DiagnosticMap{}))
	ctx = Context(ctx)
	assert.NotNil(t, ctx.Value(DiagnosticMap{}))
}

func TestWithDiagnostic(t *testing.T) {
	ctx := WithDiagnostic(context.Background(), "key", 1)
	assert.NotNil(t, ctx)
	assert.NotNil(t, ctx.Value(DiagnosticMap{}))
	assert.Contains(t, Diagnostics(ctx), "key")
	assert.Equal(t, 1, Diagnostics(ctx)["key"])
	ctx = WithDiagnostic(ctx, "otherKey", "otherValue")
	assert.Len(t, Diagnostics(ctx), 2)
	WithDiagnostic(ctx, "lastKey", true)
	assert.Len(t, Diagnostics(ctx), 3)
}
func TestDiagnostics(t *testing.T) {
	assert.Nil(t, Diagnostics(context.Background()))
	assert.Len(t, Diagnostics(WithDiagnostic(context.Background(), "key", true)), 1)
}


