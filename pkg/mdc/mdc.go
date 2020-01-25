// Provides a Mapped Diagnostic Context using the core Go Context
package mdc

import "context"

type DiagnosticMap struct {}

// Get a context with an empty DiagnosticMap value.
func Context(ctx context.Context) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, DiagnosticMap{}, make(map[string]interface{}))
}

// Adds the diagnostic key/value pair to the context's DiagnosticMap and returns
// the same context if the DiagnosticMap already exists.  If the Context doesn't
// have a DiagnosticMap, a new context is created with the DiagnosticMap added
// and the key/value pair set.
func WithDiagnostic(ctx context.Context, key string, value interface{}) context.Context{
	mdc := ctx.Value(DiagnosticMap{})
	if mdc == nil {
		mdc = make(map[string]interface{})
	}
	mdc.(map[string]interface{})[key] = value
	return context.WithValue(ctx, DiagnosticMap{}, mdc)
}

// Retrieves the DiagnosticMap from the context for use in logging (or whatever).
func Diagnostics(ctx context.Context) map[string]interface{} {
	mdc := ctx.Value(DiagnosticMap{})
	if mdc != nil {
		return mdc.(map[string]interface{})
	}
	return nil
}
