package u

func IfInterfaceEqualNilThen(data1, data2 interface{}) interface{} {
	return Compare(data1, data2, func(a, b interface{}) bool {
		return data1 != nil
	})
}
