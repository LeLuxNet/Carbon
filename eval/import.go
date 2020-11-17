package eval

import (
	"fmt"
	"github.com/leluxnet/carbon/hash"
	"github.com/leluxnet/carbon/typing"
	"os"
	"path/filepath"
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
		Cal: func(_ typing.Object, args []typing.Object, file *typing.File) typing.Throwable {
			return ImportMap(args[0].ToString(), file)
		},
	}
}

var importCache = make(map[string]typing.Map)

func AbsPath(relative string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return filepath.Join(wd, relative), nil
}

func ImportMap(name string, fromFile *typing.File) typing.Throwable {
	var fName string

	if strings.HasPrefix(name, "./") || strings.HasPrefix(name, "../") {
		fName = filepath.Join(fromFile.Path, name+FileExtension)
	} else {
		fName = fmt.Sprintf("lib/%s/_index.car", name)

		// tmp
		var err error
		fName, err = AbsPath(fName)
		if err != nil {
			return typing.NewError(err.Error())
		}
	}

	if cache, ok := importCache[fName]; ok {
		return typing.Return{Data: cache}
	}

	e := BuiltinEnv()
	_, file := RunFile(fName, e)

	m := typing.NewMap()
	for name, o := range file.Props {
		m.Items[hash.HashString(name)] = typing.Pair{Key: typing.String{Value: name}, Value: o}
	}

	importCache[fName] = m
	return typing.Return{Data: m}
}
