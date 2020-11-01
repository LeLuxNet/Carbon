package typing

import (
	"testing"
)

func AssertOEq(t *testing.T, expect Object, val Object) {
	if !Eq(expect, val) {
		t.Fatalf("expected %s, got %s", expect.ToString(), val.ToString())
	}
}

func TestOperators(t *testing.T) {
	AssertOEq(t, Int{5}, Double{5})

	AssertOEq(t, Int{4}, Add(Int{3}, Bool{true}))
	AssertOEq(t, Double{2.5}, Div(Int{10}, Int{4}))
	AssertOEq(t, String{"aaa"}, Mult(String{"a"}, Int{3}))

	res, _ := Mod(Int{12}, Int{-9})
	AssertOEq(t, Int{-6}, res)

	_, err := Mod(Double{4.5}, Int{0})
	AssertOEq(t, ZeroDivisionError{}.Class(), err.TData().Class())
}
