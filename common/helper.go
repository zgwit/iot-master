package common

import "strings"

//TODO 等golang 1.18的泛型
func foreach(array []interface{}, fn func(item interface{})) {

}

func HasId(a, b []int) bool {
	for i := len(a); i >= 0; i-- {
		for j := len(b); j >= 0; j-- {
			if a[i] == b[j] {
				return true
			}
		}
	}
	return false
}


func HasTag(a, b []string) bool {
	for i := len(a); i >= 0; i-- {
		for j := len(b); j >= 0; j-- {
			if strings.EqualFold(a[i], b[j]) {
				return true
			}
		}
	}
	return false
}