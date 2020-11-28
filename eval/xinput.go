package eval

import (
	"github.com/leluxnet/carbon/typing"
	"math/big"
	"syscall"
	"unsafe"
)

var (
	procXInputGetState *syscall.Proc
	procXInputSetState *syscall.Proc
)

type xState struct {
	PacketNumber uint32
	Gamepad      xGamepad
}

type xGamepad struct {
	Buttons      uint16
	LeftTrigger  uint8
	RightTrigger uint8
	ThumbLX      int16
	ThumbLY      int16
	ThumbRX      int16
	ThumbRY      int16
}

type xVibration struct {
	LeftMotorSpeed  uint16
	RightMotorSpeed uint16
}

var XInput = typing.Module{Name: "_xinput", Items: map[string]typing.Object{
	"_load": typing.BFunction{
		Name: "_load",
		Cal: func(_ typing.Object, params map[string]typing.Object, _ []typing.Object, _ map[string]typing.Object, _ *typing.File) typing.Throwable {
			dll, err := syscall.LoadDLL("xinput1_4.dll")
			defer func() {
				if err != nil {
					dll.Release()
				}
			}()
			if err != nil {
				dll, err = syscall.LoadDLL("xinput1_3.dll")
				if err != nil {
					dll, err = syscall.LoadDLL("xinput9_1_0.dll")
					return typing.NewError(err.Error())
				}
			}
			procXInputGetState, err = dll.FindProc("XInputGetState")
			if err != nil {
				return typing.NewError(err.Error())
			}
			procXInputSetState, err = dll.FindProc("XInputSetState")
			if err != nil {
				return typing.NewError(err.Error())
			}
			return nil
		},
	},
	"_get": typing.BFunction{
		Name: "_get",
		Dat: typing.ParamData{Params: []typing.Parameter{
			{Name: "id", Type: typing.IntClass},
		}},
		Cal: func(_ typing.Object, params map[string]typing.Object, _ []typing.Object, _ map[string]typing.Object, _ *typing.File) typing.Throwable {
			controller := params["id"].(typing.Int).Value.Uint64()

			s := &xState{}

			r, _, _ := procXInputGetState.Call(uintptr(controller), uintptr(unsafe.Pointer(s)))
			if r == 0 {
				t := typing.Tuple{Values: []typing.Object{
					typing.Int{Value: big.NewInt(int64(s.Gamepad.Buttons))},
					typing.Int{Value: big.NewInt(int64(s.Gamepad.LeftTrigger))},
					typing.Int{Value: big.NewInt(int64(s.Gamepad.RightTrigger))},
					typing.Int{Value: big.NewInt(int64(s.Gamepad.ThumbLX))},
					typing.Int{Value: big.NewInt(int64(s.Gamepad.ThumbLY))},
					typing.Int{Value: big.NewInt(int64(s.Gamepad.ThumbRX))},
					typing.Int{Value: big.NewInt(int64(s.Gamepad.ThumbRY))},
				}}
				return typing.Return{Data: t}
			}
			return typing.NewError(syscall.Errno(r).Error())
		},
	},
	"_set": typing.BFunction{
		Name: "_set",
		Dat: typing.ParamData{Params: []typing.Parameter{
			{Name: "id", Type: typing.IntClass},
			{Name: "l", Type: typing.IntClass},
			{Name: "r", Type: typing.IntClass},
		}},
		Cal: func(_ typing.Object, params map[string]typing.Object, _ []typing.Object, _ map[string]typing.Object, _ *typing.File) typing.Throwable {
			controller := params["id"].(typing.Int).Value.Uint64()

			vib := &xVibration{
				LeftMotorSpeed:  uint16(params["l"].(typing.Int).Value.Uint64()),
				RightMotorSpeed: uint16(params["r"].(typing.Int).Value.Uint64()),
			}

			r, _, _ := procXInputSetState.Call(uintptr(controller), uintptr(unsafe.Pointer(vib)))
			if r == 0 {
				return nil
			}
			return typing.NewError(syscall.Errno(r).Error())
		},
	},
}}
