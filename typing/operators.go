/* WARNING!
   This file is automatically generated through scripts/generator.py
   DON'T EDIT IT MANUALLY. ALL CHANGES WILL BE LOST! */

package typing

type Ltable interface {
    Lt(Object) (Object, Object)
}

func Lt(a, b Object) (Object, Throwable) {
    if a, ok := a.(Ltable); ok {
        res, err := a.Lt(b)
        if err != nil {
            return nil, Throw{Data: err}
        }
        if res != nil {
            return res, nil
        }
    }
    if b, ok := b.(Gtable); ok {
        res, err := b.Gt(a)
        if err != nil {
            return nil, Throw{Data: err}
        }
        return res, nil
    }
    return nil, nil
}

type Leable interface {
    Le(Object) (Object, Object)
}

func Le(a, b Object) (Object, Throwable) {
    if a, ok := a.(Leable); ok {
        res, err := a.Le(b)
        if err != nil {
            return nil, Throw{Data: err}
        }
        if res != nil {
            return res, nil
        }
    }
    if b, ok := b.(Geable); ok {
        res, err := b.Ge(a)
        if err != nil {
            return nil, Throw{Data: err}
        }
        return res, nil
    }
    return nil, nil
}

type Gtable interface {
    Gt(Object) (Object, Object)
}

func Gt(a, b Object) (Object, Throwable) {
    if a, ok := a.(Gtable); ok {
        res, err := a.Gt(b)
        if err != nil {
            return nil, Throw{Data: err}
        }
        if res != nil {
            return res, nil
        }
    }
    if b, ok := b.(Gtable); ok {
        res, err := b.Gt(a)
        if err != nil {
            return nil, Throw{Data: err}
        }
        return res, nil
    }
    return nil, nil
}

type Geable interface {
    Ge(Object) (Object, Object)
}

func Ge(a, b Object) (Object, Throwable) {
    if a, ok := a.(Geable); ok {
        res, err := a.Ge(b)
        if err != nil {
            return nil, Throw{Data: err}
        }
        if res != nil {
            return res, nil
        }
    }
    if b, ok := b.(Leable); ok {
        res, err := b.Le(a)
        if err != nil {
            return nil, Throw{Data: err}
        }
        return res, nil
    }
    return nil, nil
}

type Addable interface {
    Add(Object, bool) (Object, Object)
}

func Add(a, b Object) (Object, Throwable) {
    if a, ok := a.(Addable); ok {
        res, err := a.Add(b, true)
        if err != nil {
            return nil, Throw{Data: err}
        }
        if res != nil {
            return res, nil
        }
    }
    if b, ok := b.(Addable); ok {
        res, err := b.Add(a, false)
        if err != nil {
            return nil, Throw{Data: err}
        }
        return res, nil
    }
    return nil, nil
}

type Subable interface {
    Sub(Object, bool) (Object, Object)
}

func Sub(a, b Object) (Object, Throwable) {
    if a, ok := a.(Subable); ok {
        res, err := a.Sub(b, true)
        if err != nil {
            return nil, Throw{Data: err}
        }
        if res != nil {
            return res, nil
        }
    }
    if b, ok := b.(Subable); ok {
        res, err := b.Sub(a, false)
        if err != nil {
            return nil, Throw{Data: err}
        }
        return res, nil
    }
    return nil, nil
}

type Mulable interface {
    Mul(Object, bool) (Object, Object)
}

func Mul(a, b Object) (Object, Throwable) {
    if a, ok := a.(Mulable); ok {
        res, err := a.Mul(b, true)
        if err != nil {
            return nil, Throw{Data: err}
        }
        if res != nil {
            return res, nil
        }
    }
    if b, ok := b.(Mulable); ok {
        res, err := b.Mul(a, false)
        if err != nil {
            return nil, Throw{Data: err}
        }
        return res, nil
    }
    return nil, nil
}

type Divable interface {
    Div(Object, bool) (Object, Object)
}

func Div(a, b Object) (Object, Throwable) {
    if a, ok := a.(Divable); ok {
        res, err := a.Div(b, true)
        if err != nil {
            return nil, Throw{Data: err}
        }
        if res != nil {
            return res, nil
        }
    }
    if b, ok := b.(Divable); ok {
        res, err := b.Div(a, false)
        if err != nil {
            return nil, Throw{Data: err}
        }
        return res, nil
    }
    return nil, nil
}

type Modable interface {
    Mod(Object, bool) (Object, Object)
}

func Mod(a, b Object) (Object, Throwable) {
    if a, ok := a.(Modable); ok {
        res, err := a.Mod(b, true)
        if err != nil {
            return nil, Throw{Data: err}
        }
        if res != nil {
            return res, nil
        }
    }
    if b, ok := b.(Modable); ok {
        res, err := b.Mod(a, false)
        if err != nil {
            return nil, Throw{Data: err}
        }
        return res, nil
    }
    return nil, nil
}

type Powable interface {
    Pow(Object, bool) (Object, Object)
}

func Pow(a, b Object) (Object, Throwable) {
    if a, ok := a.(Powable); ok {
        res, err := a.Pow(b, true)
        if err != nil {
            return nil, Throw{Data: err}
        }
        if res != nil {
            return res, nil
        }
    }
    if b, ok := b.(Powable); ok {
        res, err := b.Pow(a, false)
        if err != nil {
            return nil, Throw{Data: err}
        }
        return res, nil
    }
    return nil, nil
}

type LShiftable interface {
    LShift(Object, bool) (Object, Object)
}

func LShift(a, b Object) (Object, Throwable) {
    if a, ok := a.(LShiftable); ok {
        res, err := a.LShift(b, true)
        if err != nil {
            return nil, Throw{Data: err}
        }
        if res != nil {
            return res, nil
        }
    }
    if b, ok := b.(LShiftable); ok {
        res, err := b.LShift(a, false)
        if err != nil {
            return nil, Throw{Data: err}
        }
        return res, nil
    }
    return nil, nil
}

type RShiftable interface {
    RShift(Object, bool) (Object, Object)
}

func RShift(a, b Object) (Object, Throwable) {
    if a, ok := a.(RShiftable); ok {
        res, err := a.RShift(b, true)
        if err != nil {
            return nil, Throw{Data: err}
        }
        if res != nil {
            return res, nil
        }
    }
    if b, ok := b.(RShiftable); ok {
        res, err := b.RShift(a, false)
        if err != nil {
            return nil, Throw{Data: err}
        }
        return res, nil
    }
    return nil, nil
}

type URShiftable interface {
    URShift(Object, bool) (Object, Object)
}

func URShift(a, b Object) (Object, Throwable) {
    if a, ok := a.(URShiftable); ok {
        res, err := a.URShift(b, true)
        if err != nil {
            return nil, Throw{Data: err}
        }
        if res != nil {
            return res, nil
        }
    }
    if b, ok := b.(URShiftable); ok {
        res, err := b.URShift(a, false)
        if err != nil {
            return nil, Throw{Data: err}
        }
        return res, nil
    }
    return nil, nil
}

type Andable interface {
    And(Object, bool) (Object, Object)
}

func And(a, b Object) (Object, Throwable) {
    if a, ok := a.(Andable); ok {
        res, err := a.And(b, true)
        if err != nil {
            return nil, Throw{Data: err}
        }
        if res != nil {
            return res, nil
        }
    }
    if b, ok := b.(Andable); ok {
        res, err := b.And(a, false)
        if err != nil {
            return nil, Throw{Data: err}
        }
        return res, nil
    }
    return nil, nil
}

type Orable interface {
    Or(Object, bool) (Object, Object)
}

func Or(a, b Object) (Object, Throwable) {
    if a, ok := a.(Orable); ok {
        res, err := a.Or(b, true)
        if err != nil {
            return nil, Throw{Data: err}
        }
        if res != nil {
            return res, nil
        }
    }
    if b, ok := b.(Orable); ok {
        res, err := b.Or(a, false)
        if err != nil {
            return nil, Throw{Data: err}
        }
        return res, nil
    }
    return nil, nil
}

type Xorable interface {
    Xor(Object, bool) (Object, Object)
}

func Xor(a, b Object) (Object, Throwable) {
    if a, ok := a.(Xorable); ok {
        res, err := a.Xor(b, true)
        if err != nil {
            return nil, Throw{Data: err}
        }
        if res != nil {
            return res, nil
        }
    }
    if b, ok := b.(Xorable); ok {
        res, err := b.Xor(a, false)
        if err != nil {
            return nil, Throw{Data: err}
        }
        return res, nil
    }
    return nil, nil
}

type Negable interface {
    Neg() Object
}

type Posable interface {
    Pos() Object
}

type Notable interface {
    Not() Object
}

type IndexSettable interface {
    SetIndex(key, value Object) Object
}

type IndexGettable interface {
    GetIndex(key Object) (Object, Object)
}

type Appendable interface {
    Append(Object) Object
}

type Containable interface {
    Contains(Object) (Object, Throwable)
}

type PropertyGettable interface {
    GetProperty(name string) (Object, Object)
}

type PropertySettable interface {
    SetProperty(name string, object Object) Object
}