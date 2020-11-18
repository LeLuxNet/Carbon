package eval

import (
	"crypto/rand"
	"fmt"
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
			return ImportModule(args[0].ToString(), file)
		},
	}
}

var Sys = typing.Module{Name: "sys", Items: map[string]typing.Object{
	"_urandom": typing.BFunction{
		Name: "_urandom",
		Dat: typing.ParamData{
			Params: []typing.Parameter{
				{
					Name: "len",
					Type: typing.Int{}.Class(),
				},
			},
		},
		Cal: func(_ typing.Object, args []typing.Object, _ *typing.File) typing.Throwable {
			i := args[0].(typing.Int).Value.Int64()

			b := make([]byte, i)
			_, err := rand.Read(b)
			if err != nil {
				return typing.NewError(err.Error())
			}

			return typing.Return{Data: typing.Bytes{Values: b}}
		},
	},
}}

var importCache = map[string]typing.Module{
	"sys": Sys,
}

func AbsPath(relative string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return filepath.Join(wd, relative), nil
}

func ImportModule(name string, fromFile *typing.File) typing.Throwable {
	var mName string
	var fName string

	if cache, ok := importCache[name]; ok {
		return typing.Return{Data: cache}
	}

	if strings.HasPrefix(name, "./") || strings.HasPrefix(name, "../") {
		fName = filepath.Join(fromFile.Path, name+FileExtension)
		mName = fName
	} else {
		mName = name
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

	m := typing.Module{
		Name:  mName,
		Items: file.Props,
	}

	importCache[fName] = m
	return typing.Return{Data: m}
}
