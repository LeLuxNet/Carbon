package typing

import (
	"github.com/leluxnet/carbon/hash"
	"github.com/leluxnet/carbon/math"
	"math/big"
	"strings"
	"unicode"
)

const (
	toUpperCaseS   = "toUpperCase"
	toLowerCaseS   = "toLowerCase"
	capitalizeS    = "capitalize"
	capitalizeAllS = "capitalizeAll"
	splitS         = "split"
	parseIntS      = "parseInt"
	parseDoubleS   = "parseDouble"
)

var toUpperCase = BFunction{Name: toUpperCaseS, Cal: func(this Object, _ map[string]Object, _ []Object, _ map[string]Object, _ *File) Throwable {
	t, _ := this.(String)
	return Return{String{strings.ToUpper(t.Value)}}
}}

var toLowerCase = BFunction{Name: toLowerCaseS, Cal: func(this Object, _ map[string]Object, _ []Object, _ map[string]Object, _ *File) Throwable {
	t, _ := this.(String)
	return Return{String{strings.ToLower(t.Value)}}
}}

var capitalize = BFunction{Name: capitalizeS, Cal: func(this Object, _ map[string]Object, _ []Object, _ map[string]Object, _ *File) Throwable {
	t, _ := this.(String)
	if len(t.Value) == 0 {
		return Return{String{""}}
	}

	res := []rune(strings.ToLower(t.Value))
	res[0] = unicode.ToUpper(res[0])

	return Return{String{string(res)}}
}}

var capitalizeAll = BFunction{Name: capitalizeAllS, Cal: func(this Object, _ map[string]Object, _ []Object, _ map[string]Object, _ *File) Throwable {
	t, _ := this.(String)
	return Return{String{strings.Title(strings.ToLower(t.Value))}}
}}

var split = BFunction{Name: splitS, Dat: ParamData{Params: []Parameter{{Name: "sep", Type: StringClass, Default: String{" "}}}},
	Cal: func(this Object, params map[string]Object, _ []Object, _ map[string]Object, _ *File) Throwable {
		t, _ := this.(String)

		tmpSep, _ := params["sep"]
		sep := tmpSep.(String)

		var res []Object
		for _, s := range strings.Split(t.Value, sep.Value) {
			res = append(res, String{s})
		}

		return Return{Array{res}}
	}}

var parseInt = BFunction{Name: parseIntS, Cal: func(this Object, _ map[string]Object, _ []Object, _ map[string]Object, _ *File) Throwable {
	t, _ := this.(String)

	num, success := new(big.Int).SetString(t.Value, 10)
	if !success {
		return NewError("Can't parse int")
	}

	return Return{Int{num}}
}}

var parseDouble = BFunction{Name: parseDoubleS, Cal: func(this Object, _ map[string]Object, _ []Object, _ map[string]Object, _ *File) Throwable {
	t, _ := this.(String)

	num, _, err := new(big.Float).Parse(t.Value, 10)
	if err != nil {
		return NewError("Can't parse double")
	}

	return Return{Double{num}}
}}

var StringClass = NewNativeClass("string", Properties{
	toUpperCaseS:   toUpperCase,
	toLowerCaseS:   toLowerCase,
	capitalizeS:    capitalize,
	capitalizeAllS: capitalizeAll,
	parseIntS:      parseInt,
	parseDoubleS:   parseDouble,
})

func InitStringClass() {
	StringClass.Properties[splitS] = split
}

var _ Object = String{}

type String struct {
	Value string
}

func (o String) ToString() string {
	return o.Value
}

func (o String) Class() Class {
	return StringClass
}

func (o String) Eq(other Object) (Object, Throwable) {
	if other, ok := other.(String); ok {
		return Bool{o.Value == other.Value}, nil
	}
	return nil, nil
}

func (o String) NEq(other Object) (Object, Throwable) {
	if other, ok := other.(String); ok {
		return Bool{o.Value != other.Value}, nil
	}
	return nil, nil
}

func (o String) Add(other Object, first bool) (Object, Object) {
	if first {
		return String{o.Value + other.ToString()}, nil
	} else {
		return String{other.ToString() + o.Value}, nil
	}
}

func (o String) Mul(other Object, _ bool) (Object, Object) {
	switch other := other.(type) {
	case Int:
		var b strings.Builder
		for i := new(big.Int).Set(other.Value); i.Sign() > 0; i = i.Sub(i, math.IOne) {
			b.WriteString(o.Value)
		}
		return String{b.String()}, nil
	}
	return nil, nil
}

func (o String) Hash() uint64 {
	return hash.HashString(o.Value)
}

func (o String) GetIndex(key Object) (Object, Object) {
	switch key := key.(type) {
	case Int:
		chars := []rune(o.Value)

		i, err := resolveIntIndex(len(chars), key)
		if err != nil {
			return nil, err
		}

		return Char{chars[i]}, nil
	}
	return nil, nil
}
