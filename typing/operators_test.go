package typing

import (
	"math/big"
	"testing"
)

func AssertOEq(t *testing.T, expect Object, val Object) {
	eq, err := Eq(expect, val)
	if err != nil {
		t.Fatal(err)
	}
	if !Truthy(eq) {
		t.Fatalf("expected %s, got %s", expect.ToString(), val.ToString())
	}
}

func TestOperators(t *testing.T) {
	AssertOEq(t, Int{big.NewInt(5)}, Double{big.NewFloat(5)})

	val, err := Add(Int{big.NewInt(3)}, Bool{true})
	if err != nil {
		t.Fatal(err.TData())
	}
	AssertOEq(t, Int{big.NewInt(4)}, val)

	val, err = Div(Int{big.NewInt(10)}, Int{big.NewInt(4)})
	if err != nil {
		t.Fatal(err.TData())
	}
	AssertOEq(t, Double{big.NewFloat(2.5)}, val)

	val, err = Mul(String{"a"}, Int{big.NewInt(3)})
	if err != nil {
		t.Fatal(err.TData())
	}
	AssertOEq(t, String{"aaa"}, val)

	val, err = Mod(Int{big.NewInt(12)}, Int{big.NewInt(-9)})
	if err != nil {
		t.Fatal(err.TData())
	}
	AssertOEq(t, Int{big.NewInt(-6)}, val)

	val, err = Mod(Double{big.NewFloat(9.4)}, Double{big.NewFloat(-2.1)})
	if err != nil {
		t.Fatal(err.TData())
	}
	AssertOEq(t, Double{big.NewFloat(-1.1)}, val)

	_, err = Mod(Double{big.NewFloat(4.5)}, Int{big.NewInt(0)})
	AssertOEq(t, ZeroDivisionError{}.Class(), err.TData().Class())
}
