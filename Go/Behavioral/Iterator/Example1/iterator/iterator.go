package iterator

import "reflect"

type Iterator struct {
	data    reflect.Value
	len     int
	current int
}

func (i *Iterator) Value() interface{} {
	return i.data.Index(i.current).Interface()
}

func (i *Iterator) Next() bool {
	i.current++
	return i.current < i.len
}

func Iter(data interface{}) *Iterator {
	var v reflect.Value
	if reflect.TypeOf(data).Kind() != reflect.Slice {
		v = reflect.Append(reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(data)), 0, 1), reflect.ValueOf(data))
	} else {
		v = reflect.ValueOf(data)
	}
	return &Iterator{
		data:    v,
		len:     v.Len(),
		current: -1,
	}
}
