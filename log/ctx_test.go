package log

import (
	"context"
	"reflect"
	"testing"
)

func TestCtx(t *testing.T) {
	l1 := Logger.With().Bool("test", true).Logger()
	ctx := l1.WithContext(context.Background())
	l2 := Ctx(ctx)
	if reflect.DeepEqual(l2, &Logger) {
		t.Error("returned the default logger against a context with a registered logger")
	}
	if !reflect.DeepEqual(l1, *l2) {
		t.Error("failed to return the correct logger")
	}
}

func TestCtxNoRegistration(t *testing.T) {
	l := Ctx(context.Background())
	if !reflect.DeepEqual(l, &Logger) {
		t.Error("failed to return the default logger against a context without a registered logger")
	}
}
