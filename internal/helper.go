package interval

import "strings"

//TODO 等golang 1.18的泛型
func foreach(array []interface{}, fn func(item interface{})) {

}

func hasId(a, b []int) bool {
	for i := len(a); i >= 0; i-- {
		for j := len(b); j >= 0; j-- {
			if a[i] == b[j] {
				return true
			}
		}
	}
	return false
}


func hasTag(a, b []string) bool {
	for i := len(a); i >= 0; i-- {
		for j := len(b); j >= 0; j-- {
			if strings.EqualFold(a[i], b[j]) {
				return true
			}
		}
	}
	return false
}