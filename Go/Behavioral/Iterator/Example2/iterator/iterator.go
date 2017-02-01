package iterator

import (
	"reflect"
)

func Iter(data interface{}) <-chan interface{} {
	ch := make(chan interface{}, 10)
	go func() {
		var v reflect.Value
		if reflect.TypeOf(data).Kind() != reflect.Slice {
			v = reflect.Append(reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(data)), 0, 1), reflect.ValueOf(data))
		} else {
			v = reflect.ValueOf(data)
		}
		l := v.Len()
		for i := 0; i < l; i++ {
			ch <- v.Index(i).Interface()
		}
		close(ch)
	}()
	return ch
}
