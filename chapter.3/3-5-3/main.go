package main

import (
	"log"
	"reflect"
	"strconv"
	"unsafe"
)

type MapStruct struct {
	Str     string  `map:"str"`
	StrPtr  *string `map:"strPtr"`
	Bool    bool    `map:"bool"`
	BoolPtr *bool   `map:"boolPtr"`
	Int     int     `map:"int"`
	IntPtr  *int    `map:"intPtr"`
}

func main() {
	src := MapStruct{
		Str:     "string-value",
		StrPtr:  &[]string{"string-ptr-value"}[0],
		Bool:    true,
		BoolPtr: &[]bool{false}[0],
		Int:     12345,
		IntPtr:  &[]int{67890}[0],
	}

	dest := map[string]string{}
	Encode(dest, &src)
	log.Println(dest)
}

func Encode(target map[string]string, src interface{}) error {
	v := reflect.ValueOf(src)
	e := v.Elem()
	t := e.Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		// 埋め込まれた構造体は再帰的に処理
		if f.Anonymous {
			if err := Encode(target, e.Field(i).Addr().Interface()); err != nil {
				return err
			}
			continue
		}

		// フィールドに構造体があれば再帰的に処理
		if f.Type.Kind() == reflect.Struct {
			if err := Encode(target, e.Field(i).Addr().Interface()); err != nil {
				return err
			}
			continue
		}

		// タグを取得
		// タグがなければフィールド名をそのまま使う
		key := f.Tag.Get("map")
		if key == "" {
			key = f.Name
		}

		// フィールドの型
		var k reflect.Kind
		// ポインターか否か
		var isP bool
		if f.Type.Kind() != reflect.Ptr {
			k = f.Type.Kind()
		} else {
			k = f.Type.Elem().Kind()
			isP = true
			// ポインターのポインターは無視
			if k == reflect.Ptr {
				continue
			}
		}

		switch k {
		case reflect.String:
			if isP {
				// nilなら読み込まない
				if e.Field(i).Pointer() != 0 {
					target[key] = *(*string)(unsafe.Pointer(e.Field(i).Pointer()))
				}
			} else {
				target[key] = e.Field(i).String()
			}
		case reflect.Bool:
			var b bool
			if isP {
				if e.Field(i).Pointer() != 0 {
					b = *(*bool)(unsafe.Pointer(e.Field(i).Pointer()))
				}
			} else {
				b = e.Field(i).Bool()
			}
			target[key] = strconv.FormatBool(b)
		case reflect.Int:
			var n int64
			if isP {
				if e.Field(i).Pointer() != 0 {
					n = int64(*(*int)(unsafe.Pointer(e.Field(i).Pointer())))
				}
			} else {
				n = e.Field(i).Int()
			}
			target[key] = strconv.FormatInt(n, 10)
		}
	}
	return nil
}
