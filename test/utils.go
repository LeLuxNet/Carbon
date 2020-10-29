package test

import (
	"testing"
)

func AssertEq(t *testing.T, expect interface{}, val interface{}) {
	if expect != val {
		t.Fatalf("expected %v, got %v", expect, val)
	}
}
