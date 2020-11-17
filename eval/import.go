package eval

import (
	"github.com/leluxnet/carbon/hash"
	"github.com/leluxnet/carbon/typing"
	"strings"
)

const FileExtension = ".car"

var ImportFun typing.BFunction

func InitImportFun() {
	ImportFun = typing.BFunction{
		Name: "import",
		Dat: typing.ParamData{
			Params: []typing.Parameter{
				{
					"module",
					typing.StringClass,
					typing.String{},
				},
			},
		},
		Cal: func(_ typing.Object, args []typing.Object) typing.Throwable {
			props := Import(args[0].ToString())

			m := typing.NewMap()
			for name, o := range props {
				m.Items[hash.HashString(name)] = typing.Pair{Key: typing.String{Value: name}, Value: o}
			}

			return typing.Return{Data: m}
		},
	}
}

func Import(name string) map[string]typing.Object {
	e := BuiltinEnv()
	if strings.HasPrefix(name, "./") || strings.HasPrefix(name, "../") {
		_, props := RunFile(name+FileExtension, e)
		return props
	} else {
		// TODO: Import module
		return map[string]typing.Object{}
	}
}
