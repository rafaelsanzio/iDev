package utils

import "reflect"

// FindMaxValue function to find max value in array
func FindMaxValue(arr []int) (max int) {
	max = 0
	for _, value := range arr {
		if value > max {
			max = value
		}
	}
	return max
}

//InArray function to check if valeu is in array
func InArray(v interface{}, in interface{}) (ok bool, i int) {
	val := reflect.Indirect(reflect.ValueOf(in))
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for ; i < val.Len(); i++ {
			if ok = v == val.Index(i).Interface(); ok {
				return
			}
		}
	}
	return
}
