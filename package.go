package reflectx

import (
	"runtime"
	"strings"
)

func pkgSkip(level uint) string {
	pc, _, _, ok := runtime.Caller(int(level) + 1)
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

func PackageName(level ...uint) string {
	var lvl uint
	if len(level) > 0 {
		lvl = level[0]
	}
	return pkgSkip(lvl)
}
