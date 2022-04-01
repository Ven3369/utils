package utils

// 往slice前面插入数据 和 append 相反 1.8 版本后可使用泛型来实现 不会就不会 那么痛苦
func Prepend(sc *[]interface{}, value interface{}) *[]interface{} {
	result := make([]interface{}, 1)
	result[0] = value
	for _, v := range *sc {
		result = append(result, v)
	}
	return &result
}

// 数组删除
// func SliceDel(sc *[]interface{}, index int) *[]interface{} {

// 	return sc
// }
