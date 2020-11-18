package typing

import (
	"fmt"
	"math/big"
	"strings"
)

var _ Object = Bytes{}

type Bytes struct {
	Values []byte
}

const hextable = "0123456789abcdef"

func byteToHex(b byte) (byte, byte) {
	return hextable[b>>4], hextable[b&0x0f]
}

func displayBytes(b byte) []byte {
	if b >= 32 && b <= 127 {
		return []byte{b}
	} else {
		h1, h2 := byteToHex(b)
		return []byte{0x5c, 0x78, h1, h2}
	}
}

/* func (o Bytes) ToString() string {
	var builder strings.Builder
	for _, val := range o.Values {
		builder.Write(displayBytes(val))
	}
	return builder.String()
} */

func (o Bytes) ToString() string {
	var vals []string
	for _, val := range o.Values {
		b1, b2 := byteToHex(val)
		vals = append(vals, string(b1)+string(b2))
	}

	return strings.Join(vals, " ")
}

func (o Bytes) Class() Class {
	return NewNativeClass("bytes", Properties{})
}

func (o Bytes) Eq(other Object) (Object, Throwable) {
	switch other := other.(type) {
	case Array:
		if len(o.Values) != len(other.Values) {
			return Bool{false}, nil
		}

		for i, val := range o.Values {
			eq, err := Eq(Int{big.NewInt(int64(val))}, other.Values[i])
			if err != nil {
				return nil, err
			}

			if !Truthy(eq) {
				return Bool{false}, nil
			}
		}
		return Bool{true}, nil
	case Bytes:
		if len(o.Values) != len(other.Values) {
			return Bool{false}, nil
		}

		for i, val := range o.Values {
			if val != other.Values[i] {
				return Bool{false}, nil
			}
		}
		return Bool{true}, nil
	}
	return nil, nil
}

func (o Bytes) NEq(other Object) (Object, Throwable) {
	switch other := other.(type) {
	case Array:
		if len(o.Values) != len(other.Values) {
			return Bool{true}, nil
		}

		for i, val := range o.Values {
			eq, err := Eq(Int{big.NewInt(int64(val))}, other.Values[i])
			if err != nil {
				return nil, err
			}

			if !Truthy(eq) {
				return Bool{true}, nil
			}
		}
		return Bool{false}, nil
	case Bytes:
		if len(o.Values) != len(other.Values) {
			return Bool{true}, nil
		}

		for i, val := range o.Values {
			if val != other.Values[i] {
				return Bool{true}, nil
			}
		}
		return Bool{false}, nil
	}
	return nil, nil
}

func limitInt(o Object) (byte, Object) {
	if o, ok := o.(Int); ok {
		if o.Value.Sign() <= 0 && o.Value.Cmp(big.NewInt(255)) >= 0 {
			return byte(o.Value.Int64()), nil
		} else {
			return 0, Error{"Byte out of range %d has to be between 0 and 256"}
		}
	}
	return 0, nil
}

func (o Bytes) SetIndex(index, value Object) Object {
	switch index := index.(type) {
	case Int:
		i, err := resolveIntIndex(len(o.Values), index)
		if err != nil {
			return err
		}

		v, err := limitInt(value)
		if err != nil {
			return err
		}

		o.Values[i] = v
	}
	return nil
}

func (o Bytes) GetIndex(index Object) (Object, Object) {
	switch index := index.(type) {
	case Int:
		i, err := resolveIntIndex(len(o.Values), index)
		if err != nil {
			return nil, err
		}

		val := big.NewInt(int64(o.Values[i]))
		return Int{val}, nil
	}
	return nil, Error{fmt.Sprintf("'%s' is not of type int", index.ToString())}
}

func (o Bytes) Contains(value Object) (Object, Throwable) {
	for _, v := range o.Values {
		eq, err := Eq(Int{big.NewInt(int64(v))}, value)
		if err != nil {
			return nil, err
		}

		if Truthy(eq) {
			return Bool{true}, nil
		}
	}

	return Bool{false}, nil
}

func (o Bytes) Append(value Object) Object {
	v, err := limitInt(value)
	if err != nil {
		return err
	}

	o.Values = append(o.Values, v)
	return nil
}
