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

	val, err := Add(Int{3}, Bool{true})
	if err != nil {
		t.Fatal(err.TData())
	}
	AssertOEq(t, Int{4}, val)

	val, err = Div(Int{10}, Int{4})
	if err != nil {
		t.Fatal(err.TData())
	}
	AssertOEq(t, Double{2.5}, val)

	val, err = Mul(String{"a"}, Int{3})
	if err != nil {
		t.Fatal(err.TData())
	}
	AssertOEq(t, String{"aaa"}, val)

	val, err = Mod(Int{12}, Int{-9})
	if err != nil {
		t.Fatal(err.TData())
	}
	AssertOEq(t, Int{-6}, val)

	_, err = Mod(Double{4.5}, Int{0})
	AssertOEq(t, ZeroDivisionError{}.Class(), err.TData().Class())
}
