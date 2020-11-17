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
					"name",
					typing.StringClass,
					typing.String{},
				},
			},
		},
		Cal: func(_ typing.Object, args []typing.Object) typing.Throwable {
			return ImportMap(args[0].ToString())
		},
	}
}

var importCache = make(map[string]typing.Map)

func ImportMap(name string) typing.Throwable {
	var fName string

	if strings.HasPrefix(name, "./") || strings.HasPrefix(name, "../") {
		fName = name + FileExtension
	} else {
		// TODO: Import module
		return nil
	}

	rName, err := resolveName(fName)
	if err != nil {
		return typing.NewError(err.Error())
	}

	if cache, ok := importCache[rName]; ok {
		return typing.Return{Data: cache}
	}

	e := BuiltinEnv()
	_, props := RunFile(rName, e)

	m := typing.NewMap()
	for name, o := range props {
		m.Items[hash.HashString(name)] = typing.Pair{Key: typing.String{Value: name}, Value: o}
	}

	importCache[rName] = m
	return typing.Return{Data: m}
}
