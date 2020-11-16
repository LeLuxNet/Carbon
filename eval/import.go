package eval

import (
	"github.com/leluxnet/carbon/typing"
	"strings"
)

const FileExtension = ".car"

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
