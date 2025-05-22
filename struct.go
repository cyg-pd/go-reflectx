package reflectx

import "reflect"

func Name(v any) string {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}

func FullName(v any) string {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		return t.Elem().PkgPath() + "." + t.Elem().Name()
	} else {
		return t.PkgPath() + "." + t.Name()
	}
}
