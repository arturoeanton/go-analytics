package u

func Compare(data1, data2 interface{}, fxComp func(a, b interface{}) bool) interface{} {
	if fxComp(data1, data2) {
		return data1
	}
	return data2
}

func IfInterfaceEqualThen(data1, comparer, data2 interface{}) interface{} {
	return Compare(data1, data2, func(a, b interface{}) bool {
		return data1 != comparer
	})
}

func IfStringEqualThen(data1, comparer, data2 string) string {
	return IfInterfaceEqualThen(data1, comparer, data2).(string)
}
