package reflectx

import (
	"runtime"
	"strings"
)

func pkgSkip(level int) string {
	pc, _, _, ok := runtime.Caller(level + 1)
	if !ok {
		return ""
	}

	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)

	if parts[pl-2][0] == '(' {
		return strings.Join(parts[0:pl-2], ".")
	} else {
		return strings.Join(parts[0:pl-1], ".")
	}
}

func PackageName(level ...int) string {
	lvl := 0
	if len(level) > 0 {
		lvl = level[0]
	}
	return pkgSkip(lvl)
}
