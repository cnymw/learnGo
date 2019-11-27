package strategy

import "testing"

func TestContext_ExecuteStrategy(t *testing.T) {
	context := NewContext(&ImplOne{})
	context.ExecuteStrategy()

	context.SetStrategy(&ImplTwo{})
	context.ExecuteStrategy()
}
