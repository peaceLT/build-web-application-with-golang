package helper

// 字符串切片转换为interface切片
func StringSliceToInterface(strings []string) []interface{} {
	interfaces := make([]interface{}, len(strings))
	for i, v := range strings {
		interfaces[i] = v
	}
	return interfaces
}
