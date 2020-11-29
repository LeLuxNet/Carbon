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

	procSetCursorPos *syscall.Proc
	procGetCursorPos *syscall.Proc
	procSendInput    *syscall.Proc
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

type point struct {
	x int32
	y int32
}

type input struct {
	t  uint32
	mi mouseInput
}

type mouseInput struct {
	x           int32
	y           int32
	button      int32
	dwFlags     uint32
	time        uint32
	dwExtraInfo uintptr
}

var XInput = typing.Module{Name: "_xinput", Items: map[string]typing.Object{
	"_load": typing.BFunction{
		Name: "_load",
		Cal: func(_ typing.Object, params map[string]typing.Object, _ []typing.Object, _ map[string]typing.Object, _ *typing.File) typing.Throwable {
			uDll, err := syscall.LoadDLL("user32.dll")
			if err != nil {
				return typing.NewError(err.Error())
			}
			defer uDll.Release()
			procSetCursorPos, err = uDll.FindProc("SetCursorPos")
			if err != nil {
				return typing.NewError(err.Error())
			}
			procGetCursorPos, err = uDll.FindProc("GetCursorPos")
			if err != nil {
				return typing.NewError(err.Error())
			}
			procSendInput, err = uDll.FindProc("SendInput")
			if err != nil {
				return typing.NewError(err.Error())
			}

			dll, err := syscall.LoadDLL("xinput1_4.dll")
			if err != nil {
				defer dll.Release()
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
	"_get_cursor": typing.BFunction{
		Name: "_get_cursor",
		Cal: func(_ typing.Object, params map[string]typing.Object, _ []typing.Object, _ map[string]typing.Object, _ *typing.File) typing.Throwable {
			p := &point{}

			procGetCursorPos.Call(uintptr(unsafe.Pointer(p)))
			// if r == 0 {
			t := typing.Tuple{Values: []typing.Object{
				typing.Int{Value: big.NewInt(int64(p.x))},
				typing.Int{Value: big.NewInt(int64(p.y))},
			}}
			return typing.Return{Data: t}
			// }
			// return typing.NewError(syscall.Errno(r).Error())
		},
	},
	"_set_cursor": typing.BFunction{
		Name: "_set_cursor",
		Dat: typing.ParamData{Params: []typing.Parameter{
			{Name: "x", Type: typing.IntClass},
			{Name: "y", Type: typing.IntClass},
		}},
		Cal: func(_ typing.Object, params map[string]typing.Object, _ []typing.Object, _ map[string]typing.Object, _ *typing.File) typing.Throwable {
			x := int32(params["x"].(typing.Int).Value.Int64())
			y := int32(params["y"].(typing.Int).Value.Int64())

			r, _, _ := procSetCursorPos.Call(uintptr(x), uintptr(y))
			if r == 0 {
				return nil
			}
			return nil // return typing.NewError(syscall.Errno(r).Error())
		},
	},
	"_click": typing.BFunction{
		Name: "_click",
		Dat: typing.ParamData{
			Params: []typing.Parameter{
				{Name: "data", Type: typing.IntClass},
				{Name: "flag", Type: typing.IntClass},
			},
		},
		Cal: func(_ typing.Object, params map[string]typing.Object, _ []typing.Object, _ map[string]typing.Object, _ *typing.File) typing.Throwable {
			dat := int32(params["data"].(typing.Int).Value.Uint64())
			flag := uint32(params["flag"].(typing.Int).Value.Uint64())

			data := input{t: 0, mi: mouseInput{button: dat, dwFlags: flag}}
			procSendInput.Call(uintptr(1), uintptr(unsafe.Pointer(&data)), unsafe.Sizeof(data))
			return nil
		},
	},
}}
